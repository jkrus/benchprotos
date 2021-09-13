// Copyright © 2020-2021 The EVEN Solutions Developers Team

package grpc

import (
	"io"
	"log"
	"net"

	"github.com/evenlab/go-kit/context"
	"github.com/evenlab/go-kit/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	cfg "benchprotos/config"
	"benchprotos/grpc/proto/pb"
)

type (
	// service implements grpc service interface.
	service struct {
		pb.UnimplementedQuotesServiceServer
		pb.UnimplementedPingPongServer

		close    chan struct{}
		context  context.Context
		server   *grpc.Server
		listener net.Listener
	}
)

var (
	// Make sure service implements Service interface.
	_ Service = (*service)(nil)
)

// NewService returns grpc service interface.
func NewService() Service {
	l, err := net.Listen("tcp", cfg.AppDefaultHost+":"+cfg.AppDefaultPortGRPC)
	if err != nil {
		log.Fatalf("grpc listener: %v", err)
	}

	server := grpc.NewServer()
	reflection.Register(server)

	return &service{
		close:    make(chan struct{}),
		server:   server,
		listener: l,
	}
}

// Start starts grpc service.
func Start(ctx context.Context, s Service) error {
	log.Println("Start grpc service...")
	s.Serve(ctx)

	return nil
}

// Serve accepts incoming connections.
func (s *service) Serve(ctx context.Context) {
	s.context = ctx
	s.register()

	served := make(chan struct{})
	go s.serve(served)

	<-served
	log.Printf("grpc service start complete on port %v\n", cfg.AppDefaultPortGRPC)
}

// Stop stops grpc service.
func (s *service) Stop() error {
	err := s.listener.Close()
	if err != nil {
		log.Println(err)
	}

	s.server.Stop()

	return err
}

// register performs registrations of grpc service for services.
func (s *service) register() {
	// use alphabetical order only
	pb.RegisterQuotesServiceServer(s.server, s)
}

// serve accepts incoming connections on the listener.
func (s *service) serve(served chan struct{}) {
	go func() {
		select {
		case <-s.close:
			break
		case <-s.context.Done():
			break
		}

		log.Println("Stop grpc service...")
		if err := s.Stop(); err != nil {
			log.Println(err)
		} else {
			log.Println("grpc service stop complete")
		}
	}()

	served <- struct{}{}
	if err := s.server.Serve(s.listener); err != nil {
		if !errors.Is(err, io.EOF) {
			log.Fatalf("grpc service: %v", err)
		}
	}
}
