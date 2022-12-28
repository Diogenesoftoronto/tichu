// /*
// Copyright © 2022 NAME HERE <EMAIL ADDRESS>

// */
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// messageCmd represents the message command
func MessageCmd() *cobra.Command {
	messageCmd := &cobra.Command{
		Use:   "message",
		Short: "send a message to a user",
		Long: `You can send messages to other users using this command.
		
		Example:
		message send -p <player_id> "<message>"
		
		You can also send a message to a group.
		message send -g <group_id> "<message>"
		
		You can also use this command to view messages.
		
		To view messages from a user:
		message -p <player_id>
		
		To view messages from a group:
		message -g <group_id>
		
		To view messages from all users and groups:
		message --all --global | --local | --group | --player | --server`,
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println("message called")
			return nil
		},
	}
	return messageCmd
}

func MessageSendCmd() *cobra.Command {
	messageSendCmd := &cobra.Command{
		Use:   "send",
		Short: "Send a message to a user or group",
		Long: `You can send messages to other users using this command.
		send -p <player_id> "<message>"

		You can also send a message to a group.
		send -g <group_id> "<message>"`,
		Example: "send -p <player_id> \"<message>\"",
		Args:    cobra.MaximumNArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println("send called")
			return nil
		},
	}
	return messageSendCmd
}

func init() {

	messageCmd := MessageCmd()

	messageCmd.AddCommand(MessageSendCmd())
	rootCmd.AddCommand(messageCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// mark the flags as required

	messageCmd.PersistentFlags().BoolP("global", "e", false, "for a global context")
	messageCmd.PersistentFlags().BoolP("local", "l", true, "for a local context")
	messageCmd.PersistentFlags().BoolP("group", "g", false, "for a group context")
	messageCmd.PersistentFlags().BoolP("player", "p", false, "for a player context")
	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	messageCmd.Flags().BoolP("all", "a", false, "view all messages for all users and groups")
}
