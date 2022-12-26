/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	tichu "github.com/Diogenesoftoronto/tichu/tui"
	"github.com/spf13/cobra"
)

// playCmd represents the play command

func PlayCmd() *cobra.Command {
	var playCmd = &cobra.Command{
		Use:   "play",
		Short: "use this command to play the game, if you havent figured that out yet",
		Long: `Add players to the game, and then play the game.
		You can add players by using the add command.
		You can play the game by using the play command.
		You can remove players by using the remove command.
		You can list players by using the list command.
		You can quit the game by using the quit command.
		You can get help by using the help command.`,
		Run: func(cmd *cobra.Command, args []string){
			fmt.Println("play called")
			tichu.RunGame()
		},
	}
	return playCmd
}

func init() {
	rootCmd.AddCommand(PlayCmd())

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// playCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// playCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
