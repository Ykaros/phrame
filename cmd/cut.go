/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/Ykaros/phrame/utils"
	"github.com/spf13/cobra"
)

// cutCmd represents the cut command
var cutCmd = &cobra.Command{
	Use:   "cut",
	Short: "Cut the image into grids",
	Run: func(cmd *cobra.Command, args []string) {
		grids, _ := cmd.Flags().GetInt("grid")
		err := utils.Cut(sourcePath, grids)
		if err != nil {
			fmt.Printf("Error cutting image: %v\n", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(cutCmd)
	rootCmd.Flags().IntP("grid", "g", 4, "The number of grids to cut the image into")
}
