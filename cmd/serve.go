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

	"github.com/Sirupsen/logrus"
	"github.com/asteris-llc/benchy/rpc"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		logrus.WithField("addr", viper.GetString("addr")).Info("listening")
		server := rpc.Server{
			Auth: map[string]string{
				viper.GetString("token"): viper.GetString("project"),
			},

			DatabaseAddr:     viper.GetString("influx-addr"),
			DatabaseUsername: viper.GetString("influx-username"),
			DatabasePassword: viper.GetString("influx-password"),
			DatabaseName:     viper.GetString("influx-database"),
		}

		if err := server.Listen(context.Background(), viper.GetString("addr")); err != nil {
			logrus.WithError(err).Fatal("could not serve")
		}
	},
}

func init() {
	RootCmd.AddCommand(serveCmd)

	serveCmd.Flags().String("addr", "localhost:8080", "address to serve")
	serveCmd.Flags().String("project", "", "project to ingest")
	serveCmd.Flags().String("token", "", "token to use for auth")

	// database stuff
	serveCmd.Flags().String("influx-addr", "http://localhost:8086", "influxdb address")
	serveCmd.Flags().String("influx-username", "", "influxdb username")
	serveCmd.Flags().String("influx-password", "", "influxdb password")
	serveCmd.Flags().String("influx-database", "benchy", "influxdb database")
}
