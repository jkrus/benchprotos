package crpc

import (
	ctx "github.com/evenlab/go-kit/context"

	"benchprotos/crpc/cpb"
)

type (
	// Server embedded server interface.
	Server interface {
		// Serve accepts incoming connections.
		Serve(ctx.Context)

		// Stop stops network service.
		Stop() error
	}

	// Service describes grpc service interface.
	Service interface {
		// QuotesUpdate_Server embedded crpc service interface.
		cpb.QuotesUpdate_Server

		// Server embedded server interface.
		Server
	}
)
