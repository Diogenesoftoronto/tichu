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
	leaderboardCmd := &cobra.Command{
		Use:   "leaderboard [commands] [flags] [args]",
		Short: "Figure out who is winning",
		Long: `This command will show you the current leaderboard
	You can view the global, local, or personal leaderboard
		
		leaderboard --global
		
		leaderboard --local <server_id> 
		
		leaderboard --personal <player_id>

	You can also view the leaderboard for a specific game.
		
		leaderboard game <game_id>
	You can also view the current rank of a specific player in game.
	
		leaderboard game <game_id> --personal <player_id>
	`,
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println("leaderboard called")
			fmt.Fprintf(cmd.OutOrStdout(), "")
			return nil
		},
	}
	return leaderboardCmd
}

func LeaderboardGameCmd() *cobra.Command {
	leaderboardGameCmd := &cobra.Command{
		Use: "game <game_id> args",
		Short: "View the leaderboard for a specific game",
		Long: `This command will show you the current leaderboard for a specific game.
		You can view the global, local, or personal leaderboard
			
			game --global <game_id> 
			
			game --local <game_id>  <server_id> 
			
			game --personal <game_id>  <player_id>
		`,
		Example: "leaderboard game <game_id> --personal <player_id>",
		Args: cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println("leaderboard game called")
			fmt.Fprintf(cmd.OutOrStdout(), "")
			return nil
		},
	}
	return leaderboardGameCmd
}

func init() {
	leaderboardCmd := LeaderboardCmd()
	gameCmd := LeaderboardGameCmd()
	leaderboardCmd.AddCommand(gameCmd)
	rootCmd.AddCommand(leaderboardCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// leaderboardCmd.PersistentFlags().String("foo", "", "A help for foo")
	
	leaderboardCmd.PersistentFlags().BoolP("global", "g", false, "view the global leaderboard")
	leaderboardCmd.PersistentFlags().BoolP("local", "s", false, "view the server leaderboard")
	leaderboardCmd.PersistentFlags().BoolP("personal", "p", false, "view a leaderboard for a specific player")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// leaderboardCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
