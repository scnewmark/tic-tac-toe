package ui

import (
	"fmt"
	"strconv"
	"strings"

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
			c.Add(button(w, row, col))
		}
	}
	board = logic.NewBoard(size, size)
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
	btn.Disable()

	if logic.CurrentTurn == logic.X {
		btn.SetText("X")
		board.Insert(row, col, logic.X)

		if checkWin(w, btn, row, col) {
			return
		}

		logic.CurrentTurn = logic.O
	} else {
		btn.SetText("O")
		board.Insert(row, col, logic.O)

		if checkWin(w, btn, row, col) {
			return
		}

		logic.CurrentTurn = logic.X
	}

	if board.Full() {
		createPopup(w, "Game is a tie!")
	}
}

// Reset resets the current boardstate to empty.
func Reset() {
	for key, btn := range Buttons {
		btn.SetText("")
		btn.Enable()

		values := strings.Split(key, "-")
		row, _ := strconv.Atoi(values[0])
		col, _ := strconv.Atoi(values[1])

		board.Insert(row, col, logic.EMPTY)
	}
	logic.CurrentTurn = logic.X
}

func createPopup(w fyne.Window, t string) {
	btn := widget.NewButton("Play Again", func() {})

	text := widget.NewFormItem(fmt.Sprintf("\t\t\t\t\t\t%s", t), widget.NewLabel(""))
	fbtn := widget.NewFormItem("", btn)

	popup := widget.NewModalPopUp(widget.NewForm(text, fbtn), w.Canvas())

	btn.OnTapped = func() {
		Reset()
		popup.Hide()
	}

	popup.ShowAtPosition(fyne.NewPos(200, 235))
}

func checkWin(w fyne.Window, btn *widget.Button, row int, col int) bool {
	var win *logic.Win = board.GetWin()

	if win.Exists && win.Player == logic.X {
		createPopup(w, "X won the game!")
		return true
	} else if win.Exists && win.Player == logic.O {
		createPopup(w, "O won the game!")
		return true
	}
	return false
}
