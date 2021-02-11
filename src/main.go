package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"github.com/scnewmark/tic-tac-toe/src/ui"
)

func main() {
	app := app.New()
	window := app.NewWindow("Tic Tac Toe")

	ui.LoadMenu(window)
	ui.LoadForm(window)

	window.Resize(fyne.NewSize(512, 512))
	window.CenterOnScreen()
	window.ShowAndRun()
}
