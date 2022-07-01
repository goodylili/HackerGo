/*
Copyright © 2022 NAME HERE <ukejegoodness599@gmail.com>

*/

package cmd

import (
	"Hacker-Go/scraper"
	"fmt"
	"log"
	"strconv"

	"github.com/spf13/cobra"
)

// randomCmd represents the random command
var randomCmd = &cobra.Command{
	Use:   "random",
	Short: "The random command would get you news from a random page",
	Long: `
	The random command would get you news from a random page.
	You can specify the number of news you want from the random page or get all the news.
	
	
	all = gets all the jobs from the page 
	int = specify a number of jobs (has to be below 30)
	
	`,

	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			log.Fatalln("You didn't specify the additional arguments\nUse the --help flag for comprehensive help on how to use this tool")
		}
		if args[0] == "all" {
			fmt.Println("Fetching you those buzzing tech gists; This might take a while...")
			scraper.AllFromPage(scraper.RandomNewsPage())
		} else {
			numberStringConv, err := strconv.Atoi(args[0])
			if err != nil {
				log.Fatalln("Please specify a number from 1 - 30")
			}
			if numberStringConv > 30 {
				log.Fatalln("Please specify a number from 1 - 30")
			} else {
				fmt.Println("Fetching you those buzzing tech gists; This might take a while...")
				scraper.NumberOfNews(scraper.RandomNewsPage(), numberStringConv)
			}

		}
		//fmt.Println("random called")
	},
}

func init() {
	rootCmd.AddCommand(randomCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// randomCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// randomCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
