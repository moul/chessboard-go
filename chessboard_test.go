package chessboard

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestChessboard(t *testing.T) {
	Convey("Testing Chessboard", t, func() {
		Convey("Board 3x3, 1 king, 2 knights", func() {
			board := NewChessboard(3, 3)
			board.Pieces[KingPieceType] = 1
			board.Pieces[KnightPieceType] = 2

			results, err := board.Solve()
			fmt.Println(results)
			So(err, ShouldBeNil)
			So(len(results), ShouldEqual, 12)
		})
	})
}
