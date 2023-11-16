/*
Copyright © 2023 Ykaros
*/
package cmd

import (
	"fmt"
	"github.com/Ykaros/phrame/utils"
	"github.com/spf13/cobra"
	"os"
)

// TODO: 0) One image or dir; 1) More save options; 2) More color options
var rootCmd = &cobra.Command{
	Use:   "phrame",
	Short: "A CLI tool written in Go to add a frame to photo(s)",
	Run: func(cmd *cobra.Command, args []string) {
		sourceDir, destinationDir := args[0], args[1]
		borderRatio, _ := cmd.Flags().GetFloat64("borderRatio")

		err := utils.AddFrames(sourceDir, destinationDir, borderRatio)
		if err != nil {
			fmt.Printf("Error adding frames: %v\n", err)
			os.Exit(1)
		}
	},
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	rootCmd.Flags().StringP("source", "s", "", "Source directory containing images")
	rootCmd.Flags().StringP("destination", "d", "", "Output directory for images with frames")
	rootCmd.Flags().Float64P("borderRatio", "b", 0.1, "Border ratio for the frame")
}
