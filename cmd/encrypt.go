/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// encryptCmd represents the encrypt command
var encryptCmd = &cobra.Command{
	Use:   "encrypt",
	Short: "rc4 encrypt",
	Long:  `nil`,
	Run: func(cmd *cobra.Command, args []string) {
		result := RC4_encrypt(String_to_byte(key), String_to_byte(text))
		fmt.Printf("ciphertext: %s\n", string(result))
	},
}

func init() {
	rootCmd.AddCommand(encryptCmd)
}
