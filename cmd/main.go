package main

import (
	"context"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
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
	// instantiate the static discovery provider
	disco := static.NewDiscovery(&discoConfig)

	clusterConfig := goakt.
		NewClusterConfig().
		WithDiscovery(disco).
		WithPartitionCount(20).
		WithMinimumPeersQuorum(1).
		WithReplicaCount(1).
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

	// wait for the cluster to be ready
	time.Sleep(time.Duration(rand.Float32()*5)*time.Second + time.Second)

	// spawn the actors
	log.DefaultLogger.Infof("Spawning actors InCluster: %t - %s", actorSystem.InCluster(), actorSystem.Host())
	actorSystem.Spawn(ctx, "Source", actors.NewSource())
	actorSystem.Spawn(ctx, "Plugin-multiply", actors.NewPlugin())
	actorSystem.Spawn(ctx, "Plugin-demo", actors.NewPlugin())

	// wait for ctrl+c and exit
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	// signal.Notify(c, os.Kill)
	signal.Notify(c, syscall.SIGTERM)
	signal.Notify(c, syscall.SIGINT)

	<-c
	err = actorSystem.Stop(ctx)
	if err != nil {
		log.DefaultLogger.Errorf("app stopped with error: %v", err)
	} else {
		log.DefaultLogger.Info("app exited")
	}
}
