package rpc

import (
	"github.com/asteris-llc/benchy/rpc/pb"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
)

// NewIngesterClient gets an ingester client for a given address
func NewIngesterClient(addr string) (pb.IngesterClient, error) {
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		return nil, errors.Wrapf(err, "dialing %s failed", addr)
	}

	return pb.NewIngesterClient(conn), nil
}
