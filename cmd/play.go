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
		Use:   "play [command] [args]",
		Short: "use this command to play the game (if you havent figured that out yet)",
		Long: `Add players to the game, and then play the game.
		You can add players by using the add command.
		
		play add <player_id>...

		You can play the game by using the play command.
		
		play [--no-tui | -nt]

		Without arguments this command will generate a new game, launch the tui,
		and play the game with the current players. with the --no-tui or -nt flag
		it will generate a new game. 
		
		play game [<game_id>]
		if no game_id is provided. If a game_id is provided
		it will play the game with the current players. 
		You can also continue a game in the tui with this argument.

		
		You can remove players by using the remove command,
		keep in mind you need to be authorized to do this. 
		Also note that removing a player will end the game they are in.
		
		play remove <player_id>
		
		You can list players in your game by using the list command.
		You can also list the games you are in by using the list command.

		play list [-p | --players] <game_id> 
		play list [-g | --games] 

		You can quit the game by using the quit command.

		play quit`,
		ValidArgs: []string{"", "add", "remove", "list", "quit", "game"},
		Example:   "play add 1234567890",
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println("play called")
			tichu.RunGame()
			return nil
		},
	}
	return playCmd
}

func PlayGameCmd() *cobra.Command {
	var playGameCmd = &cobra.Command{
		Use:   "game [flags] [args]",
		Short: "Play the game",
		Long: `You can use this command to play the game.
		You can also continue a game in the tui with this argument.

		play game [<game_id>]
		if no game_id is provided. If a game_id is provided
		it will play the game with the current players.
		`,
		Example: "play game 1234567890",
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println("play game called")
			return nil
		},
	}
	return playGameCmd
}

func PlayListCmd() *cobra.Command {
	var playListCmd = &cobra.Command{
		Use:   "list [flags] [args]]",
		Short: "List players in the game",
		Long: `You can use this command to list players in the game.
		You can also list the games you are in.

		play list [-p | --players] <game_id> | [ -g | --games] 
		
		if the -p or --players flag is used, it will list the players in you must provide the game id.`,
		Example: "play list -p 1234567890",
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println("play list called")
			return nil
		},
	}
	return playListCmd
}

func PlayAddCmd() *cobra.Command {
	var playAddCmd = &cobra.Command{
		Use:   "add [args]",
		Short: "Add players to the game",
		Long: `You can use this command to add players to the game.
		You can add multiple players at once.
		
		play add <player_id>...`,
		Example: "play add 1234567890 0987654321",
		Args:    cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println("play add called")
			return nil
		},
	}
	return playAddCmd
}

func PlayRemoveCmd() *cobra.Command {
	var playRemoveCmd = &cobra.Command{
		Use:   "remove [args]",
		Short: "Remove players from the game",
		Long: `You can use this command to remove players from the game.
		You can remove multiple players at once.

		play remove <player_id>...`,
		Example: "play remove 1234567890 0987654321",
		Args:    cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println("play remove called")
			return nil
		},
	}
	return playRemoveCmd
}

func PlayQuitCmd() *cobra.Command {
	var playQuitCmd = &cobra.Command{
		Use:   "quit <game_id>",
		Short: "Quit a game",
		Long: `You can use this command to quit a game.
		You must provide the game id.

		play quit <game_id>...`,
		Example: "play quit <game_id>",
		Args:    cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println("play quit called")
			return nil
		},
	}
	return playQuitCmd
}
func init() {

	playCmd := PlayCmd()
	listCmd := PlayListCmd()
	gameCmd := PlayGameCmd()

	playCmd.AddCommand(listCmd)
	playCmd.AddCommand(gameCmd)
	playCmd.AddCommand(PlayAddCmd())
	playCmd.AddCommand(PlayRemoveCmd())
	playCmd.AddCommand(PlayQuitCmd())
	rootCmd.AddCommand(playCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// playCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// playCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	playCmd.Flags().BoolP("no-tui", "nt", false, "start a game without the tui")
	gameCmd.Flags().BoolP("no-tui", "nt", false, "start a game without the tui")
	listCmd.Flags().BoolP("players", "p", false, "List players in the game")
	listCmd.Flags().BoolP("games", "g", false, "List games you are in")
}
