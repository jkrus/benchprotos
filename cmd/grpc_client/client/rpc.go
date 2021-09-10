// Copyright © 2020-2021 The EVEN Solutions Developers Team

package client

import (
	"context"
	"fmt"
	"time"

	"google.golang.org/grpc"
)

const (
	// RPCAddress is the address of a gRPC server to connect to.
	RPCAddress = "localhost:4317"
)

var (
	con *grpc.ClientConn // global variables are harmful
	ctx context.Context  // global variables are harmful

	cancel  func()
	timeout = 30 * time.Second // why not constant?
)

func connect() error {
	var err error
	// Set up a connection to the server
	con, err = grpc.Dial(RPCAddress, grpc.WithInsecure())
	if err != nil {
		return fmt.Errorf("rpc connect: %w", err)
	}

	freshContext()

	return nil
}

func terminate() {
	defer cancel()
	if con != nil {
		_ = con.Close()
	}
}

func freshContext() {
	// Create a context
	ctx, cancel = context.WithTimeout(context.Background(), timeout)
}
