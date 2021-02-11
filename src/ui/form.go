package ui

import (
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"github.com/scnewmark/tic-tac-toe/src/logic"
)

var (
	sizes  []string = []string{"3x3", "4x4", "5x5", "6x6", "7x7", "8x8", "9x9", "10x10"}
	pieces []string = []string{"X", "O"}
	modes  []string = []string{"Single Player", "Two Player"}
)

// LoadForm creates a new selection form and adds it to the window content.
func LoadForm(w fyne.Window) {
	// reset variables for new games
	logic.CurrentMode = logic.SINGLEPLAYER
	logic.CurrentTurn = logic.X

	sizeMenu := widget.NewSelect(sizes, func(s string) {})
	pieceMenu := widget.NewSelect(pieces, func(s string) {})
	playerMode := widget.NewRadioGroup(modes, func(s string) {})

	form := &widget.Form{
		SubmitText: "Play",
		Items: []*widget.FormItem{
			{Text: "Select a piece:", Widget: pieceMenu},
			{Text: "Select a board:", Widget: sizeMenu},
			{Text: "Select a mode:", Widget: playerMode}},
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

	w.SetContent(form)
}
