package test

import (
	"testing"

	"github.com/scnewmark/tic-tac-toe/src/logic"
)

func TestAI(t *testing.T) {
	arr := logic.Tree().Children[0]
	board := logic.NewBoard(3, 3)
	xpos, ypos := board.Move(arr)
	t.Log(xpos, ypos)
}
