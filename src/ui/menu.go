package ui

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

// LoadMenu loads the main menu for the game.
func LoadMenu(w fyne.Window) {
	newGame := fyne.NewMenuItem("New Game", func() {
		Reset()
		LoadForm(w)
	})

	resetGame := fyne.NewMenuItem("Reset Game", Reset)

	light := fyne.NewMenuItem("Light Theme", func() {
		fyne.CurrentApp().Settings().SetTheme(theme.LightTheme())
		TitleColor = color.Black
		Title.Color = TitleColor
	})
	dark := fyne.NewMenuItem("Dark Theme", func() {
		fyne.CurrentApp().Settings().SetTheme(theme.DarkTheme())
		TitleColor = color.White
		Title.Color = TitleColor
	})

	gameMenu := fyne.NewMenu("Game", newGame, resetGame)
	themeMenu := fyne.NewMenu("Theme", light, dark)
	mainMenu := fyne.NewMainMenu(gameMenu, themeMenu)

	w.SetMainMenu(mainMenu)
}
