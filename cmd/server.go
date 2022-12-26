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
		server leave
		
		You can view the list of servers,
		server list`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("server called")
		},
	}
	return serverCmd
}

func init() {
	rootCmd.AddCommand(ServerCmd())

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serverCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serverCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
