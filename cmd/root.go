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

var rootCmd = &cobra.Command{
	Use:   "phrame",
	Short: "A CLI tool written in Go to add a frame to photo(s)",
	// At least a source is required
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		sourcePath := args[0]

		var destinationPath string
		// Output path is optional and can be passed as a flag or as an argument
		if len(args) > 1 {
			destinationPath = args[1]
		} else {
			destinationPath, _ = cmd.Flags().GetString("destination")
		}

		borderRatio, _ := cmd.Flags().GetFloat64("borderRatio")
		//colorOption, _ := cmd.Flags().GetBool("color")

		err := utils.AddFrames(sourcePath, destinationPath, borderRatio)
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
	rootCmd.Flags().StringP("source", "s", "", "Original image(s) location")
	rootCmd.Flags().StringP("destination", "d", "", "Output directory for images with frames")
	rootCmd.Flags().Float64P("borderRatio", "b", 0.1, "Border ratio for the frame")
	//rootCmd.Flags().BoolP("color", "c", false, "Whether the frame is colored or not")
	// rootCmd.Flags().IntVarP(&formatOption, "format", "f", 1, "Whether the frame is colored or not")
}
