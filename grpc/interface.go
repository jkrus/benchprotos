// Copyright © 2020-2021 The EVEN Solutions Developers Team

package grpc

import (
	ctx "github.com/evenlab/go-kit/context"

	"benchprotos/grpc/proto/pb"
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
		// UnsafeQuotesServiceServer embedded grpc service interface.
		pb.UnsafeQuotesServiceServer

		// UnsafePingPongServer embedded grpc service interface.
		pb.UnsafePingPongServer

		// Server embedded server interface.
		Server
	}
)
