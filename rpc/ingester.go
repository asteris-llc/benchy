package rpc

import (
	"fmt"
	"io"

	"github.com/asteris-llc/benchy/rpc/pb"
	"github.com/pkg/errors"
)

type ingester struct {
	Verifier *TokenVerifier
}

func (i *ingester) AddBenchmark(stream pb.Ingester_AddBenchmarkServer) error {
	var count uint64

	for {
		benchmark, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return errors.Wrap(err, "receiving benchmark")
		}

		// verify that the benchmark's project matches the token
		if err := i.Verifier.VerifyForProjectInContext(stream.Context(), benchmark.Project); err != nil {
			err := stream.SendAndClose(&pb.WriteStatus{
				Status: &pb.WriteStatus_Error{Error: err.Error()},
			})
			if err != nil {
				return errors.Wrap(err, "sending final status")
			}
			return nil
		}

		fmt.Printf("%+v\n", benchmark)
		count++
	}
	fmt.Println("")

	err := stream.SendAndClose(&pb.WriteStatus{
		Status: &pb.WriteStatus_Stats_{Stats: &pb.WriteStatus_Stats{
			Written: count,
		}},
	})
	if err != nil {
		return errors.Wrap(err, "sending final status")
	}

	return nil
}
