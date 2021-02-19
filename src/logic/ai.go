package logic

import (
	"math"
	"math/rand"
)

// Representations for infinity for minimax algorithm.
const (
	NegativeInfinity int64 = math.MinInt64
	PositiveInfinity int64 = math.MaxInt64
)

// Move returns the best move based on the current boardstate board.
func (b Board) Move(n *Node) (xpos int, ypos int) {
	var (
		board  [3][3]int
		lowest int64 = 0
	)

	for _, child := range n.Children {
		if child.Weight-heuristic(child) < lowest {
			lowest = child.Weight
			board = child.Value
		}
	}

	var x, y int = pos(n.Value, board)
	for {
		if available(convert(b.Matrix), x, y) {
			break
		}
		x, y = rand.Intn(3), rand.Intn(3)
	}

	return x, y
}

// Find finds the current boardstate.
func Find(node *Node, b [][]int) *Node {
	for _, child := range node.Children {
		node := Find(child, b)
		if node.Value == convert(b) {
			return node
		}
	}
	return node
}

// available checks whether or not a position is available.
func available(board [3][3]int, row int, col int) bool {
	if board[row][col] == EMPTY {
		return true
	}
	return false
}

// convert converts a slice to an array.
func convert(b [][]int) (dest [3][3]int) {
	for row := 0; row < 3; row++ {
		for col := 0; col < 3; col++ {
			dest[row][col] = b[row][col]
		}
	}
	return dest
}

// pos returns the new coordinate pair for the chosen board.
func pos(old [3][3]int, new [3][3]int) (x int, y int) {
	for row := 0; row < 3; row++ {
		for col := 0; col < 3; col++ {
			if old[row][col] != new[row][col] {
				return row, col
			}
		}
	}
	return -1, -1
}

// heuristic is a function for modifying the value of a child node based on conditions that will lead to a win.
func heuristic(n *Node) int64 {
	if len(n.Children) == 7 && n.Value[1][1] == O {
		return 500
	}
	if block(n.Value) {
		return 15
	}
	return 0
}

// block checks for whether or not a child node blocks X from winning.
func block(b [3][3]int) bool {
	for i := 0; i < 3; i++ {
		// -------------------------------------------------------
		// Rows
		if b[i][0] == X && b[i][1] == X && b[i][2] == O {
			return true
		}

		if b[i][0] == O && b[i][1] == X && b[i][2] == X {
			return true
		}

		if b[i][0] == X && b[i][1] == O && b[i][2] == X {
			return true
		}

		// -------------------------------------------------------
		// Columns
		if b[0][i] == X && b[1][i] == O && b[2][i] == X {
			return true
		}

		if b[0][i] == O && b[1][i] == X && b[2][i] == X {
			return true
		}

		if b[0][i] == X && b[1][i] == X && b[2][i] == O {
			return true
		}
	}

	// -------------------------------------------------------
	// Diagonals
	if b[0][0] == X && b[1][1] == X && b[2][2] == O {
		return true
	}

	if b[0][0] == X && b[1][1] == O && b[2][2] == X {
		return true
	}

	if b[0][0] == O && b[1][1] == X && b[2][2] == X {
		return true
	}

	if b[2][0] == X && b[1][1] == X && b[0][2] == O {
		return true
	}

	if b[2][0] == X && b[1][1] == O && b[0][2] == X {
		return true
	}

	if b[2][0] == O && b[1][1] == X && b[0][2] == X {
		return true
	}

	return false
}
