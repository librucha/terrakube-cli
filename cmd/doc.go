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

	"github.com/spf13/cobra"
)

var docLong = `
This command shows API doc for Terrakube CLI
`

var docCmd = &cobra.Command{
	Use:   "doc",
	Short: "get API doc",
	Long:  docLong,
	Run: func(cmd *cobra.Command, args []string) {
		getDoc()
	},
}

func init() {
	rootCmd.AddCommand(docCmd)
}

func getDoc() {
	client := newClient()

	resp, err := client.Doc.Get()

	if err != nil {
		fmt.Println(err)
		return
	}

	renderOutput(resp, output)
}
