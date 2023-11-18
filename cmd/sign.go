/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// signCmd represents the sign command
var signCmd = &cobra.Command{
	Use:   "sign",
	Short: "Sign the framed image",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("sign called")
	},
}

func init() {
	rootCmd.AddCommand(signCmd)
	rootCmd.Flags().StringP("signature", "s", "", "Signature to add to the image")
	rootCmd.Flags().IntP("fontSize", "fs", 18, "Font size options")
	rootCmd.Flags().StringP("fontColor", "fc", "0", "Font color options")
}
