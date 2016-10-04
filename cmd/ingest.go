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
	"time"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// ingestCmd represents the ingest command
var ingestCmd = &cobra.Command{
	Use:   "ingest",
	Short: "ingest a metric",

	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		if err := RootCmd.PersistentPreRunE(cmd, args); err != nil {
			return err
		}

		switch viper.GetString("timestamp") {
		case "":
			viper.Set("timestamp-unix", time.Now().Unix())

		default:
			parsed, err := time.Parse(time.RFC3339, viper.GetString("timestamp"))
			if err != nil {
				return errors.Wrap(err, "could not parse timestamp, expecting RFC3339 value")
			}

			viper.Set("timestamp-unix", parsed.Unix())
		}

		return nil
	},
}

func init() {
	RootCmd.AddCommand(ingestCmd)

	ingestCmd.PersistentFlags().StringP("timestamp", "t", "", "timestamp for this metric (default now)")
	ingestCmd.PersistentFlags().String("rpc-addr", "localhost:8080", "address of ingestion RPC server")
	ingestCmd.PersistentFlags().String("rpc-token", "", "RPC token")
}
