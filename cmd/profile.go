/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// profileCmd represents the profile command
var profileCmd = &cobra.Command{
	Use:   "profile",
	Short: "Manage and view profiles",
	Long: `You can use this command to manage and view profiles of users, groups, servers, and cards.
	Almost all objects have a profile, and you can view the profile of any object using this command.
	
	To view the profile of a user:
	profile -p <player_id>
	
	To view the profile of a group:
	profile --group <group_id>
	
	To view the profile of a server:
	profile -s <server_id>
	
	To view the profile of a card:
	profile -c <card_id>
	
	To view the profile of a deck:
	profile -d <deck_id>
	
	To view the profile of a game:
	profile --game <game_id>
	
	To view your own profile:
	profile
	
	To edit your own profile:
	profile -e <field> <value>`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("profile called")
	},
}

func init() {
	rootCmd.AddCommand(profileCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// profileCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// profileCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
