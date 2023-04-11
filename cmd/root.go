/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var key string
var text string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "Go_RC4",
	Short: "nil",
	Long: `
	encrypt --key "SU15VTE" --text "Hello World!"
	decrypt --key "SU15VTE" --text "4dc1dacdad813f8f50a9b9c7"`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&text, "text", "t", "", "A Text")
	rootCmd.PersistentFlags().StringVarP(&key, "key", "k", "", "The key to use for encryption/decryption")
}
