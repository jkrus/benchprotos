// Copyright © 2020-2021 The EVEN Solutions Developers Team

package app

import (
	"os"

	"github.com/evenlab/go-kit/context"
	"github.com/goava/di"
	"github.com/urfave/cli/v2"
)

const (
	appDefaultName  = "benchprotos"
	appDefaultUsage = "BenchMarks protos"
)

type (
	// App describes cli application.
	App struct {
		*cli.App
	}
)

// NewApp returns an application.
func NewApp() *App {
	cli.VersionFlag = &cli.BoolFlag{
		Name:    "version",
		Aliases: []string{"ver"},
		Usage:   "print the version",
	}

	return &App{
		App: &cli.App{
			Name:    appDefaultName,
			Usage:   appDefaultUsage,
			Version: version(),
		},
	}
}

// Start starts the application.
func Start(ctx context.Context, dic *di.Container, app *App) error {
	if err := dic.Invoke(addStartCommand); err != nil {
		return err
	}

	return app.RunContext(ctx, os.Args)
}
