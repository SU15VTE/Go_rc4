/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// decryptCmd represents the decrypt command
var decryptCmd = &cobra.Command{
	Use:   "decrypt",
	Short: "rc4 decrypt",
	Long:  `nil`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("decrypt called")
		fmt.Println("de called")
		fmt.Printf("Key:	%s\n", key)
		fmt.Printf("text:	%s\n", text)
		result := RC4_decrypt(String_to_byte(key), String_to_byte(text))
		fmt.Printf("plaintext: %s\n", string(result))
	},
}

func init() {
	rootCmd.AddCommand(decryptCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// decryptCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// decryptCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
