package rpc

import (
	"fmt"
	"io"
	"time"

	"github.com/asteris-llc/benchy/rpc/pb"
	"github.com/influxdata/influxdb/client/v2"
	"github.com/pkg/errors"
)

type ingester struct {
	Verifier *TokenVerifier

	// database stuff
	InfluxConfig client.HTTPConfig
	Database     string
}

func (i *ingester) AddBenchmark(stream pb.Ingester_AddBenchmarkServer) error {
	var count uint64

	influx, err := client.NewHTTPClient(i.InfluxConfig)
	if err != nil {
		return errors.Wrap(err, "could not get influx client")
	}

	batch, err := client.NewBatchPoints(client.BatchPointsConfig{
		Precision: "s",
		Database:  i.Database,
	})
	if err != nil {
		return errors.Wrap(err, "could not start batch")
	}

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

		// add our point to the batch
		if bench := benchmark.GetGoTestBench(); bench != nil {
			point, err := client.NewPoint(
				"benchmark",
				map[string]string{
					"project": benchmark.Project,
					"name":    bench.Name,
					"package": bench.Package,
				},
				map[string]interface{}{
					"n":                 bench.N,
					"nsPerOp":           bench.NsPerOp,
					"allocedBytesPerOp": bench.AllocedBytesPerOp,
					"allocsPerOp":       bench.AllocsPerOp,
					"mbPerS":            bench.MbPerS,
					"measured":          bench.Measured,
				},
				time.Unix(benchmark.Timestamp, 0),
			)
			if err != nil {
				if inerr := stream.SendAndClose(&pb.WriteStatus{
					Status: &pb.WriteStatus_Error{Error: err.Error()},
				}); inerr != nil {
					return errors.Wrap(err, "sending final status")
				}
				return errors.Wrap(err, "creating point")
			}
			batch.AddPoint(point)
		}

		count++
	}

	if err := influx.Write(batch); err != nil {
		if inerr := stream.SendAndClose(&pb.WriteStatus{
			Status: &pb.WriteStatus_Error{Error: err.Error()},
		}); inerr != nil {
			return errors.Wrap(err, "sending final status")
		}

		return errors.Wrap(err, "wrinting data points")
	}
	fmt.Println("")

	err = stream.SendAndClose(&pb.WriteStatus{
		Status: &pb.WriteStatus_Stats_{Stats: &pb.WriteStatus_Stats{
			Written: count,
		}},
	})
	if err != nil {
		return errors.Wrap(err, "sending final status")
	}

	return nil
}
