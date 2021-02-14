package logic

// Declarations for Empty spaces, spaces filled by X, and spaces filled by O.
const (
	EMPTY int = iota
	X
	O
)

// Declarations of single player and two player game modes
const (
	SINGLEPLAYER int = (iota + 1) << 5
	TWOPLAYER
)

var (
	// CurrentTurn holds the current turn value.
	CurrentTurn int = X

	// CurrentMode holds the current mode value.
	CurrentMode int = SINGLEPLAYER
)

// Board represents a playing board.
type Board struct {
	// Size is the length of each row and column.
	Size int

	// Matrix is the 2D array used to manage the boardstate.
	Matrix [][]int
}

// Win represents a win situation if such a condition exists.
type Win struct {
	// Player is the player who won, if any.
	Player int

	// Exists is a bool indicating whether or not someone has won.
	Exists bool
}

// NewBoard creates a new board of the size specified by ints dy and dx.
func NewBoard(dy int, dx int) *Board {
	b := Board{
		Size:   dy,
		Matrix: make([][]int, dy),
	}

	for i := range b.Matrix {
		b.Matrix[i] = make([]int, dx)
	}

	return &b
}

// Insert is a helper method for inserting a piece at the selected row and column.
func (b Board) Insert(r int, c int, p int) {
	b.Matrix[r][c] = p
}

// Full returns a boolean value indicative of whether or not the board is full.
func (b Board) Full() bool {
	for row := 0; row < b.Size; row++ {
		for col := 0; col < b.Size; col++ {
			if b.Matrix[row][col] == EMPTY {
				return false
			}
		}
	}
	return true
}

// GetWin returns the win situation if one exists
func (b Board) GetWin() *Win {
	var win *Win

	win = b.findWin(true)
	if win.Exists {
		return win
	}

	win = b.findWin(false)
	if win.Exists {
		return win
	}

	return win
}

// findWin returns whether a win exists based on bool rfirst
func (b Board) findWin(rfirst bool) *Win {
	for row := 0; row < b.Size; row++ {
		var xtaken, otaken int = 0, 0
		for col := 0; col < b.Size; col++ {
			if rfirst {
				if b.Matrix[row][col] == X {
					xtaken++
				} else if b.Matrix[row][col] == O {
					otaken++
				}
			} else {
				if b.Matrix[col][row] == X {
					xtaken++
				} else if b.Matrix[col][row] == O {
					otaken++
				}
			}
		}
		if winner := b.win(xtaken, otaken); winner.Exists {
			return winner
		}
	}

	var xtaken, otaken = 0, 0
	for i := 0; i < b.Size; i++ {
		if b.Matrix[i][i] == X {
			xtaken++
		} else if b.Matrix[i][i] == O {
			otaken++
		}
	}

	if winner := b.win(xtaken, otaken); winner.Exists {
		return winner
	}

	xtaken, otaken = 0, 0
	for i := 0; i < b.Size; i++ {
		if b.Matrix[b.Size-i-1][i] == X {
			xtaken++
		} else if b.Matrix[b.Size-i-1][i] == O {
			otaken++
		}
	}

	if winner := b.win(xtaken, otaken); winner.Exists {
		return winner
	}

	if xtaken == b.Size {
		return &Win{Exists: true, Player: X}
	} else if otaken == b.Size {
		return &Win{Exists: true, Player: O}
	}

	return &Win{Exists: false}
}

func (b Board) win(x int, o int) *Win {
	if x == b.Size {
		return &Win{Exists: true, Player: X}
	} else if o == b.Size {
		return &Win{Exists: true, Player: O}
	}
	return &Win{Exists: false}
}
