package actors

import (
	"context"
	"encoding/binary"
	"encoding/json"
	"math"
	"time"

	messages "bizmate.it/WASMPluginTest/src/protos"
	extism "github.com/extism/go-sdk"
	"github.com/tochemey/goakt/v2/log"

	goakt "github.com/tochemey/goakt/v2/actors"
)

// FIXME: This structure is good for prototyping only, we need to implement a better system
type WASMPlugin struct {
	self   *goakt.PID
	ctx    context.Context
	cancel context.CancelFunc
	plugin *extism.Plugin

	init bool

	nodesOut []string //Actors' names
}

func NewWASMPlugin() *WASMPlugin {
	return &WASMPlugin{}
}

func (p *WASMPlugin) PreStart(context.Context) (err error) {
	log.DefaultLogger.Info("[WASMPLUGIN] Starting")

	p.init = false //There's probably a better way to do this

	return nil
}

func (p *WASMPlugin) Receive(ctx *goakt.ReceiveContext) {
	switch ctx.Message().(type) {
	case *messages.Init:
		p.handleInit(ctx.Message().(*messages.Init))
	case *messages.Event:
		p.handleEvent(ctx.Message().(*messages.Event))
	default:
		ctx.Unhandled()
	}
}

func (p *WASMPlugin) PostStop(context.Context) error {
	return nil
}

func (p *WASMPlugin) handleInit(ev *messages.Init) {
	if ev.WASMUrl == "" {
		log.DefaultLogger.Error("[WASMPLUGIN] An empty WASMUrl was passed")
		return
	}

	err := json.Unmarshal([]byte(ev.OutputsJSON), &p.nodesOut)
	if err != nil {
		log.DefaultLogger.Warnf("[WASMPLUGIN] An error occurred whilst getting output nodes thus the message cannot be forwarded, error message: %s", err.Error())
		p.nodesOut = []string{}
	}

	manifest := extism.Manifest{
		Wasm: []extism.Wasm{
			extism.WasmUrl{
				Url: ev.GetWASMUrl(),
			},
		},
	}
	//It is FUNDAMENTAL that the context passed is context.Background() otherwise it will block the entire program.
	plugin, err := extism.NewPlugin(context.Background(), manifest, extism.PluginConfig{EnableWasi: true}, []extism.HostFunction{})
	if err != nil {
		log.DefaultLogger.Errorf("[WASMPLUGIN] An error occurred whist instantiating plugin: %s", err.Error())
		return
	}

	p.plugin = plugin

	p.init = true
	log.DefaultLogger.Info("[WASMPLUGIN] A plugin has started")
	return
}

func (p *WASMPlugin) handleEvent(ev *messages.Event) {

	if !p.init {
		log.DefaultLogger.Error("[WASMPLUGIN] An event was passed before the plugin was set up.")
		return
	}

	floatBuffer := make([]byte, 8)
	binary.LittleEndian.PutUint64(floatBuffer, extism.EncodeF32(ev.Value))

	exit, out, err := p.plugin.Call("run", floatBuffer) //TODO: Fetch data from previous nodes
	if err != nil {
		log.DefaultLogger.Errorf("[WASMPLUGIN] Plugin#run() returned with code %d, error message: %s", exit, err.Error())
	}
	if len(p.nodesOut) == 0 {
		return
	}

	floatOut := math.Float32frombits(binary.LittleEndian.Uint32(out))
	log.DefaultLogger.Infof("[WASMPLUGIN] FIXME: Implement forwarding to the next nodes, output was: %f", floatOut)

	for i := 0; i < len(p.nodesOut); i++ {
		log.DefaultLogger.Infof("[WASMPLUGIN] Forwarding to: %s", p.nodesOut[i])
		_, err := p.self.SendSync(p.ctx, p.nodesOut[i], &messages.Event{
			Ts:     time.Now().UnixMilli(),
			Serial: "",
			Value:  floatOut,
		}, 1*time.Second)
		if err != nil {
			log.DefaultLogger.Errorf("[WASMPLUGIN] An error occurred during forwarding: %s", err.Error())
		}
		log.DefaultLogger.Infof("[WASMPLUGIN] Forwarded")
	}
}
