package main

import (
	"context"
	"os"
	"os/signal"
	"runtime"
	"time"

	"bizmate.it/cep/internal/actors"
	"bizmate.it/cep/internal/global"
	"bizmate.it/cep/internal/proto"
	goakt "github.com/tochemey/goakt/v2/actors"
	"github.com/tochemey/goakt/v2/discovery/static"
	"github.com/tochemey/goakt/v2/log"
)

func main() {
	ctx := context.Background()

	logger := log.New(log.DebugLevel, os.Stdout)

	// define the discovery options
	discoConfig := static.Config{
		Hosts: []string{
			"node1:3322",
			"node2:3322",
			"node3:3322",
		},
	}
	// instantiate the dnssd discovery provider
	disco := static.NewDiscovery(&discoConfig)

	clusterConfig := goakt.
		NewClusterConfig().
		WithDiscovery(disco).
		WithPartitionCount(20).
		WithMinimumPeersQuorum(2).
		WithReplicaCount(2).
		WithDiscoveryPort(3322).
		WithPeersPort(3320).
		WithKinds(new(actors.Plugin), new(actors.Source))

	// grab the host
	host, _ := os.Hostname()

	// create the actor system.
	actorSystem, err := goakt.NewActorSystem("iotbuilder",
		goakt.WithLogger(logger),
		goakt.WithPassivationDisabled(),
		goakt.WithActorInitMaxRetries(3),
		goakt.WithShutdownTimeout(30*time.Second),
		goakt.WithRemoting(host, 50052),
		goakt.WithCluster(clusterConfig),
	)
	if err != nil {
		panic(err)
	}

	// start the actor system
	err = actorSystem.Start(ctx)
	if err != nil {
		panic(err)
	}

	// Initialize a plugin loader
	p, err := proto.NewPluginPlugin(ctx)
	if err != nil {
		panic(err)
	}
	global.Loader = p

	// spawn the actors
	log.DefaultLogger.Infof("Spawning actors InCluster: %t - %s", actorSystem.InCluster(), actorSystem.Host())
	actorSystem.Spawn(ctx, "Source", actors.NewSource(), goakt.WithSupervisorStrategies(
		goakt.NewSupervisorStrategy(goakt.PanicError{}, goakt.NewRestartDirective()),
		goakt.NewSupervisorStrategy(&runtime.PanicNilError{}, goakt.NewRestartDirective()),
	))
	actorSystem.Spawn(ctx, "Plugin-multiply", actors.NewPlugin(), goakt.WithSupervisorStrategies(
		goakt.NewSupervisorStrategy(goakt.PanicError{}, goakt.NewRestartDirective()),
		goakt.NewSupervisorStrategy(&runtime.PanicNilError{}, goakt.NewRestartDirective()),
	))
	actorSystem.Spawn(ctx, "Plugin-demo", actors.NewPlugin(), goakt.WithSupervisorStrategies(
		goakt.NewSupervisorStrategy(goakt.PanicError{}, goakt.NewRestartDirective()),
		goakt.NewSupervisorStrategy(&runtime.PanicNilError{}, goakt.NewRestartDirective()),
	))

	// wait for ctrl+c and exit
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	<-c
	err = actorSystem.Stop(ctx)
	if err != nil {
		log.DefaultLogger.Errorf("app stopped with error: %v", err)
	} else {
		log.DefaultLogger.Info("app exited")
	}
}
