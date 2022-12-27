/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func ProfileCmd() *cobra.Command {

	// profileCmd represents the profile command
	var profileCmd = &cobra.Command{
		Use:   "profile [command] [flags] [args]",
		Short: "Manage and view profiles",
		Long: `You can use this command to manage and view profiles of users, groups, servers, and cards.
		Almost all objects have a profile, and you can view the profile of any object using this command.
		
		To view the profile of a player:
		profile player <player_id>
		
		To view the profile of a group:
		profile group <group_id>
		
		To view the profile of a server:
		profile server <server_id>
		
		To view the profile of a card:
		profile card <card_id>
		
		To view the profile of a deck:
		profile deck <deck_id>
		
		To view the profile of a game:
		profile game <game_id>
		
		To view your own profile:
		profile
		
		To edit your own profile:
		profile edit -me <field> <value>`,
		ValidArgs: []string{"", "player", "group", "server", "card", "deck", "game", "edit"},
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("profile called")
		},
	}
	return profileCmd
}

func ProfileEditCmd() *cobra.Command {

	// profileCmd represents the profile command
	var profileEditCmd = &cobra.Command{
		Use:     "edit [flags] <field> <value>",
		Short:   "Edit your profile",
		Example: "edit -me name Dio",
		Long: `You can use this command to edit your profile.
		You can edit your profile by specifying the field to edit and the value to set it to.
		
		To edit your name:
		profile edit -me name <name>
		
		To edit your bio:
		profile edit -me bio <bio>
		
		To edit your profile picture:
		profile edit -me picture <picture_url>
		
		To edit your banner:
		profile edit -me banner <banner_url>
		
		To edit your status:
		profile edit -me status <status>`,
		Args: cobra.MinimumNArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("profile edit called")
		},
	}
	return profileEditCmd
}



func ProfileGroupCmd() *cobra.Command {
	var profileGroupCmd = &cobra.Command{
		Use:     "group <group_id>",
		Short:   "View a group's profile",
		Example: "group 123456789",
		Long: `You can use this command to view a group's profile.
		You can view a group's profile by specifying the group's id.

		To view a group's profile:
		profile group <group_id>`,
		Args: cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("profile group called")
		},
	}
	return profileGroupCmd
}

func ProfileServerCmd() *cobra.Command {
	var profileServerCmd = &cobra.Command{
		Use:     "server",
		Short:   "View a server's profile",
		Example: "profile server 123456789",
		Long: `You can use this command to view a server's profile.
		You can view a server's profile by specifying the server's id.

		To view a server's profile:
		profile server <server_id>`,
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println("profile server called")
			return nil
		},
	}
	return profileServerCmd
}

func ProfileCardCmd() *cobra.Command {
	var profileCardCmd = &cobra.Command{
		Use:     "card",
		Short:   "View a card's profile",
		Example: "profile card 123456789",
		Long: `You can use this command to view a card's profile.
		You can view a card's profile by specifying the card's id.

		To view a card's profile:
		profile card <card_id>`,
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println("profile card called")
			return nil
		},
	}
	return profileCardCmd
}

func ProfileDeckCmd() *cobra.Command {
	var profileDeckCmd = &cobra.Command{
		Use:     "deck [args]",
		Short:   "View a deck's profile",
		Example: "profile deck 123456789",
		Long: `You can use this command to view a deck's profile.
		You can view a deck's profile by specifying the deck's id.

		To view a deck's profile:
		profile deck <deck_id>`,
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println("profile deck called")
			return nil
		},
	}
	return profileDeckCmd
}

func ProfileGameCmd() *cobra.Command {
	var profileGameCmd = &cobra.Command{
		Use:     "game [args]",
		Short:   "View a game's profile",
		Example: "profile game 123456789",
		Long: `You can use this command to view a game's profile.
		You can view a game's profile by specifying the game's id.

		To view a game's profile:
		profile game <game_id>`,
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println("profile game called")
			return nil
		},
	}
	return profileGameCmd
}

func ProfilePlayerCmd() *cobra.Command {
	var profilePlayerCmd = &cobra.Command{
		Use:     "player [args]",
		Short:   "View a player's profile",
		Example: "profile player 123456789",
		Long: `You can use this command to view a player's profile.
		You can view a player's profile by specifying the player's id.

		To view a player's profile:
		profile player <player_id>`,
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println("profile player called")
			return nil
		},
	}
	return profilePlayerCmd
}


func init() {
	// declare the profile command
	profileCmd := ProfileCmd()
	// add flags to the commands
	editCmd := ProfileEditCmd()
	editCmd.Flags().BoolP("me", "m", false, "Edit your profile")
	// add subcommands to the profile command
	profileCmd.AddCommand(editCmd)
	profileCmd.AddCommand(ProfilePlayerCmd())
	profileCmd.AddCommand(ProfileGroupCmd())
	profileCmd.AddCommand(ProfileServerCmd())
	profileCmd.AddCommand(ProfileCardCmd())
	profileCmd.AddCommand(ProfileDeckCmd())
	profileCmd.AddCommand(ProfileGameCmd())
	// add profile command to root command
	rootCmd.AddCommand(profileCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// profileCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// profileCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

