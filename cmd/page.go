/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"log"
	"strconv"

	"github.com/spf13/cobra"
)

// pageCmd represents the page command
var pageCmd = &cobra.Command{
	Use:   "page",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			log.Fatal("You didn't specify the additional arguments\nUse the --help flag for comprehensive help on how to use this tool")
		}
		numberStringConv, err := strconv.Atoi(args[0])
		if err != nil {
			log.Fatalln("Please specify a number not a string")
		}
		if numberStringConv > 30 {
			log.Fatalln()
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
