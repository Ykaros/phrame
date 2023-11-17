/*
Copyright © 2023 Ykaros
*/
package cmd

import (
	"fmt"
	"github.com/Ykaros/phrame/utils"
	"github.com/spf13/cobra"
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
			destinationPath, _ = cmd.Flags().GetString("output")

			// Check if the input path is a directory and output path is not specified
			if utils.IsDir(sourcePath) && destinationPath == "" {
				fmt.Print("Do you want to give a name to the output directory?")
				fmt.Scanln(&destinationPath)
			}
		}

		borderRatio, _ := cmd.Flags().GetFloat64("border")
		squareOption, _ := cmd.Flags().GetBool("square")
		colorOption, _ := cmd.Flags().GetString("color")

		c, err := utils.ParseHexColorFast(colorOption)
		if err != nil {
			fmt.Printf("Invalid color format: %v\n", err)
		}
		err = utils.AddFrames(sourcePath, destinationPath, borderRatio, squareOption, c)
		if err != nil {
			fmt.Printf("Error adding frames: %v\n", err)
		}
	},
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	rootCmd.PersistentFlags().StringP("input", "i", "", "Original image(s) location")
	rootCmd.PersistentFlags().StringP("output", "o", "", "Output directory for images with frames")
	rootCmd.PersistentFlags().StringP("color", "c", "0", "Frame color options")
	rootCmd.Flags().Float64P("border", "b", 0.1, "Border ratio for the frame")
	rootCmd.Flags().BoolP("square", "s", false, "Whether the frame is square or not")
}
