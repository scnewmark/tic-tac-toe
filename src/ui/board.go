package ui

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/scnewmark/tic-tac-toe/src/logic"
)

var (
	// Buttons represents the group of buttons that make up the visible board.
	Buttons map[string]*widget.Button = make(map[string]*widget.Button)

	board *logic.Board
)

// LoadBoard loads a board of the specified size.
func LoadBoard(w fyne.Window, size int) *fyne.Container {
	c := container.New(layout.NewGridLayout(size))
	for row := 0; row < size; row++ {
		for col := 0; col < size; col++ {
			c.Objects = append(c.Objects, button(w, row, col))
		}
	}
	board = logic.NewBoard(size, size)
	c.Refresh()
	return c
}

// button is an internal helper method for creating a new button.
func button(w fyne.Window, row int, col int) *widget.Button {
	btn := widget.NewButton("", func() { updateState(w, row, col) })

	Buttons[fmt.Sprintf("%d-%d", row, col)] = btn
	return btn
}

// updateState is an internal method to manage the visible boardstate. Triggered when a button is clicked.
func updateState(w fyne.Window, row int, col int) {
	btn := Buttons[fmt.Sprintf("%d-%d", row, col)]
	if logic.CurrentTurn == logic.X {
		btn.SetText("X")
		board.Matrix[row][col] = logic.X

		logic.CurrentTurn = logic.O
	} else {
		btn.SetText("O")
		board.Matrix[row][col] = logic.O

		logic.CurrentTurn = logic.X
	}
	btn.Disable()

	if board.Full() {
		btn := widget.NewButton("Play Again", func() {})

		text := widget.NewFormItem("\t\t\t\t\t\tGame is a tie!", widget.NewLabel(""))
		fbtn := widget.NewFormItem("", btn)

		popup := widget.NewModalPopUp(widget.NewForm(text, fbtn), w.Canvas())

		btn.OnTapped = func() {
			w.Hide()
			LoadForm(w)
			w.Show()
			popup.Hide()
		}

		popup.ShowAtPosition(fyne.NewPos(200, 235))
	}
}
