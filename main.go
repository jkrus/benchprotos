// Copyright © 2020-2021 The EVEN Solutions Developers Team

package main

import (
	"log"

	"github.com/evenlab/go-kit/context"
	"github.com/goava/di"

	"benchprotos/app"
	"benchprotos/grpc"
)

func main() {

	// create the application DI-container
	c, err := di.New(
		di.Provide(app.NewApp),         // provide the app
		di.Provide(context.NewContext), // provide the app context
		di.Provide(grpc.NewService),    // provide the app network service
	)
	if err != nil {
		log.Fatal(err)
	}

	// invoke application starter
	if err = c.Invoke(app.Start); err != nil {
		log.Fatal(err)
	}
}
