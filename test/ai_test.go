package test

import (
	"testing"

	"github.com/scnewmark/tic-tac-toe/src/logic"
)

func TestAI(t *testing.T) {
	arr := logic.Tree().Children[0]
	xpos, ypos := logic.Move(arr)
	t.Log(xpos, ypos)
}
