package ui

import (
	"fyne.io/fyne/v2"
)

// LoadMenu loads the main menu for the game.
func LoadMenu(w fyne.Window) {
	newGame := &fyne.MenuItem{
		Label: "New Game",
		Action: func() {
			// TODO: Create a new game
		},
	}
	loadGame := &fyne.MenuItem{
		Label: "Load Game",
		Action: func() {
			// TODO: Load a game
		},
	}
	saveGame := &fyne.MenuItem{
		Label: "Save Game",
		Action: func() {
			// TODO: Save a current game
		},
	}
	gameMenu := []*fyne.MenuItem{newGame, loadGame, saveGame}
	menu := &fyne.Menu{
		Label: "Game",
		Items: gameMenu,
	}
	mainMenu := fyne.NewMainMenu(menu)
	w.SetMainMenu(mainMenu)
}
