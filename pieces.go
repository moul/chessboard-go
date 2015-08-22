package chessboard

// King model
type King struct {
	Piece *Piece
}

// Movements are King's movements
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

// Symbol is the King's symbol
func (k *King) Symbol() string {
	return "♚"
}

// NewKing returns a King-flavored Piece
func NewKing(board *Board, x, y uint) *Piece {
	piece := Piece{
		Board: board,
		XPos:  int(x),
		YPos:  int(y),
	}
	piece.PieceInterface = &King{Piece: &piece}
	return &piece
}

// Queen model
type Queen struct {
	Piece *Piece
}

// Movements are Queen's movements
func (k *Queen) Movements() []PieceMovement {
	movements := append(k.Piece.HorizontalMovements(), k.Piece.VerticalMovements()...)
	movements = append(movements, k.Piece.DiagonalMovements()...)
	return movements
}

// Symbol is the Queen's symbol
func (k *Queen) Symbol() string {
	return "♛"
}

// NewQueen returns a Queen-flavored Piece
func NewQueen(board *Board, x, y uint) *Piece {
	piece := Piece{
		Board: board,
		XPos:  int(x),
		YPos:  int(y),
	}
	piece.PieceInterface = &Queen{Piece: &piece}
	return &piece
}

// Rook model
type Rook struct {
	Piece *Piece
}

// Movements are Rook's movements
func (k *Rook) Movements() []PieceMovement {
	return append(k.Piece.HorizontalMovements(), k.Piece.VerticalMovements()...)
}

// Symbol is the Rook's symbol
func (k *Rook) Symbol() string {
	return "♜"
}

// NewRook returns a Rook-flavored Piece
func NewRook(board *Board, x, y uint) *Piece {
	piece := Piece{
		Board: board,
		XPos:  int(x),
		YPos:  int(y),
	}
	piece.PieceInterface = &Rook{Piece: &piece}
	return &piece
}

// Bishop model
type Bishop struct {
	Piece *Piece
}

// Movements are Bishop's movements
func (k *Bishop) Movements() []PieceMovement {
	return k.Piece.DiagonalMovements()
}

// Symbol is the Bishop's symbol
func (k *Bishop) Symbol() string {
	return "♝"
}

// NewBishop returns a Bishop-flavored Piece
func NewBishop(board *Board, x, y uint) *Piece {
	piece := Piece{
		Board: board,
		XPos:  int(x),
		YPos:  int(y),
	}
	piece.PieceInterface = &Bishop{Piece: &piece}
	return &piece
}

// Knight model
type Knight struct {
	Piece *Piece
}

// Movements are Knight's movements
func (k *Knight) Movements() []PieceMovement {
	return []PieceMovement{
		// top-left movements
		{-2, +1}, {-1, +2},
		// top-right movements
		{+2, +1}, {+1, +2},
		// bottom-right movements
		{+2, -1}, {+1, -2},
		// bottom-left movements
		{-2, -1}, {-1, -2},
	}
}

// Symbol is the Knight's symbol
func (k *Knight) Symbol() string {
	return "♞"
}

// NewKnight returns a Knight-flavored Piece
func NewKnight(board *Board, x, y uint) *Piece {
	piece := Piece{
		Board: board,
		XPos:  int(x),
		YPos:  int(y),
	}
	piece.PieceInterface = &Knight{Piece: &piece}
	return &piece
}
