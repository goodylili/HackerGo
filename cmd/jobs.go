// Package cmd /*
package cmd

import (
	"HackerGo/scraper"
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"strconv"
)

// jobsCmd represents the jobs command
var jobsCmd = &cobra.Command{
	Use:   "jobs",
	Short: "Jobs fetches jobs from y combinators hacker news  website ||- arguments = (all, int)",
	Long: `
	The Job command fetches the most recent jobs from y combinators hacker news  website so you can focus on writing code while fueling your hustle.
	Only the first page is accessible for 'timely' reasons. You can visit Y combinators website for more jobs".
	
	You can use these arguments with the job subcommand
	
	all = gets all the jobs from the page and
	
	int = specify a number of jobs (has to be below 30)
	
	`,

	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			log.Fatal("You didn't specify the additional arguments\nUse the --help flag for comprehensive help on how to use this tool")
		}
		if args[0] == "all" {
			fmt.Println("Fetching you those gigs; This might take a while...")
			scraper.AllJobsFromPage()
		} else {
			numberStringConv, err := strconv.Atoi(args[0])
			if err != nil {
				log.Fatalln("Please specify a number from 1 - 30")
			}
			if numberStringConv > 30 {
				log.Fatalln("Please specify a number from 1 - 30")
			} else {
				fmt.Println("Fetching you those gigs; This might take a while...")
				scraper.NumberOfJobs(numberStringConv)
			}

		}

		//fmt.Println("jobs called")
	},
}

func init() {
	rootCmd.AddCommand(jobsCmd)

}
