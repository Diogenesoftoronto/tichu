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
	var rootCmd = &cobra.Command{
		Use:   "tichu [command] [flags]",
		Short: "A card game for 4 players.",
		Long: `Tichu is played with a standard 52-card deck, plus 4 jokers.

	Each player is dealt 14 cards, and 
	the remaining cards are placed face down in the middle of the table.

	The object of the game is to win tricks containing cards of higher 
	value than those played by the other players.
	Each trick is won by the player who played the highest-ranking card of the suit led.

	Players may also play special cards called Tichus or Grand Tichus 
	to win tricks without playing the highest card.

	At the end of the game, the player with the most points wins.

	Points are awarded for winning tricks and for 
	holding certain cards in hand at the end of the game.`,
		// Uncomment the following line if your bare application
		// has an action associated with it:
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(`Welcome to Tichu! 
Run 'tichu play' to begin, or 'tichu help' to learn more.`)
		},
	}
	return rootCmd
}

var rootCmd = RootCmd()

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the RootCmd().
func Execute() {
	err := rootCmd.Execute()
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
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
