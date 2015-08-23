package chessboard

import "fmt"

var (
	ErrForbiddenCoordinates = fmt.Errorf("forbidden coordinates")
	ErrForbiddenIndex       = fmt.Errorf("forbidden index")
)

type Board struct {
	Length int
	Height int
}

func NewBoard(length, height uint) *Board {
	return &Board{
		Length: int(length),
		Height: int(height),
	}
}

func (b *Board) GetSize() int {
	return b.Length * b.Height
}

func (b *Board) NewVector() BoardVector {
	vector := make(BoardVector, b.GetSize())
	return vector
}

func (b *Board) ValidateCoordinates(coordinates Coordinates) error {
	if coordinates.XPos < 0 || coordinates.XPos >= b.Length || coordinates.YPos < 0 || coordinates.YPos >= b.Height {
		return ErrForbiddenCoordinates
	}
	return nil
}

func (b *Board) CoordinatesToIndex(coordinates Coordinates) (BoardIndex, error) {
	err := b.ValidateCoordinates(coordinates)
	if err != nil {
		return -1, err
	}

	index := BoardIndex(coordinates.YPos*b.Length + coordinates.XPos)
	return index, nil
}

func (b *Board) IndexToCoordinates(index BoardIndex) (Coordinates, error) {
	if index < 0 || int(index) >= b.GetSize() {
		return Coordinates{-1, -1}, ErrForbiddenIndex
	}

	x := int(index) % b.Length
	y := int(index) / b.Length
	return Coordinates{x, y}, nil
}

func (b *Board) GetPositions() []Coordinates {
	positions := []Coordinates{}
	for y := 0; y < b.Height; y++ {
		for x := 0; x < b.Length; x++ {
			positions = append(positions, Coordinates{x, y})
		}
	}
	return positions
}

type BoardVector []bool

type Coordinates struct {
	XPos int
	YPos int
}

type BoardIndex int
