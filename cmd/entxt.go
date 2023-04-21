/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// entxtCmd represents the entxt command
var entxtCmd = &cobra.Command{
	Use:   "entxt",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		txt, err := ReadFile(text)
		if err != nil {
			fmt.Println("ReadFile ERROR!")
		} else {
			result := RC4_encrypt(String_to_byte(key), txt)
			fmt.Println(string(result))
		}
	},
}

func init() {
	rootCmd.AddCommand(entxtCmd)
}
