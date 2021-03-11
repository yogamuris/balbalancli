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
	"github.com/spf13/cobra"
	"github.com/yogamuris/balbalancli/cmd/handler"
)

var leagueFlag string

// standingCmd represents the standing command
var standingCmd = &cobra.Command{
	Use:   "standing",
	Short: "Get standing of the League",
	Long: `Get standing of the League. For example:
	balbalan standing -L PL.`,

	RunE: func(cmd *cobra.Command, args []string) error {
		if err := handler.GetStanding(leagueFlag); err != nil {
			return err
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(standingCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// standingCmd.PersistentFlags().String("league", "", "Standing League")

	standingCmd.Flags().StringVarP(&leagueFlag, "league", "l", "", "Standing League")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// standingCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
