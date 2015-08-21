package chessboard

// King model
type King struct{}

func (k *King) Movements() []PieceMovement {
	return []PieceMovement{
		// Horizontal movements
		{+1, 0}, {-1, 0},
		// Vertical movements
		{0, +1}, {0, -1},
		// Diagonal movements
		{+1, +1}, {-1, -1}, {+1, -1}, {-1, +1},
	}
}

func (k *King) Symbol() string {
	return "♚"
}

func NewKing(board *Board, x, y uint) *Piece {
	return &Piece{
		PieceInterface: &King{},
		Board:          board,
		XPos:           int(x),
		YPos:           int(y),
	}
}

/*
// Queen model
type Queen struct{}

// Rook model
type Rook struct{}

// Bishop model
type Bishop struct{}

// Knight model
type Knight struct{}
*/
