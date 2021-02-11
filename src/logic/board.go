package logic

// Declarations for Empty spaces, spaces filled by X, and spaces filled by O.
const (
	EMPTY int = iota
	X
	O
)

// Board represents a playing board.
type Board struct {
	// Size is the length of each row and column.
	Size int

	// Matrix is the 2D array used to manage the boardstate.
	Matrix [][]int
}

// New creates a new board of the size specified by int s.
func New(dy int, dx int) *Board {
	b := Board{
		Size:   3,
		Matrix: make([][]int, dy),
	}

	for i := range b.Matrix {
		b.Matrix[i] = make([]int, dx)
	}

	return &b
}

// Insert is a helper method for inserting a piece at the selected row and column.
func (b Board) Insert(r int, c int, p int) {
	if b.valid(r, c, p) {
		b.Matrix[r][c] = p
	}
}

// Available is a helper method for checking if a space is available.
func (b Board) Available(r int, c int) bool {
	if b.valid(r, c) {
		return b.Matrix[r][c] == EMPTY
	}
	return false
}

func (b Board) valid(r int, c int, opts ...int) bool {
	if r > b.Size || c > b.Size {
		return false
	} else if len(opts) > 0 && opts[0] > int(O) {
		return false
	}
	return true
}
