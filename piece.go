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

func (p *Piece) HorizontalMovements() []PieceMovement {
	movements := []PieceMovement{}
	for x := -p.LeftDistance(); x < p.RightDistance()+1; x++ {
		if x != 0 {
			movements = append(movements, PieceMovement{x, 0})
		}
	}
	return movements
}

func (p *Piece) VerticalMovements() []PieceMovement {
	movements := []PieceMovement{}
	for y := -p.TopDistance(); y < p.BottomDistance()+1; y++ {
		if y != 0 {
			movements = append(movements, PieceMovement{0, y})
		}
	}
	return movements
}

func (p *Piece) DiagonalMovements() []PieceMovement {
	movements := []PieceMovement{}

	leftDistance := p.LeftDistance()
	topDistance := p.TopDistance()
	rightDistance := p.RightDistance()
	bottomDistance := p.BottomDistance()

	leftTopMinDistance := min(leftDistance, topDistance)
	for i := 0; i < leftTopMinDistance; i++ {
		movements = append(movements, PieceMovement{-(i + 1), -(i + 1)})
	}

	rightTopMinDistance := min(rightDistance, topDistance)
	for i := 0; i < rightTopMinDistance; i++ {
		movements = append(movements, PieceMovement{+(i + 1), -(i + 1)})
	}

	rightBottomMinDistance := min(rightDistance, bottomDistance)
	for i := 0; i < rightBottomMinDistance; i++ {
		movements = append(movements, PieceMovement{+(i + 1), +(i + 1)})
	}

	leftBottomMinDistance := min(leftDistance, bottomDistance)
	for i := 0; i < leftBottomMinDistance; i++ {
		movements = append(movements, PieceMovement{-(i + 1), +(i + 1)})
	}

	return movements
}

func (p *Piece) LeftDistance() int {
	return p.XPos
}

func (p *Piece) TopDistance() int {
	return p.YPos
}

func (p *Piece) RightDistance() int {
	return p.Board.Length - p.XPos - 1
}

func (p *Piece) BottomDistance() int {
	return p.Board.Height - p.YPos
}

type PieceMovement struct {
	Horizontal int
	Vertical   int
}
