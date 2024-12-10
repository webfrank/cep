package main

import (
	"context"
	"os"
	"os/signal"
	"time"

	"bizmate.it/WASMPluginTest/src/actors"
	messages "bizmate.it/WASMPluginTest/src/protos"
	goakt "github.com/tochemey/goakt/v2/actors"
	"github.com/tochemey/goakt/v2/log"
)

func main() {
	ctx := context.Background()
	logger := log.DefaultLogger

	actorSystem, err := goakt.NewActorSystem("ActorSystem",
		goakt.WithPassivationDisabled(),
		goakt.WithLogger(logger),
		goakt.WithActorInitMaxRetries(3),
		goakt.WithSupervisorDirective(goakt.NewResumeDirective()))
	if err != nil {
		/*
			log.DefaultLogger.Fatalf("[MAIN] A fatal error occurred whilst creating the actor system: %s", err.Error())
			os.Exit(-1)
		*/
		panic(err)
	}

	//Since Start returns only one argument it's way more concise to put it inside the if statement directly
	if actorSystem.Start(ctx) != nil {
		/*
			log.DefaultLogger.Fatalf("[MAIN] A fatal error occurred whilst starting the actor system: %s", err.Error())
			os.Exit(-1)
		*/
		panic(err)
	}

	//goakt's developer suggests to wait one second to allow the actor system to start properly
	time.Sleep(1 * time.Second)

	// Start plugin one
	a, err := actorSystem.Spawn(ctx, "PLUGIN_01_FOO", actors.NewWASMPlugin())
	if err != nil {
		panic(err)
	}

	a.SendSync(ctx, "PLUGIN_01_FOO", &messages.Init{
		WASMUrl:            "http://127.0.0.1:3000/wasm/timesTwo/plugin.wasm",
		OutputsJSON:        `["PLUGIN_02_BAR"]`,
		EnvVars:            "{}",
		AcceptsEmptyOutput: false,
	}, 2*time.Second)

	// Start plugin two
	b, err := actorSystem.Spawn(ctx, "PLUGIN_02_BAR", actors.NewWASMPlugin())
	if err != nil {
		panic(err)
	}

	b.SendSync(ctx, "PLUGIN_02_BAR", &messages.Init{
		WASMUrl:            "http://127.0.0.1:3000/wasm/timesTwo/plugin.wasm",
		OutputsJSON:        `[]`,
		EnvVars:            "{}",
		AcceptsEmptyOutput: false,
	}, 2*time.Second)

	log.DefaultLogger.Info("STARTED ALL PLUGINS")

	a.SendAsync(ctx, "PLUGIN_01_FOO", &messages.Event{
		Ts:     time.Now().UnixMilli(),
		Serial: "",
		Value:  0.2,
	})

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	<-c
	err = actorSystem.Stop(ctx)
	if err != nil {
		logger.Errorf("Application stopped with error: %s", err.Error())
	} else {
		logger.Info("Application stopped successfully")
	}
}
