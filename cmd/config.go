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

	"github.com/fatih/color"

	"github.com/spf13/cobra"
	"github.com/yogamuris/balbalancli/handler"
)

var tokenFlag string

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Setting the BalbalanCLI Configuration.",
	Long:  `Setting the BalbalanCLI Configuration.`,
	Run: func(cmd *cobra.Command, args []string) {

		if tokenFlag == "" {
			color.Red(fmt.Sprintf("Error: flag required. \nExample : \n%s", cmd.Example))
		} else {
			handler.SetToken(tokenFlag)
		}
	},
}

func init() {
	rootCmd.AddCommand(configCmd)
	configCmd.Flags().StringVarP(&tokenFlag, "token", "t", "", "Set API Token")
	configCmd.Example = `
	balbalan config -t YOUR_TOKEN_VALUE
	balbalan config --token YOUR_TOKEN_VALUE			
	`
}
