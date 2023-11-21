/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/Ykaros/phrame/utils"
	"github.com/spf13/cobra"
)

var grids string

// cutCmd represents the cut command
var cutCmd = &cobra.Command{
	Use:     "cut",
	Short:   "Cut the image into grids",
	Example: "phrame cut [image_path] -g [4 or 9]",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 1 {
			sourcePath = args[0]
		} else if len(args) > 1 {
			sourcePath = args[0]
			grids = args[1]
		}
		err := utils.Cut(sourcePath, grids)
		if err != nil {
			fmt.Printf("Error cutting image: %v\n", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(cutCmd)
	rootCmd.PersistentFlags().StringVarP(&grids, "grid", "g", "4", "The number of grids to cut the image into (4 or 9)")
}
