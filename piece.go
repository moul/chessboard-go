package chessboard

type PieceInterface interface {
	Movements() []PieceMovement
	Symbol() string
}

// Piece is a generic piece
type Piece struct {
	Board *Board
	XPos  int
	YPos  int

	PieceInterface
}

func (p *Piece) GetIndex() (int, error) {
	return p.Board.CoordinatesToIndex(p.XPos, p.YPos, 0, 0)
}

func (p *Piece) GetTerritory() ([]bool, error) {
	vector := p.Board.NewVector()

	index, err := p.GetIndex()
	if err != nil {
		return nil, err
	}

	vector[index] = true

	for _, movement := range p.PieceInterface.Movements() {
		index, err = p.Board.CoordinatesToIndex(p.XPos, p.YPos, movement.Horizontal, movement.Vertical)
		if err == nil {
			vector[index] = true
		}
	}

	return vector, nil
}

func (p *Piece) Movements() []PieceMovement {
	return nil
}

type PieceMovement struct {
	Horizontal int
	Vertical   int
}
