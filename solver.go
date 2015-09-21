package chessboard

import "fmt"

type Solver struct {
	Length int
	Height int

	Pieces map[PieceType]int
}

func NewSolver(length, height uint) *Solver {
	return &Solver{
		Length: int(length),
		Height: int(height),
		Pieces: make(map[PieceType]int),
	}
}

func (cb *Solver) NewIndexVector() SolverVector {
	size := cb.GetVectorSize()
	vector := make(SolverVector, size)
	for i := 0; i < size; i++ {
		vector[i] = i
	}
	return vector
}

func (cb *Solver) GroupedPermutations() error {
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

func (cb *Solver) Solve() ([]Board, error) {
	solutions := []Board{}

	cb.GroupedPermutations()

	//for piecesSet := range cb.Tree()

	return solutions, nil
}

func (cb *Solver) GetVectorSize() int {
	return cb.Length * cb.Height
}

func (cb *Solver) NewVector() SolverVector {
	vector := make(SolverVector, cb.GetVectorSize())
	return vector
}

type SolverVector []int
