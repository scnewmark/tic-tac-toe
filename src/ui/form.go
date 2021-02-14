package ui

import (
	"image/color"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/scnewmark/tic-tac-toe/src/logic"
)

var (
	// Title is the form title.
	Title *canvas.Text

	// TitleColor is the form title color.
	TitleColor color.Color = color.White

	sizes []string = []string{"3x3", "4x4", "5x5", "6x6", "7x7", "8x8", "9x9", "10x10"}
	modes []string = []string{"Single Player", "Two Player"}
)

// LoadForm creates a new selection form and adds it to the window content.
func LoadForm(w fyne.Window) {
	// reset variables for new games
	logic.CurrentMode = logic.SINGLEPLAYER
	logic.CurrentTurn = logic.X

	label := widget.NewLabel("For boards larger than 3x3, single player mode is disabled")
	label.Hide()

	playerMode := widget.NewRadioGroup(modes, func(s string) {})
	sizeMenu := widget.NewSelect(sizes, func(s string) {
		if s != "3x3" {
			playerMode.SetSelected("Two Player")
			playerMode.Disable()
			label.Show()
		} else {
			playerMode.Enable()
			label.Hide()
		}
	})

	form := &widget.Form{
		SubmitText: "Play",
		Items: []*widget.FormItem{
			{Text: "Board size:", Widget: sizeMenu},
			{Text: "Mode:", Widget: playerMode},
			{Text: "", Widget: label}},
		OnSubmit: func() {
			w.Hide()

			conv := sizeMenu.Selected[0:1]
			if sizeMenu.Selected == "10x10" {
				conv = sizeMenu.Selected[0:2]
			}
			if playerMode.Selected == "Two Player" {
				logic.CurrentMode = logic.TWOPLAYER
			}

			if val, err := strconv.Atoi(conv); err == nil {
				c := LoadBoard(w, val)
				w.SetContent(c)
				w.Show()
			}
		},
	}

	Title = canvas.NewText("Welcome to Tic Tac Toe", TitleColor)

	Title.TextSize = 30
	Title.Alignment = fyne.TextAlignCenter
	Title.TextStyle.Bold = true

	home := container.New(layout.NewVBoxLayout(), widget.NewLabel(""), Title, widget.NewLabel(""), widget.NewSeparator(), widget.NewLabel(""), form)

	w.SetContent(home)
}
