/*
Copyright © 2021 DANIEL PORTO <daniel.porto@gmail.com>

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
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/spf13/cobra"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

var cfgFile string
var keys []string
var keys_str string
var key string
var wallets []string
var wallets_str string

var origin int
var target int

var host string
var port string
var disable_events bool

var contract string   // used as bank middleware for the transaction
var money int64         //amount for transfer
var trxgaslimit uint64 //used for suggesting the fee for the transaction
var duration int      //used for duration of experiment
var threads int       // used for workload
var verbosity string  //set log verbosity
var client_id string  // client identifier to match transaction response
var debug bool

var latency_factor_unity string
var latency_factor int64

func Log(message string, a ...interface{}) {
	var m string
	m = fmt.Sprintf(message, a...)

	now := time.Now()
	location, _ := time.LoadLocation("Europe/Lisbon")
	now_ts_ms := now.UnixNano() / 1_000_000
	now_date := now.In(location).Format("2006-01-02 15:04:05")
	fmt.Printf("%v|%v [BankTransferClient]: %v\n", now_date, now_ts_ms, m)
}

func LogWtag(tag, message string, a ...interface{}) {
	var m string
	m = fmt.Sprintf(message, a...)

	now := time.Now()
	location, _ := time.LoadLocation("Europe/Lisbon")
	now_ts_ms := now.UnixNano() / 1_000_000
	now_date := now.In(location).Format("2006-01-02 15:04:05")
	fmt.Printf("%v|%v [BankTransferClient]: %v - %v\n", now_date, now_ts_ms, tag, m)
}

func LogFatal(message string, a ...interface{}) {
	LogWtag("Fatal Error", message, a...)
	os.Exit(-1)
}
func LogError(message string, a ...interface{}) {
	LogWtag("Error", message, a...)
}

func LogDebug(message string, a ...interface{}) {
	if !debug {
		return
	}
	LogWtag("Debug", message, a...)
}

// rootCmd represents the base command when called without any subcommands
// 1 eth in (wei)ght = 1 +18 zeros.
// 0.000001 eth = 1 + 12 zeros wei
var rootCmd = &cobra.Command{
	Use:   "bank",
	Short: "A tool to transfer funds between accounts",
	Long:  `A golang client that interact with a quorum network to transfer funds between accounts in the blockchain.`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		if strings.TrimSpace(strings.ToUpper(verbosity)) == "DEBUG" {
			debug = true
			LogDebug("Debug mode enabled")

		}

		switch latency_factor_unity {
		case "microseconds":
			latency_factor = 1_000
			Log("Latency for request/reply are computed in  microseconds ")
		case "milliseconds":
			latency_factor = 1_000_000
			Log("Latency for request/reply are computed in  milliseconds ")
		}
	},
}

//Note:
//// 1 eth in (wei)ght = 1 +18 zeros.
//// 0.000001 eth = 1 + 12 zeros wei.
//For example: transfer 0.000001 wei to another account
//bank workload -o transfer --host "192.168.10.166" --port 23000 \
//-c 10
//--keys "1be3b50b31734be48452c29d714941ba165ef0cbf3ccea8ca16c45e3d8d45fb0,9bdd6a2e7cc1ca4a4019029df3834d2633ea6e14034d6dcc3b944396fe13a08b" \
//--wallets "0xed9d02e382b34818e88b88a309c7fe71e65f419d,0xca843569e3427144cead5e4d5999a3d0ccf92b8e" \
//--money 1000000000000
//--origin 0  --target 1

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	//rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.counter.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	//rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	rootCmd.PersistentFlags().StringVarP(&host, "host", "a", "", "hostname or ip address of a node in a blockchain network to submit the operation")
	rootCmd.MarkPersistentFlagRequired("host")
	rootCmd.PersistentFlags().StringVarP(&port, "port", "p", "", "url of the server to connect to")
	rootCmd.MarkPersistentFlagRequired("port")

	rootCmd.PersistentFlags().StringVarP(&verbosity, "verbosity", "v", "info", "Log level (trace, debug, info, warn, error, fatal, panic")
	rootCmd.PersistentFlags().Uint64VarP(&trxgaslimit, "gaslimit", "g", 3000000, "Gas limit for the transaction")
	rootCmd.PersistentFlags().BoolVarP(&disable_events, "events", "e", false, "url of the server to connect to")
	rootCmd.PersistentFlags().StringVarP(&latency_factor_unity, "req_latency_unity", "", "microseconds", "latency between the request and reply: microseconds/milliseconds")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".bank" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".bank")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}
