/*
Copyright © 2022 NAME HERE <ukejegoodness599@gmail.com>

*/

package cmd

import (
	"HackerGo/scraper"
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"strconv"
)

// pageCmd represents the page command
var pageCmd = &cobra.Command{
	Use:   "page",
	Short: "The page command scrapes the page you specify",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 2 {
			log.Fatal("You didn't specify the additional arguments\nUse the --help flag for comprehensive help on how to use this tool")
		}
		firstArg, err := strconv.Atoi(args[0])
		secondArg := args[1]
		if err != nil {
			log.Fatal("You didn't specify the additional arguments\nUse the --help flag for comprehensive help on how to use this tool")
		}

		if err != nil && firstArg >= 30 {
			log.Fatalln("Please specify a number as the first argument not a string remember it has to be less than 30")
		}
		if secondArg == "all" {
			fmt.Println("Fetching you those buzzing tech gists; This might take a while...")
			scraper.AllFromPage(firstArg)
		} else {
			fmt.Println("Fetching you those buzzing tech gists; This might take a while...")
			secondArg, _ := strconv.Atoi(args[1])
			scraper.NumberOfNews(firstArg, secondArg)
		}

	},
}

func init() {
	rootCmd.AddCommand(pageCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// pageCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// pageCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
