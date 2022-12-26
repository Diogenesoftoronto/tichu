/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func DeckCmd() *cobra.Command {
	// deckCmd represents the deck command
	var deckCmd = &cobra.Command{
		Use:   "deck [commands] [flags] [args]",
		Short: "Manage the cards in a deck",
		Long: `This command will allow you to manage the cards in a deck.
		
		You can create a new deck, 
		deck create <deck_id>
		
		add cards to a deck, 
		deck add <deck_id> <card_id>
		
		remove cards from a deck, 
		deck remove <deck_id> <card_id>
		
		and shuffle a deck.
		deck shuffle <deck_id>
		
		You can also view the cards in a deck.
		deck <deck_id>`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("deck called")
		},
	}
	return deckCmd
}

func init() {
	rootCmd.AddCommand(DeckCmd())

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deckCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deckCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
