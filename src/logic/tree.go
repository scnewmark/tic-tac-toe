package logic

// Node represents a game tree node.
type Node struct {
	// Value contains the value of this node.
	Value [3][3]int

	// Weight represents the weight of a move.
	Weight int64

	// Children is a slice containing children of this node.
	Children []*Node
}

// Constants defined for assigning weights to positions.
const (
	owin  int64 = -1000
	nowin int64 = 0
	xwin  int64 = 1000
)

// previous represents the most recently inserted value
var previous int = O

// Tree generates a new game tree.
func Tree() *Node {
	var root *Node = &Node{
		Value:    [3][3]int{},
		Children: []*Node{},
		Weight:   nowin,
	}
	generateChildren(root)
	return root
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
				if !Winner(matrix).Exists {
					generateChildren(node.Children[child])
					node.Children[child].Weight = nowin
				} else {
					if Winner(matrix).Player == X {
						node.Children[child].Weight = xwin
					} else if Winner(matrix).Player == O {
						node.Children[child].Weight = owin
					}
				}
				child++
			}
		}
	}
}

// Winner is a helper function for checking if a child node has a winner.
func Winner(matrix [3][3]int) *Win {
	for row := 0; row < 3; row++ {
		if matrix[row][0] == X && matrix[row][1] == X && matrix[row][2] == X {
			return &Win{Exists: true, Player: X}
		} else if matrix[row][0] == O && matrix[row][1] == O && matrix[row][2] == O {
			return &Win{Exists: true, Player: O}
		}
	}

	for col := 0; col < 3; col++ {
		if matrix[0][col] == X && matrix[1][col] == X && matrix[2][col] == X {
			return &Win{Exists: true, Player: X}
		} else if matrix[0][col] == O && matrix[1][col] == O && matrix[2][col] == O {
			return &Win{Exists: true, Player: O}
		}
	}

	var xtaken, otaken = 0, 0
	for i := 0; i < 3; i++ {
		if matrix[i][i] == X {
			xtaken++
		} else if matrix[i][i] == O {
			otaken++
		}
	}

	if xtaken == 3 {
		return &Win{Exists: true, Player: X}
	} else if otaken == 3 {
		return &Win{Exists: true, Player: O}
	}

	if matrix[0][2] == X && matrix[1][1] == X && matrix[2][0] == X {
		return &Win{Exists: true, Player: X}
	} else if matrix[0][2] == O && matrix[1][1] == O && matrix[2][0] == O {
		return &Win{Exists: true, Player: O}
	}

	return &Win{Exists: false}
}
