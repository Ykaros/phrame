/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/Ykaros/phrame/utils"
	"github.com/spf13/cobra"
)

// signCmd represents the sign command
var signCmd = &cobra.Command{
	Use:   "sign",
	Short: "Sign the framed image",
	Run: func(cmd *cobra.Command, args []string) {
		signature, _ := cmd.Flags().GetString("signature")
		fontSize, _ := cmd.Flags().GetInt("fontSize")
		fontColor, _ := cmd.Flags().GetString("fontColor")

		for signature == "" {
			fmt.Print("You need to give a signature: ")
			fmt.Scanln(&signature)
		}
		c, err := utils.Hex2Color(frameColor)
		if err != nil {
			fmt.Printf("Invalid color format: %v\n", err)
		}

		fc, err := utils.Hex2Color(fontColor)
		if err != nil {
			fmt.Printf("Invalid color format: %v\n", err)
		}
		fmt.Println(sourcePath, destinationPath, borderRatio, squareOption, c, fc, signature, fontSize)
		err = utils.AddFrames(sourcePath, destinationPath, borderRatio, squareOption, c, fc, signature, fontSize)
		if err != nil {
			fmt.Printf("Error adding frames: %v\n", err)
		}

	},
}

func init() {
	rootCmd.AddCommand(signCmd)
	rootCmd.PersistentFlags().StringP("signature", "s", "", "Signature to add to the image")
	rootCmd.PersistentFlags().IntP("fontSize", "x", 30, "Font size options")
	rootCmd.PersistentFlags().StringP("fontColor", "y", "1", "Font color options")
}
