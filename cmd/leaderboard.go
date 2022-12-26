// /*
// Copyright © 2022 NAME HERE <EMAIL ADDRESS>

// */
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// // leaderboardCmd represents the leaderboard command
func LeaderboardCmd() *cobra.Command {
leaderboardCmd := &cobra.Command {
	Use:   "leaderboard [flags] [args]",
	Short: "Figure out who is winning",
	Long: `This command will show you the current leaderboard
	You can view the global, local, or personal leaderboard
		
		tichu leaderboard --global
		
		tichu leaderboard --local <server_id> 
		
		tichu leaderboard --personal <player_id>

	You can also view the leaderboard for a specific game.
		
		tichu leaderboard --game <game_id>
	You can also view the current rank of a specific player in game.
	
		tichu leaderboard --game <game_id> --personal <player_id>
	`,
	RunE: func(cmd *cobra.Command, args []string) (error) {
		fmt.Println("leaderboard called")
		fmt.Fprintf(cmd.OutOrStdout(), "")
		return nil
		},
	}
	return leaderboardCmd
}

func init() {
	leaderboardCmd := LeaderboardCmd()
	rootCmd.AddCommand(leaderboardCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// leaderboardCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// leaderboardCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
