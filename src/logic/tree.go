package logic

// Node represents a game tree node.
type Node struct {
	// Value contains the value of this node.
	Value [3][3]int

	// Children is a slice containing children of this node.
	Children []*Node
}

// previous represents the most recently inserted value
var previous int = O

// Tree generates a new game tree.
func Tree() {
	var root *Node = &Node{
		Value:    [3][3]int{},
		Children: []*Node{},
	}
	generateChildren(root)
}

// generateChildren recursively generates children for a node until the tree is complete.
func generateChildren(node *Node) {
	var child, player int

	if previous == X {
		player = O
	} else {
		player = X
	}

	for i := 0; i < 3; i++ {
		for k := 0; k < 3; k++ {
			var matrix [3][3]int = node.Value
			if matrix[i][k] == EMPTY {
				matrix[i][k] = player
				previous = player
			}
			if matrix != node.Value {
				node.Children = append(node.Children, &Node{
					Value:    matrix,
					Children: []*Node{},
				})
				generateChildren(node.Children[child])
				child++
			}
		}
	}
}
