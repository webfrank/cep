package actors

import (
	"context"
	"math/rand/v2"
	"time"

	"bizmate.it/cep/internal/proto"
	goakt "github.com/tochemey/goakt/v2/actors"
	"github.com/tochemey/goakt/v2/goaktpb"
	"github.com/tochemey/goakt/v2/log"
)

type Source struct {
	self   *goakt.PID
	ctx    context.Context
	cancel context.CancelFunc
}

func NewSource() *Source {
	return &Source{}
}

func (s *Source) PreStart(ctx context.Context) (err error) {
	log.DefaultLogger.Info("[SOURCE] Starting")

	return nil
}

func (s *Source) Receive(ctx *goakt.ReceiveContext) {
	switch ctx.Message().(type) {
	case *goaktpb.PostStart:
		s.self = ctx.Self()
		s.ctx, s.cancel = context.WithCancel(context.Background())

		log.DefaultLogger.Infof("[SOURCE] Starting %s - %s", s.self.Name(), s.self.ID())
		go s.run()
	}
}

func (s *Source) PostStop(context.Context) error {
	log.DefaultLogger.Info("[SOURCE] Stopping")
	s.cancel()

	return nil
}

func (s *Source) run() {
	timer := time.NewTicker(3 * time.Second)
	defer timer.Stop()

	for {
		select {
		case <-s.ctx.Done():
			return
		case <-timer.C:
			log.DefaultLogger.Infof("[SOURCE] Tick: %s", s.self.ID())
			response, err := s.self.SendSync(s.ctx, "Plugin-multiply", &proto.Event{
				Ts:     1,
				Serial: "ABC",
				Value:  rand.Float32(),
			}, 5*time.Second)
			if err != nil {
				log.DefaultLogger.Errorf("[SOURCE] Error sending event: %v", err)
				continue
			}
			response, err = s.self.SendSync(s.ctx, "Plugin-demo", response, 5*time.Second)
			if err != nil {
				log.DefaultLogger.Errorf("[SOURCE] Error sending event: %v", err)
				continue
			}
		}
	}
}
