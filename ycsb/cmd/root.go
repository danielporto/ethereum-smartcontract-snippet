// Package cmd
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

	log "github.com/sirupsen/logrus"

	"github.com/spf13/cobra"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

var cfgFile string
var key string
var host string
var port string
var disable_events bool

var contract string  
var trxgaslimit int32 //used for suggesting the fee for the transaction
var duration int //used for duration of experiment
var threads int      // used for workload
var verbosity string //set log verbosity

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "ycsb",
	Short: "A simple smartcontract to maintain a ycsb key value store",
	Long: `A golang client that interact with a quorum network to deploy a smartcontract
that maintains a key value store. It installs, change and read the KV store from the blockchain. 
For example:
ycsb deploy --host 192.168.10.166 --port 22000 \
   			--key "1be3b50b31734be48452c29d714941ba165ef0cbf3ccea8ca16c45e3d8d45fb0"`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		lvl, err := log.ParseLevel(verbosity)
		if err != nil {
			log.Fatal("Invalid log level", err)
		}
		log.SetLevel(lvl)
		log.Debugf("Debug mode enabled")
	},
}

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

	rootCmd.PersistentFlags().StringVarP(&host, "host", "", "", "hostname or ip address of a node in a blockchain network to submit the operation")
	rootCmd.MarkPersistentFlagRequired("host")
	rootCmd.PersistentFlags().StringVarP(&port, "port", "", "", "url of the server to connect to")
	rootCmd.MarkPersistentFlagRequired("port")

	rootCmd.PersistentFlags().StringVarP(&verbosity, "verbosity", "", "info", "Log level (trace, debug, info, warn, error, fatal, panic")
	rootCmd.PersistentFlags().Int32VarP(&trxgaslimit, "gaslimit", "", 3000000, "Gas limit for the transaction")
	rootCmd.PersistentFlags().BoolVarP(&disable_events, "events", "", false, "url of the server to connect to")
	rootCmd.PersistentFlags().StringVarP(&key, "key", "", "", "key of a network node")
	rootCmd.MarkPersistentFlagRequired("key")
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

		// Search config in home directory with name ".ycsb" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".ycsb")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}
