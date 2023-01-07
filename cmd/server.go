/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func ServerCmd() *cobra.Command {

	// serverCmd represents the server command
	var serverCmd = &cobra.Command{
		Use:   "server [commands] [args]",
		Short: "Manage, view, and join servers",
		Long: `Manage, view, and join servers
		You can create a new server,
		server create 
		
		You can join a server,
		server join <server_id>
	
		You can delete a server(you must be the server owner to do this),
		server delete <server_id>
		
		You can leave a server,
		server leave <server_id>
		
		You can view the list of servers,
		server`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("server called")
		},
	}
	return serverCmd
}

func ServerCreateCmd() *cobra.Command {

	// serverCmd represents the server command
	var serverCreateCmd = &cobra.Command{
		Use:   "create [flags] [args]",
		Short: "Create a new server",
		Long: `Create a new server
		You can create a new server,
		server create `,
		Example: "server create",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("server create called")
		},
	}
	return serverCreateCmd
}

func ServerJoinCmd() *cobra.Command {

	// serverCmd represents the server command
	var serverJoinCmd = &cobra.Command{
		Use:   "join [flags] [args]",
		Short: "Join a server",
		Long: `Join a server
		You can join a server,
		server join <server_id>`,
		Example: "server join <server_id>",
		Args:    cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("server join called")
		},
	}
	return serverJoinCmd
}

func ServerDeleteCmd() *cobra.Command {

	// serverCmd represents the server command
	var serverDeleteCmd = &cobra.Command{
		Use:   "delete [flags] [args]",
		Short: "Delete a server",
		Long: `Delete a server
		You can delete a server(you must be the server owner to do this),
		server delete <server_id>`,
		Example: "server delete <server_id>",
		Args:    cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("server delete called")
		},
	}
	return serverDeleteCmd
}

func ServerLeaveCmd() *cobra.Command {

	// serverCmd represents the server command
	var serverLeaveCmd = &cobra.Command{
		Use:   "leave [flags] [args]",
		Short: "Leave a server",
		Long: `Leave a server
		You can leave a server,
		server leave <server_id>`,
		Example: "server leave <server_id>",
		Args:    cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("server leave called")
		},
	}
	return serverLeaveCmd
}

func init() {
	serverCmd := ServerCmd()

	serverCmd.AddCommand(ServerCreateCmd())
	serverCmd.AddCommand(ServerJoinCmd())
	serverCmd.AddCommand(ServerDeleteCmd())
	serverCmd.AddCommand(ServerLeaveCmd())
	rootCmd.AddCommand(serverCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serverCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serverCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
