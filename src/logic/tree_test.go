package logic_test

import (
	"testing"

	"github.com/scnewmark/tic-tac-toe/src/logic"
)

var (
	xwins int
	owins int
	draws int
)

func Test(t *testing.T) {
	tree := logic.Tree()
	leafNodes := countLeafNodes(tree)
	t.Logf("Total leaf nodes: %d\n", leafNodes)
	t.Logf("X Wins: %d\n", xwins)
	t.Logf("O Wins: %d\n", owins)
	t.Logf("Draws: %d\n", draws)
}

func countLeafNodes(root *logic.Node) (nodes int) {
	for _, child := range root.Children {
		if len(child.Children) < 1 {
			nodes++
			if logic.Winner(child.Value).Exists {
				if logic.Winner(child.Value).Player == logic.X {
					xwins++
				} else if logic.Winner(child.Value).Player == logic.O {
					owins++
				}
			} else {
				draws++
			}
		}
		n := countLeafNodes(child)
		nodes += n
	}
	return nodes
}
