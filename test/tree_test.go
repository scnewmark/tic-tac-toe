package test

import (
	"fmt"
	"os"
	"testing"
	"text/tabwriter"

	"github.com/scnewmark/tic-tac-toe/src/logic"
)

var (
	xwins int64
	owins int64
	draws int64
)

func TestTree(t *testing.T) {
	tree := logic.Tree()
	leafNodes := countLeafNodes(tree)

	writer := tabwriter.NewWriter(os.Stdout, 0, 8, 1, '\t', tabwriter.AlignRight)

	fmt.Fprintln(writer, "Value\tResult\tExpected")
	fmt.Fprintln(writer, "--------\t--------\t--------")
	fmt.Fprintf(writer, "Leaf nodes\t%d\t%t\n", leafNodes, int64(leafNodes) == int64(255168))
	fmt.Fprintf(writer, "X wins\t%d\t%t\n", xwins, xwins == int64(131184))
	fmt.Fprintf(writer, "O wins\t%d\t%t\n", owins, owins == int64(77904))
	fmt.Fprintf(writer, "Draws\t%d\t%t\n\n", draws, draws == int64(46080))

	writer.Flush()
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
