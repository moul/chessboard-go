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
			So(permutation.Pieces, ShouldResemble, []PieceType{KingPieceType, KnightPieceType, KnightPieceType})
			So(permutation.Depth, ShouldEqual, 3)
			So(permutation.RangeSize, ShouldEqual, 3)
		})
		Convey("Testing generator", func() {
			permutation := NewPermutation(map[PieceType]int{
				KingPieceType: 1,
			}, 3)
			So(permutation.Pieces, ShouldResemble, []PieceType{KingPieceType})
			So(permutation.Depth, ShouldEqual, 1)
			So(permutation.RangeSize, ShouldEqual, 3)
			for perm := range permutation.Generator() {
				fmt.Println(perm)
			}
		})
	})
}
