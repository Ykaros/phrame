/*
Copyright © 2023 Ykaros
*/
package cmd

import (
	"fmt"
	"github.com/Ykaros/phrame/utils"
	"github.com/spf13/cobra"
	"image/color"
)

var (
	sourcePath      string
	destinationPath string
	borderRatio     float64
	squareOption    bool
	frameColor      string
)

var rootCmd = &cobra.Command{
	Use:     "phrame",
	Short:   "A CLI tool written in Go to add a frame to photo(s)",
	Example: "phrame -q -i [input_path] -o [output_path] -r [border_ratio] -c [frame_color]",
	Run: func(cmd *cobra.Command, args []string) {

		// Check if the input path is a directory and output path is not specified
		if utils.IsDir(sourcePath) && destinationPath == "" {
			fmt.Print("Do you want to give a name to the output directory?")
			fmt.Scanln(&destinationPath)
		}

		c, err := utils.Hex2Color(frameColor)
		if err != nil {
			fmt.Printf("Invalid color format: %v\n", err)
		}
		err = utils.AddFrames(sourcePath, destinationPath, borderRatio, squareOption, c, color.RGBA{0, 0, 0, 255}, "", 18)
		if err != nil {
			fmt.Printf("Error adding frames: %v\n", err)
		}
	},
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&sourcePath, "input", "i", "", "Original image(s) location")
	rootCmd.PersistentFlags().StringVarP(&destinationPath, "output", "o", "", "Output directory for images with frames")
	rootCmd.PersistentFlags().StringVarP(&frameColor, "color", "c", "0", "Frame color options: 0 or 1 or #RRGGBB or #RRGGBBAA")
	rootCmd.PersistentFlags().Float64VarP(&borderRatio, "border", "r", 0.1, "Border ratio for the frame: [0, 1]")
	rootCmd.Flags().BoolVarP(&squareOption, "square", "q", false, "Whether the frame is square or not")
}
