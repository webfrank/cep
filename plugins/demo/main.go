//go:build tinygo.wasm

package main

import (
	"context"
	"fmt"

	"bizmate.it/cep/internal/proto"
)

type MyPlugin struct {
	hostFunctions proto.HostFunctions
}

// main is required for TinyGo to compile to Wasm.
func main() {
	proto.RegisterPlugin(&MyPlugin{
		hostFunctions: proto.NewHostFunctions(),
	})
}

var _ proto.Plugin = (*MyPlugin)(nil)

func (m MyPlugin) Handle(ctx context.Context, request *proto.Event) (*proto.Event, error) {
	// Logging via the host function
	m.hostFunctions.Logger(ctx, &proto.Message{
		Line: fmt.Sprintf("DEMO: %v", request.Value),
	})
	return request, nil
}
