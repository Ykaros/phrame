/*
Copyright © 2023 Ykaros
*/
package cmd

import (
	"fmt"
	"github.com/Ykaros/phrame/utils"
	"github.com/spf13/cobra"
	"os"
	"strconv"
)

//var formatOption int

var rootCmd = &cobra.Command{
	Use:   "phrame",
	Short: "A CLI tool written in Go to add a frame to photo(s)",
	Run: func(cmd *cobra.Command, args []string) {

		sourcePath, _ := cmd.Flags().GetString("source")
		destinationPath, _ := cmd.Flags().GetString("destination")
		borderRatio, _ := cmd.Flags().GetFloat64("borderRatio")
		//colorOption, _ := cmd.Flags().GetBool("color")

		for i, arg := range args {
			switch i {
			case 0:
				sourcePath = arg
			case 1:
				destinationPath = arg
			case 2:
				borderRatio, _ = strconv.ParseFloat(arg, 64)
				//case 3:
				//	colorOption, _ = strconv.ParseBool(arg)
			}
		}

		// Check if the source exists and source type
		fileInfo, err := os.Stat(sourcePath)
		if err != nil {
			fmt.Printf("Error accessing file(s) at %s: %v\n", sourcePath, err)
			return
		}

		// Determine if the source is a directory or a file
		if fileInfo.IsDir() {
			err := utils.AddFrames(sourcePath, destinationPath, borderRatio)
			if err != nil {
				fmt.Printf("Error adding frames: %v\n", err)
				os.Exit(1)
			}
		} else {
			err := utils.AddFrame(sourcePath, destinationPath, borderRatio)
			if err != nil {
				fmt.Printf("Error adding frames: %v\n", err)
				os.Exit(1)
			}
		}
	},
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	rootCmd.Flags().StringP("source", "s", "", "Original image(s) location")
	rootCmd.Flags().StringP("destination", "d", "", "Output directory for images with frames")
	rootCmd.Flags().Float64P("borderRatio", "b", 0.1, "Border ratio for the frame")
	//rootCmd.Flags().BoolP("color", "c", false, "Whether the frame is colored or not")
	// rootCmd.Flags().IntVarP(&formatOption, "format", "f", 1, "Whether the frame is colored or not")
}
