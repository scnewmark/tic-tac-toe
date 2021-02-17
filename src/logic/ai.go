package logic

import (
	"math"
)

// Representations for infinity for minimax algorithm.
const (
	NegativeInfinity int64 = math.MinInt64
	PositiveInfinity int64 = math.MaxInt64
)

// AI represents the computer opponent.
type AI struct {
	// Tree represents the completed game tree
	Tree *Node
}

// Position represents an (x, y) coordinate pair.
type Position struct {
	// X represents the X coordinate.
	X int

	// Y represents the Y coordinate.
	Y int
}

// NewAI instantiates a new AI variable and returns a pointer to it.
func NewAI() *AI {
	return &AI{
		Tree: Tree(),
	}
}

// Move returns the best move based on the current boardstate board.
func (a *AI) Move(board [][]int) int64 {
	value := minimax(a.Tree, 9, true)
	return value
}

// minimax is an implementation of the minimax algorithm
//
// https://en.wikipedia.org/wiki/Minimax#Pseudocode
func minimax(n *Node, d int, maximizing bool) int64 {
	if d == 0 || len(n.Children) < 1 {
		return heuristic(n) + n.Weight
	}
	if maximizing {
		value := NegativeInfinity
		for _, child := range n.Children {
			value = max(heuristic(child)+child.Weight, minimax(child, d-1, false))
		}
		return value
	}
	value := PositiveInfinity
	for _, child := range n.Children {
		value = min(heuristic(child)-child.Weight, minimax(child, d-1, true))
	}
	return value
}

// heuristic is a function for modifying the value of a child node based on conditions that will lead to a win.
func heuristic(n *Node) int64 {
	var forked bool = fork(n.Value)
	if forked {
		return 10
	}
	return 0
}

// fork returns a boolean indicative of whether or not board b is a fork.
func fork(b [3][3]int) bool {
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
