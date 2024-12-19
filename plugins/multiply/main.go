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

func (m MyPlugin) Handle(ctx context.Context, request *proto.Event) (*proto.Event, error) {
	// Logging via the host function
	m.hostFunctions.Logger(ctx, &proto.Message{
		Line: fmt.Sprintf("MULTIPLY: %v", request.Value),
	})
	request.Value = request.Value * 2
	return request, nil
}
