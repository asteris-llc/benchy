package rpc

import (
	"context"
	"log"
	"net"
	"net/http"

	"golang.org/x/sync/errgroup"

	"github.com/asteris-llc/benchy/rpc/pb"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/pkg/errors"
	"github.com/soheilhy/cmux"
	"google.golang.org/grpc"
)

// NewGRPC constructs all servers and handlers
func NewGRPC() (*grpc.Server, error) {
	server := grpc.NewServer()

	pb.RegisterIngesterServer(server, new(ingester))

	return server, nil
}

// NewREST constructs a thingy for REST interface
func NewREST(ctx context.Context, addr string) (*http.Server, error) {
	mux := runtime.NewServeMux()
	err := pb.RegisterIngesterHandlerFromEndpoint(ctx, mux, addr, []grpc.DialOption{grpc.WithInsecure()})
	if err != nil {
		return nil, err
	}

	server := &http.Server{
		Handler: mux,
	}

	return server, nil
}

// Listen on the given address for all the server-y duties
func Listen(ctx context.Context, addr string) error {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		return errors.Wrap(err, "failed to listen")
	}

	// set up context and wait groups for cancelling all of this
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	wg, ctx := errgroup.WithContext(ctx)

	mux := cmux.New(lis)

	// start GRPC listener
	grpcSrv, err := NewGRPC()
	if err != nil {
		return errors.Wrap(err, "failed to create grpc server")
	}
	grpcLis := mux.Match(cmux.HTTP2HeaderField("content-type", "application/grpc"))
	wg.Go(func() error {
		log.Println("serving grpc")
		return grpcSrv.Serve(grpcLis)
	})

	// start REST listener
	restSrv, err := NewREST(ctx, addr)
	if err != nil {
		return errors.Wrap(err, "failed to create REST server")
	}
	restLis := mux.Match(cmux.HTTP1())
	wg.Go(func() error {
		log.Println("waiting for close")
		<-ctx.Done()
		log.Println("closing")
		return restLis.Close()
	})
	wg.Go(func() error {
		log.Println("serving http")
		return restSrv.Serve(restLis)
	})

	// start our cmux listener
	wg.Go(mux.Serve)

	// wait for all the listeners to return
	return wg.Wait()
}
