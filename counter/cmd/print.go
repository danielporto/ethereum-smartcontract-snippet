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
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"

	"github.com/danielporto/ethereum-smartcontract-snippet/counter/contracts"
	"github.com/spf13/cobra"
)

// printCmd represents the print command
var printCmd = &cobra.Command{
	Use:   "print",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("print called")
		executePrint()
	},
}

func init() {
	runCmd.AddCommand(printCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// printCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// printCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func executePrint() {
		// https://goethereumbook.org/en/smart-contract-load/
		// connects
		client, err := ethclient.Dial("http://localhost:7545")
		if err != nil {
			log.Fatal(err)
		}

		// use the transaction id of the contract deployed to return a reference to the contract
		address := common.HexToAddress("0x50235A46F1401070613dA7154f1154b39A6c8686")
		instance, err := contracts.NewCounter(address, client)
		if err != nil {
			log.Fatal(err)
		}

		// read a value from the blockchain
		counter, err := instance.GetCount(nil)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Current counter value: %v\n", counter)

}