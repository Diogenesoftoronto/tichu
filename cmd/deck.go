/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"strconv"

	"encoding/json"

	"github.com/Diogenesoftoronto/tichu/deck"
	"github.com/google/uuid"
	"github.com/spf13/cobra"
)

type DB struct {
	decks []*deck.Deck
}

func NewDB() *DB {
	return new(DB)
}

func DeckCmd() *cobra.Command {
	// deckCmd represents the deck command
	var deckCmd = &cobra.Command{
		Use:   "deck [flags] [commands] [args]",
		Short: "Manage the cards in a deck",
		Long: `This command will allow you to manage the cards in a deck.
		
		You can create a new deck, 
		deck create <deck_type>
		
		add cards to a deck, 
		deck add <deck_id> <card_id>
		
		remove cards from a deck, 
		deck remove <deck_id> <card_id>...
		
		and shuffle a deck.
		deck shuffle <deck_id>
		
		You can also view the cards in a deck.
		deck <deck_id>`,
		Args:      cobra.MinimumNArgs(1),
		ValidArgs: []string{"", "create", "add", "remove", "shuffle", "view"},
		Example:   "deck create deck",
		RunE: func(cmd *cobra.Command, args []string) error {
			// Create a mock database with a uuid as the id of the deck
			// that I want to retrieve
			dbptr := NewDB()
			// I must create a deck with a particular uuid
			Deck := deck.New(1)
			Id, err := uuid.Parse(args[0])
			if err != nil {
				return err
			}
			Deck.Id = Id
			// create an array of decks
			deck_arr := []*deck.Deck{Deck}
			// Then place decks into a db
			dbptr.decks = deck_arr

			// check the json flag to see if it is true
			flag, err := cmd.Flags().GetBool("json")
			if err != nil {
				return err
			}
			// if the json flag is true then set the message to json instead of a raw struct
			// TODO: might be a good case for generics https://gobyexample.com/generics
			var message any
			if flag {
				// to send as json i need to marshall the deck
				message, err = json.Marshal(Deck); if err != nil {
					return err
				}
			} else {

				message = Deck

			}
			sender := cmd.OutOrStdout()
			fmt.Fprint(sender, message)
			return nil
		},
	}
	return deckCmd
}

func DeckCreateCmd() *cobra.Command {
	// deckCmd represents the deck command
	var deckCreateCmd = &cobra.Command{
		Use:     "create [flags] ...<arg>",
		Short:   "Create a new deck",
		Example: "create 1",
		Long: `You can use this command to create a new deck.
		
		To create a new deck:
		deck create <deck_type>
		
		You can create multiple types of decks by passing the --amount flag.
		deck create --amount <amount_int> <deck_type>
		
		Note that this argument returns the deck id of the new deck.`,
		Args: cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			flag, err := cmd.Flags().GetInt("amount")
			if err != nil {
				return err
			}
			deck_type, err := strconv.ParseInt(args[0], 10, 8)
			if err != nil {
				return err
			}
			deck_int := int(deck_type)
			dtype, err := deck.ParseType(deck_int)
			if err != nil {
				return err
			}
			for i := 0; i < flag; i++ {
				d := deck.New(dtype)
				message := d.Id
				sender := cmd.OutOrStdout()
				fmt.Fprintln(sender, message)
			}
			return nil
		},
	}
	return deckCreateCmd
}

func DeckAddCmd() *cobra.Command {
	// deckCmd represents the deck command
	var deckAddCmd = &cobra.Command{
		Use:     "add [flags] <deck_id> <card_id>",
		Short:   "Add a card to a deck",
		Example: "add 123456789012345678901234 123456789012345678901234",
		Long: `You can use this command to add a card to a deck.

		To add a card to a deck:
		deck add <deck_id> <card_id>`,
		Args: cobra.MinimumNArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("deck add called")
		},
	}
	return deckAddCmd
}

func DeckRemoveCmd() *cobra.Command {
	// deckCmd represents the deck command
	var deckRemoveCmd = &cobra.Command{
		Use:     "remove [flags] <deck_id> <card_id>",
		Short:   "Remove a card from a deck",
		Example: "remove 123 123",
		Long: `You can use this command to remove a card from a deck.

		To remove a card from a deck:
		deck remove <deck_id> <card_id>`,
		Args: cobra.MinimumNArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("deck remove called")
		},
	}
	return deckRemoveCmd
}

func DeckShuffleCmd() *cobra.Command {
	// deckCmd represents the deck command
	var deckShuffleCmd = &cobra.Command{
		Use:     "shuffle [flags] <deck_id>",
		Short:   "Shuffle a deck",
		Example: "shuffle 123456789012345678901234",
		Long: `You can use this command to shuffle a deck.

		To shuffle a deck:
		deck shuffle <deck_id>`,
		Args: cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("deck shuffle called")
		},
	}
	return deckShuffleCmd
}

// func DeckViewCmd() *cobra.Command {
// 	// deckCmd represents the deck command
// 	var deckViewCmd = &cobra.Command{
// 		Use:     "view [flags] <deck_id>",
// 		Short:   "View the cards in a deck",
// 		Example: "view 123456789012345678901234",
// 		Long: `You can use this command to view the cards in a deck.

// 		To view the cards in a deck:
// 		deck view <deck_id>`,
// 		Args: cobra.MinimumNArgs(1),
// 		Run: func(cmd *cobra.Command, args []string) {
// 			fmt.Println("deck view called")
// 		},
// 	}
// 	return deckViewCmd
// }

func init() {

	deckCmd := DeckCmd()
	creatCmd := DeckCreateCmd()

	deckCmd.AddCommand(creatCmd)
	deckCmd.AddCommand(DeckAddCmd())
	deckCmd.AddCommand(DeckRemoveCmd())
	deckCmd.AddCommand(DeckShuffleCmd())
	rootCmd.AddCommand(deckCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deckCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:

	// deckCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	creatCmd.Flags().IntP("amount", "a", 1, "The amount of decks to create")
	deckCmd.Flags().Bool("json", false, "Send output as json")
}
