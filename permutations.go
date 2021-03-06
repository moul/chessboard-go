package chessboard

import "sync"

type Permutation struct {
	Pieces    []PieceType
	Depth     int
	RangeSize int
	Indexes   map[int]int
}

func NewPermutation(pieces map[PieceType]int, rangeSize int) *Permutation {
	permutation := Permutation{
		Pieces:  make([]PieceType, 0),
		Indexes: make(map[int]int),
	}

	for kind, quantity := range pieces {
		for ; quantity > 0; quantity-- {
			permutation.Pieces = append(permutation.Pieces, kind)
		}
	}

	permutation.Depth = len(permutation.Pieces)

	if rangeSize != 0 {
		permutation.RangeSize = rangeSize
	} else {
		permutation.RangeSize = permutation.Depth
	}

	return &permutation
}

func (p *Permutation) Generator() <-chan string {
	c := make(chan string, p.RangeSize)

	go func() {
		defer close(c)

		if len(p.Pieces) == 0 {
			return
		}

		var wg sync.WaitGroup
		wg.Add(p.RangeSize)

		for i := 1; i <= p.RangeSize; i++ {
			go func(i int) {
				// Permute()
				wg.Done()
			}(i)
		}

		wg.Wait()
	}()

	return c
}
