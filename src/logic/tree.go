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
func Tree() *Node {
	var root *Node = &Node{
		Value:    [3][3]int{},
		Children: []*Node{},
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
				if !winner(matrix) {
					generateChildren(node.Children[child])
				}
				child++
			}
		}
	}
}

func winner(matrix [3][3]int) bool {
	for row := 0; row < 3; row++ {
		if (matrix[row][0] == X && matrix[row][1] == X && matrix[row][2] == X) || (matrix[row][0] == O && matrix[row][1] == O && matrix[row][2] == O) {
			return true
		}
	}

	for col := 0; col < 3; col++ {
		if (matrix[0][col] == X && matrix[1][col] == X && matrix[2][col] == X) || (matrix[0][col] == O && matrix[1][col] == O && matrix[2][col] == O) {
			return true
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

	if xtaken == 3 || otaken == 3 {
		return true
	}

	if (matrix[0][2] == X && matrix[1][1] == X && matrix[2][0] == X) || (matrix[0][2] == O && matrix[1][1] == O && matrix[2][0] == O) {
		return true
	}

	return false
}
