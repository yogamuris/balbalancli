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
	"github.com/yogamuris/balbalancli/handler"
)

var leagueFlag string

// standingCmd represents the standing command
var standingCmd = &cobra.Command{
	Use:   "standing",
	Short: "Get standing of the League",
	Long: `Get standing of the League.
	Available league :
	- CL  : UEFA Champions League
	- PL  : English Premier League
	- PD  : Primera Division
	- SA  : Serie A
	- BL1 : Bundesliga
	- FL1 : League 1
	- DED : Eredivisie
	- PPL : Primeira Liga
	`,

	Run: func(cmd *cobra.Command, args []string) {
		if leagueFlag == "" {
			leagueFlag = "All"
		}
		handler.GetStanding(leagueFlag)
	},
}

func init() {
	rootCmd.AddCommand(standingCmd)

	standingCmd.Flags().StringVarP(&leagueFlag, "league", "l", "All", "Get the league standing.")

	standingCmd.Example = `
	balbalan standing -l PL
	balbalan standing --league PL			
	`
}
