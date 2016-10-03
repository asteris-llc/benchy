package rpc

import (
	"fmt"
	"io"

	"github.com/asteris-llc/benchy/rpc/pb"
	"github.com/pkg/errors"
)

type ingester struct{}

func (*ingester) AddBenchmark(stream pb.Ingester_AddBenchmarkServer) error {
	var count uint64

	for {
		benchmark, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return errors.Wrap(err, "receiving benchmark")
		}

		fmt.Printf("%+v", benchmark)
		count++
	}

	err := stream.SendAndClose(&pb.WriteStatus{
		Status: &pb.WriteStatus_Stats_{&pb.WriteStatus_Stats{
			Written: count,
		}},
	})
	if err != nil {
		return errors.Wrap(err, "sending final status")
	}

	return nil
}
