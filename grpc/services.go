// Copyright © 2020-2021 The EVEN Solutions Developers Team

package grpc

import (
	"context"
	"time"

	"benchprotos/grpc/proto/pb"
)

// QuotesUpdate implements service api interface.
func (s *service) QuotesUpdate(_ context.Context, req *pb.QuotesUpdateRequest) (*pb.QuotesUpdateResponse, error) {

	time.Sleep(500)

	return &pb.QuotesUpdateResponse{Response: req.String()}, nil
}

func (s *service) PingMessage(_ context.Context, req *pb.PingPongRequest) (*pb.PingPongResponse, error) {

	time.Sleep(500)

	return &pb.PingPongResponse{Result: req.GetPingMessage()}, nil
}

/*
func server(c net.Conn) error {
	// Create a new locally implemented HashFactory.
	main := hashes.HashFactory_ServerToClient(hashFactory{})
	// Listen for calls, using the HashFactory as the bootstrap interface.
	conn := rpc.NewConn(rpc.StreamTransport(c), rpc.MainInterface(main.Client))
	// Wait for connection to abort.
	err := conn.Wait()
	return err
}*/
