// Copyright © 2020-2021 The EVEN Solutions Developers Team

package app

import (
	"log"

	"github.com/evenlab/go-kit/context"
	"github.com/goava/di"
	"github.com/urfave/cli/v2"

	"github.com/jkrus/benchprotos/grpc"
)

// addStartCommand appends start action to application.
func addStartCommand(ctx context.Context, dic *di.Container, app *App) {
	app.Commands = append(app.Commands, &cli.Command{
		Name:  "start",
		Usage: "Starts " + appDefaultUsage,
		Action: func(cc *cli.Context) error {
			// provide cli context
			if err := dic.Provide(func() *cli.Context { return cc }); err != nil {
				return ErrProvideCliContext(err)
			}

			return invokeServices(dic)
		},
		After: func(cc *cli.Context) error {
			<-cc.Done() // wait while context canceled

			ctx.Cancel()
			ctx.WgWait() // wait while all workers finished

			log.Println("Application shutdown complete")

			return nil
		},
	})
}

// invokeServices tries to invoke required
// services from application DI container.
func invokeServices(dic *di.Container) error {
	// invoke grpc service starter
	if err := dic.Invoke(grpc.Start); err != nil {
		return ErrStartServiceGRPC(err)
	}

	return nil
}
