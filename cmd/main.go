package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/apex/log/handlers/json"

	srv "go-fiber-api/server"

	"github.com/apex/log"
	cli "github.com/urfave/cli/v2"
)

func main() {
	app := cli.NewApp()
	app.Name = "vfeatureflag"
	app.Usage = "vfeatureflag is a tool for managing feature flags"
	app.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:     "config",
			Usage:    "Config file for vfeatureflag",
			Required: false,
		},
	}

	log.SetHandler(json.New(os.Stderr))

	app.Action = func(ctx *cli.Context) error {
		config, err := srv.ReadConfig(ctx.String("config"))
		if err != nil {
			return fmt.Errorf("failed to load config. %w", err)
		}

		server, err := srv.NewServer(config)
		if err != nil {
			return fmt.Errorf("failed to create server. %w", err)
		}

		fiberApp := srv.SetupFiber(server, srv.CreateFiberConfig())

		serverCh := make(chan error)
		go func() {
			log.Info("Starting vFeatureFlag server")

			if err := fiberApp.Listen("0.0.0.0:3000"); err != nil {
				serverCh <- err
			}
			serverCh <- nil
		}()

		// Wait for a signal to gracefully shutdown the server
		ch := make(chan os.Signal, 1)
		signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)

		select {
		case err := <-serverCh:
			if err != nil {
				return fmt.Errorf("failed to start server: %w", err)
			}
		case sig := <-ch:
			log.Infof("Server shutting down, received signal %v", sig)
			if err := fiberApp.Shutdown(); err != nil {
				return fmt.Errorf("error shutting down server: %w", err)
			}
		}

		return nil
	}

	if err := app.Run(os.Args); err != nil {
		log.Errorf("unexpected exit error: %s", err)
		os.Exit(1)
	}
}
