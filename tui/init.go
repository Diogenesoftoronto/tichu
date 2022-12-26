package game

import (
	// "io/fs"

	tea "github.com/charmbracelet/bubbletea"
)

// type picture struct {
// 	fs.File
// 	path string
// }

// type player struct {
// 	name string
// 	score int32
// 	rank int32
// 	description string
// 	picture picture
// }

// type icon struct {
// 	picture picture
// 	icon_id rune
// }

// type choice struct {
// 	title string
// 	description string
// 	icon icon
// 	id int
// }

// type selection struct {
// 	chosen bool
// 	choice choice
// }

// type cursor struct {
// 	position int
// }

// type model struct {
// 	players []player
// 	choices []choice
// 	cursor cursor
// 	selected selection

// }

// MVP: getting started struct for the game
type simpleModel struct {
	players  []string
	choices  []string
	cursor   int
	selected map[int]struct{}
}

// Init is called once when the program starts. It returns the initial model
func InitialModel() simpleModel {
	// TODO: load players from the database
	return simpleModel{
		players:  []string{"You", "Ally", "Enemy 0", "Enemy 1"},
		choices:  []string{"Draw", "Grand Tichu", "Message"},
		cursor:   0,
		selected: make(map[int]struct{}),
	}
}

func (m simpleModel) Init() tea.Cmd {
	return nil
}
