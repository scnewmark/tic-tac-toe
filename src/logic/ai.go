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
func Move(n *Node) (xpos int, ypos int) {
	var board [3][3]int
	value := minimax(n, len(n.Children), false)

	for _, child := range n.Children {
		if child.Weight-heuristic(child) == value {
			board = child.Value
		}
	}

	var x, y int = pos(n.Value, board)
	for {
		if available(board, x, y) {
			break
		}
		x, y = rand.Intn(3), rand.Intn(3)
	}

	return x, y
}

// Find finds the current boardstate.
func Find(node *Node, b [][]int) (res *Node) {
	var depth int = depth(b)
	for _, child := range node.Children {
		res = child
		if len(res.Children) != depth {
			Find(child, b)
		} else {
			break
		}
	}
	return res
}

func depth(b [][]int) (depth int) {
	for row := 0; row < 3; row++ {
		for col := 0; col < 3; col++ {
			if b[row][col] == EMPTY {
				depth++
			}
		}
	}
	return depth
}

func available(board [3][3]int, row int, col int) bool {
	if board[row][col] == EMPTY {
		return true
	}
	return false
}

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

// minimax is an implementation of the minimax algorithm
//
// https://en.wikipedia.org/wiki/Minimax#Pseudocode
func minimax(n *Node, d int, maximizing bool) int64 {
	if d == 0 || len(n.Children) < 1 {
		return n.Weight + heuristic(n)
	}
	if maximizing {
		value := NegativeInfinity
		for _, child := range n.Children {
			value = max(child.Weight+heuristic(child), minimax(child, d-1, false))
		}
		return value
	}
	value := PositiveInfinity
	for _, child := range n.Children {
		value = min(child.Weight-heuristic(child), minimax(child, d-1, true))
	}
	return value
}

// heuristic is a function for modifying the value of a child node based on conditions that will lead to a win.
func heuristic(n *Node) int64 {
	if fork(n.Value) {
		return 10
	} else if two(n.Value) {
		return 5
	}
	return 0
}

// two returns a boolean indicative of whether or not board b is a two in a row.
func two(b [3][3]int) bool {
	xv, xd, xh, ov, od, oh := check(b)

	var x []int = []int{xv, xd, xh}
	var o []int = []int{ov, od, oh}

	for i := 0; i < len(x); i++ {
		for k := 0; k < len(o); k++ {
			if x[i]+x[k] == 1 && i != k {
				return true
			} else if o[i]+o[k] == 1 && i != k {
				return true
			}
		}
	}

	return false
}

// fork returns a boolean indicative of whether or not board b is a fork.
func fork(b [3][3]int) bool {
	xv, xd, xh, ov, od, oh := check(b)

	var x []int = []int{xv, xd, xh}
	var o []int = []int{ov, od, oh}

	for i := 0; i < len(x); i++ {
		for k := 0; k < len(o); k++ {
			if x[i]+x[k] > 1 && i != k {
				return true
			} else if o[i]+o[k] > 1 && i != k {
				return true
			}
		}
	}

	return false
}

func check(b [3][3]int) (int, int, int, int, int, int) {
	var xv, xd, xh int = 0, 0, 0
	var ov, od, oh int = 0, 0, 0

	var xtaken, otaken = 0, 0
	for row := 0; row < len(b); row++ {
		for col := 0; col < len(b[row]); col++ {
			if b[row][col] == X {
				xtaken++
			} else if b[row][col] == O {
				otaken++
			}
		}

		if xtaken == 2 && otaken == 0 {
			xh++
		} else if otaken == 2 && xtaken == 0 {
			oh++
		}
	}

	xtaken, otaken = 0, 0
	for row := 0; row < len(b[0]); row++ {
		for col := 0; col < len(b); col++ {
			if b[col][row] == X {
				xtaken++
			} else if b[col][row] == O {
				otaken++
			}
		}
		if xtaken == 2 && otaken == 0 {
			xv++
		} else if otaken == 2 && xtaken == 0 {
			ov++
		}
	}

	xtaken, otaken = 0, 0
	for i := 0; i < len(b); i++ {
		if b[i][i] == X {
			xtaken++
		} else if b[i][i] == O {
			otaken++
		}
	}

	if xtaken == 2 && otaken == 0 {
		xd++
	} else if otaken == 2 && xtaken == 0 {
		od++
	}

	xtaken, otaken = 0, 0
	for i := 0; i < len(b); i++ {
		if b[3-i-1][i] == X {
			xtaken++
		} else if b[3-i-1][i] == O {
			otaken++
		}
	}

	if xtaken == 2 && otaken == 0 {
		xd++
	} else if otaken == 2 && xtaken == 0 {
		od++
	}

	return xv, xd, xh, ov, od, oh
}

// max retuns the higher number of a and b.
func max(a int64, b int64) int64 {
	if a > b {
		return a
	}
	return b
}

// min returns the lower number of a and b.
func min(a int64, b int64) int64 {
	if a < b {
		return a
	}
	return b
}
