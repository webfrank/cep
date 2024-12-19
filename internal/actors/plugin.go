package actors

import (
	"context"
	"fmt"
	"regexp"

	"bizmate.it/cep/internal/functions"
	"bizmate.it/cep/internal/global"
	"bizmate.it/cep/internal/proto"
	goakt "github.com/tochemey/goakt/v2/actors"
	"github.com/tochemey/goakt/v2/goaktpb"
	"github.com/tochemey/goakt/v2/log"
)

type Plugin struct {
	self   *goakt.PID
	ctx    context.Context
	cancel context.CancelFunc
	plugin proto.Plugin
}

func NewPlugin() *Plugin {
	return &Plugin{}
}

func (p *Plugin) PreStart(ctx context.Context) (err error) {
	log.DefaultLogger.Info("[PLUGIN] Starting")

	return nil
}

func (p *Plugin) Receive(ctx *goakt.ReceiveContext) {
	switch ctx.Message().(type) {
	case *goaktpb.PostStart:
		p.self = ctx.Self()
		p.ctx, p.cancel = context.WithCancel(context.Background())

		re := regexp.MustCompile(`-(\w+)$`)
		matches := re.FindStringSubmatch(p.self.Name())
		if len(matches) < 2 {
			log.DefaultLogger.Error("[PLUGIN] Invalid actor name format")
			return
		}
		name := matches[1]

		log.DefaultLogger.Infof("[PLUGIN] Starting %s - %s", name, p.self.ID())

		err := p.loadPlugin(p.ctx, name)
		if err != nil {
			log.DefaultLogger.Errorf("[PLUGIN] Error loading plugin: %v", err)
			ctx.Unhandled()
		}
	case *proto.Event:
		if p.plugin == nil {
			log.DefaultLogger.Error("[PLUGIN] Plugin not loaded")
			ctx.Unhandled()
			return
		}
		log.DefaultLogger.Infof("[PLUGIN] Tick: %s", p.self.ID())
		response, err := p.plugin.Handle(p.ctx, ctx.Message().(*proto.Event))
		if err != nil {
			log.DefaultLogger.Errorf("[PLUGIN] Error handling event: %v", err)
			return
		}
		ctx.Response(response)
	}
}

func (p *Plugin) PostStop(context.Context) error {
	log.DefaultLogger.Info("[PLUGIN] Stopping")
	// p.plugin.Close(p.ctx)
	p.cancel()
	return nil
}

func (p *Plugin) loadPlugin(ctx context.Context, name string) (err error) {
	// Load a plugin
	plugin, err := global.Loader.Load(ctx, fmt.Sprintf("plugins/%s/plugin.wasm", name), functions.MyHostFunctions{})
	if err != nil {
		return err
	}
	p.plugin = plugin

	return nil
}
