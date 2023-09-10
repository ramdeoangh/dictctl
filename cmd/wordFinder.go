/*
Copyright © 2023 ramdeo.angh@gmail.com
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var inputword string

var wordFinderCmd = &cobra.Command{
	Use:   "wordfinder",
	Short: "wordfinder - Search the any vocab you want",
	Long: `A word finder command is a CLI tool to find any vocabs from Merriam Webster
dictionary canyon. For example:

wordfinder --help`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		InfoLogger.Println("arguments", fmt.Sprint(args))
		InfoLogger.Println("output", filterResult(args[0]))
		fmt.Println(filterResult(args[0]))
	},
}

func init() {
	rootCmd.AddCommand(wordFinderCmd)

	wordFinderCmd.Flags().StringVarP(&inputword, "inputword", "w", "exercise", "Get meaning of search word")
}
