package chessboard

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestPermutation(t *testing.T) {
	Convey("Testing Permutation", t, func() {
		Convey("NewPermutation", func() {
			permutation := NewPermutation(map[PieceType]int{
				KingPieceType:   1,
				KnightPieceType: 2,
			}, 0)
			fmt.Println(permutation.Pieces)
			So(permutation.Pieces, ShouldResemble, []PieceType{KingPieceType, KnightPieceType, KnightPieceType})
			So(permutation.Depth, ShouldEqual, 3)
			So(permutation.RangeSize, ShouldEqual, 3)
		})
	})
}
