package chessboard

import "fmt"

type Chessboard struct {
	Length int
	Height int

	Pieces map[PieceType]int
}

func NewChessboard(length, height uint) *Chessboard {
	return &Chessboard{
		Length: int(length),
		Height: int(height),
		Pieces: make(map[PieceType]int),
	}
}

func (cb *Chessboard) NewIndexVector() ChessboardVector {
	size := cb.GetVectorSize()
	vector := make(ChessboardVector, size)
	for i := 0; i < size; i++ {
		vector[i] = i
	}
	return vector
}

func (cb *Chessboard) GroupedPermutations() error {
	for kind, quantity := range cb.Pieces {
		if quantity <= 0 {
			continue
		}
		vector := cb.NewIndexVector()
		for combination := range combinations(vector, quantity) {
			fmt.Println(combination)
		}
		fmt.Println(kind, quantity)
	}
	//fmt.Println(vector)
	return nil
}

func (cb *Chessboard) Solve() ([]Board, error) {
	solutions := []Board{}

	cb.GroupedPermutations()

	//for piecesSet := range cb.Tree()

	return solutions, nil
}

func (cb *Chessboard) GetVectorSize() int {
	return cb.Length * cb.Height
}

func (cb *Chessboard) NewVector() ChessboardVector {
	vector := make(ChessboardVector, cb.GetVectorSize())
	return vector
}

type ChessboardVector []int
