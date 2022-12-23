/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// leaderboardCmd represents the leaderboard command
func LeaderboardCmd() *cobra.Command {
leaderboardCmd := &cobra.Command {
	Use:   "leaderboard",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	RunE: func(cmd *cobra.Command, args []string) (error) {
		fmt.Println("leaderboard called")
		fmt.Fprintf(cmd.OutOrStdout(), "")
		return nil
		},
	}
	return leaderboardCmd
}

func init() {
	RootCmd().AddCommand(LeaderboardCmd())

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// leaderboardCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// leaderboardCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
