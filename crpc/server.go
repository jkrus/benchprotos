package crpc

import (
	"log"
	"net"

	"github.com/evenlab/go-kit/context"
	"zombiezen.com/go/capnproto2/rpc"

	cfg "benchprotos/config"
	"benchprotos/crpc/cpb"
)

type capnservice struct {
	quote cpb.Quote

	close    chan struct{}
	context  context.Context
	conn     net.Conn
	listener net.Listener
	options  rpc.ConnOption
}

var (
	// Make sure service implements Service interface.
	_ Service = (*capnservice)(nil)
)

func NewService() Service {
	l, err := net.Listen("tcp", cfg.AppDefaultHost+":"+cfg.AppDefaultPortCRPC)
	if err != nil {
		log.Fatalf("crpc listener: %v", err)
	}

	return &capnservice{
		listener: l,
		//		options: &rpc.Options{},
		//		policy: &crpc.Policy{},
	}
}

// Start starts crpc service.
func Start(ctx context.Context, s Service) error {
	log.Println("Start crpc service...")
	s.Serve(ctx)

	return nil
}

// Serve accepts incoming connections.
func (s *capnservice) Serve(ctx context.Context) {
	s.context = ctx

	go func() {
		select {
		case <-s.close:
			break
		case <-s.context.Done():
			break
		}

		log.Println("Stop crpc service...")
		if err := s.Stop(); err != nil {
			log.Println(err)
		} else {
			log.Println("crpc service stop complete")
		}
	}()

	log.Printf("crpc service start complete on port %v\n", cfg.AppDefaultPortCRPC)

	for {
		conn, err := s.listener.Accept()
		if err != nil {
			log.Fatal(err)
		}

		go s.serve(conn)
	}

}

// Stop stops grpc service.
func (s *capnservice) Stop() error {
	err := s.listener.Close()
	if err != nil {
		log.Println(err)
	}

	return err
}

// serve accepts incoming connections on the listener.
func (s *capnservice) serve(c net.Conn) {
	go func() {
		main := cpb.QuotesUpdate_ServerToClient(s)
		conn := rpc.NewConn(rpc.StreamTransport(c), rpc.MainInterface(main.Client))
		err := conn.Wait()
		if err != nil {
			log.Fatalf("crpc service: %v", err)
		}
	}()
}
