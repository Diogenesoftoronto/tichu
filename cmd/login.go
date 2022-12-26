// /*
// Copyright © 2022 NAME HERE <EMAIL ADDRESS>

// */
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// loginCmd represents the login command
func LoginCmd() *cobra.Command {
	var loginCmd = &cobra.Command{
		Use:   "login",
		Short: "login to your account",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("login called")
		},
	}
	return loginCmd
}

func init() {
	rootCmd.AddCommand(LoginCmd())

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// loginCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// loginCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
