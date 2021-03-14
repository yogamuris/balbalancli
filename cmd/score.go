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

var league string

// scoreCmd represents the score command
var scoreCmd = &cobra.Command{
	Use:   "score",
	Short: "Show the latest score.",
	Long: `Show the latest score.
	Available league :
	- CL  : UEFA Champions League
	- PL  : English Premier League
	- PD  : Primera Division
	- SA  : Serie A
	- BL1 : Bundesliga
	- FL1 : League 1
	- DED : Eredivisie
	- PPL : Primeira Liga`,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := handler.GetLatestScore(league)
		if err != nil {
			return err
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(scoreCmd)
	scoreCmd.Flags().StringVarP(&league, "league", "l", "All", "Get the league latest score.")

	scoreCmd.Example = `
	balbalan score -l PL
	balbalan score --league PL			
	`
}
