package ui

import (
	"fyne.io/fyne/v2"
)

// LoadMenu loads the main menu for the game.
func LoadMenu(w fyne.Window) {
	newGame := fyne.NewMenuItem("New Game", func() {
		Reset()
		LoadForm(w)
	})

	resetGame := fyne.NewMenuItem("Reset Game", Reset)

	menu := fyne.NewMenu("Game", newGame, resetGame)
	mainMenu := fyne.NewMainMenu(menu)

	w.SetMainMenu(mainMenu)
}
