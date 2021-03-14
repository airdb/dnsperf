/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"airdb.io/airdb/dnsperf/dnslib"
	"github.com/spf13/cobra"
)

// ptrCmd represents the query command
var ptrCmd = &cobra.Command{
	Use:   "ptr",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		client := dnslib.PtrClient{DNSServer: ptrFlag.DNSServer}
		client.StressPtr()
	},
}

// ptrCmd represents the query command
var aRecordCmd = &cobra.Command{
	Use:   "a",
	Short: "A brief description of your command",
	Long: `A to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		client := dnslib.PtrClient{DNSServer: ptrFlag.DNSServer}
		client.StressRecordA()
	},
}


func init() {
	rootCmd.AddCommand(ptrCmd)
	rootCmd.AddCommand(aRecordCmd)

	ptrCmd.PersistentFlags().StringVarP(&ptrFlag.DNSServer, "dns-server", "", "8.8.8.8", "config file (default is $HOME/.dnstest.yaml)")
}

type PtrFlag struct {
	DNSServer string
}

var ptrFlag PtrFlag
