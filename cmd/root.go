/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// RootCmd() represents the base command when called without any subcommands
func RootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "tichu - multiplayer card game",
		Short: "A card game for 4 players.",
		Long: `A card game for 4 players.
	tichu play begins a tichu game with 4 players.
	
	tichu cheer cheers you on to victory.
	
	tichu help lists all commands.
	
	tichu end ends the game.
	
	tichu leaderboard lists the players and the score.
	
	tichu message sends a message to the other players.`,
		// Uncomment the following line if your bare application
		// has an action associated with it:
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Fprintf(cmd.OutOrStdout(), "Welcome to Tichu!\n")
			return nil
		},
	}
	return cmd
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the RootCmd().
func Execute() {
	err := RootCmd().Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// RootCmd().PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.tichu.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	RootCmd().Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
