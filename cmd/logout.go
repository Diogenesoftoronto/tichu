/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func LogoutCmd() *cobra.Command {
	// logoutCmd represents the logout command
	var logoutCmd = &cobra.Command{
		Use:   "logout",
		Short: "logout of your account",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("logout called")
		},
	}
	return logoutCmd
}
func init() {
	rootCmd.AddCommand(LogoutCmd())

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// logoutCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// logoutCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
