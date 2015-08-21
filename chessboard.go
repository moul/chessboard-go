package chessboard

import "fmt"

var (
	ErrForbiddenCoordinates = fmt.Errorf("forbidden coordinates")
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

func (b *Board) NewVector() []bool {
	vector := make([]bool, b.GetSize())
	return vector
}

func (b *Board) ValidateCoordinates(xPos, yPos int) error {
	if xPos < 0 || xPos >= b.Length || yPos < 0 || yPos >= b.Height {
		return ErrForbiddenCoordinates
	}
	return nil
}

func (b *Board) CoordinatesToIndex(xPos, yPos, xShift, yShift int) (int, error) {
	xTarget := xPos + xShift
	yTarget := yPos + yShift

	err := b.ValidateCoordinates(xTarget, yTarget)
	if err != nil {
		return 0, err
	}

	return yTarget*b.Length + xTarget, nil
}
