// Copyright © 2016 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"context"
	"log"
	"os"

	"io"

	"github.com/asteris-llc/benchy/parse/gobench"
	"github.com/asteris-llc/benchy/rpc"
	"github.com/asteris-llc/benchy/rpc/pb"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// gobenchCmd represents the gobench command
var gobenchCmd = &cobra.Command{
	Use:   "gobench",
	Short: "ingest a set of go benchmarks",
	Long: `Pipe the output of "go test -bench='.'" here to ingest. Output will
be displayed before ingestion.`,
	PreRunE: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return errors.New("need one argument, a project name")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()

		log.Printf("connecting to %q", viper.GetString("rpc-addr"))
		client, err := rpc.NewIngesterClient(viper.GetString("rpc-addr"))
		if err != nil {
			log.Fatal(errors.Wrap(err, "could not get ingester client"))
		}

		reader := io.TeeReader(os.Stdin, os.Stdout)
		benchmarks, err := gobench.Parse(reader)
		if err != nil {
			log.Fatal(errors.Wrap(err, "failed to read benchmarks"))
		}

		// send the benchmarks off to the server
		stream, err := client.AddBenchmark(ctx)
		if err != nil {
			errors.Wrap(err, "failed to get stream")
		}

		// send everything off
		for pkg, benchmarks := range benchmarks {
			for _, benchmark := range benchmarks {
				err := stream.Send(&pb.Benchmark{
					Project: args[0],
					Kind: &pb.Benchmark_GoTestBench_{GoTestBench: &pb.Benchmark_GoTestBench{
						Name:              benchmark.Name,
						Package:           pkg,
						N:                 uint64(benchmark.N),
						NsPerOp:           benchmark.NsPerOp,
						AllocedBytesPerOp: benchmark.AllocedBytesPerOp,
						AllocsPerOp:       benchmark.AllocsPerOp,
						MbPerS:            benchmark.MBPerS,
						Measured:          int64(benchmark.Measured),
					}},
				})
				if err != nil {
					log.Fatal(errors.Wrap(err, "failed to send benchmark"))
				}
			}
		}

		receipt, err := stream.CloseAndRecv()
		if err != nil {
			errors.Wrap(err, "closing failed")
		}

		if errResult := receipt.GetError(); errResult != "" {
			log.Fatalf("FAIL\tgot an error: %s", errResult)
		} else if result := receipt.GetStats(); result != nil {
			log.Printf("ok  \tsuccess, %d written", result.Written)
		}
	},
}

func init() {
	ingestCmd.AddCommand(gobenchCmd)
}
