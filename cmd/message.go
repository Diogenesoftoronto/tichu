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
		message <user_id> "<message>"
		
		You can also send a message to a group.
		message <group_id> "<message>"
		
		You can also use this command to view messages.
		
		To view messages from a user:
		message -v <user_id>
		
		To view messages from a group:
		message -v <group_id>
		
		To view messages from all users and groups:
		message -v -all`,
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println("message called")
			return nil
		},
	}
	return messageCmd
}

func init() {
	rootCmd.AddCommand(MessageCmd())

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// messageCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// messageCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
