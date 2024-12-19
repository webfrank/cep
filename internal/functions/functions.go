package functions

import (
	"context"

	"github.com/tochemey/goakt/v2/log"

	"bizmate.it/cep/internal/proto"
	"github.com/knqyf263/go-plugin/types/known/emptypb"
)

// myHostFunctions implements HostFunctions
type MyHostFunctions struct{}

// HttpGet is embedded into the plugin and can be called by the plugin.
func (MyHostFunctions) Logger(_ context.Context, request *proto.Message) (*emptypb.Empty, error) {
	log.DefaultLogger.Infof("[PLUGIN] %s", request.GetLine())

	return &emptypb.Empty{}, nil
}
