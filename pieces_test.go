package chessboard

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestKing(t *testing.T) {
	Convey("Testing King", t, func() {
		Convey("Board=3x3, King=1x1", func() {
			territory, err := NewKing(NewBoard(3, 3), 1, 1).GetTerritory()
			So(err, ShouldBeNil)
			So(territory, ShouldResemble, []bool{
				true, true, true,
				true, true, true,
				true, true, true,
			})
		})
		Convey("Board=3x3, King=0x0", func() {
			territory, err := NewKing(NewBoard(3, 3), 0, 0).GetTerritory()
			So(err, ShouldBeNil)
			So(territory, ShouldResemble, []bool{
				true, true, false,
				true, true, false,
				false, false, false,
			})
		})
		Convey("Board=3x3, King=1x0", func() {
			territory, err := NewKing(NewBoard(3, 3), 1, 0).GetTerritory()
			So(err, ShouldBeNil)
			So(territory, ShouldResemble, []bool{
				true, true, true,
				true, true, true,
				false, false, false,
			})
		})
		Convey("Board=3x3, King=2x0", func() {
			territory, err := NewKing(NewBoard(3, 3), 2, 0).GetTerritory()
			So(err, ShouldBeNil)
			So(territory, ShouldResemble, []bool{
				false, true, true,
				false, true, true,
				false, false, false,
			})
		})
		Convey("Board=3x3, King=2x1", func() {
			territory, err := NewKing(NewBoard(3, 3), 2, 1).GetTerritory()
			So(err, ShouldBeNil)
			So(territory, ShouldResemble, []bool{
				false, true, true,
				false, true, true,
				false, true, true,
			})
		})
		Convey("Board=3x3, King=2x2", func() {
			territory, err := NewKing(NewBoard(3, 3), 2, 2).GetTerritory()
			So(err, ShouldBeNil)
			So(territory, ShouldResemble, []bool{
				false, false, false,
				false, true, true,
				false, true, true,
			})
		})
		Convey("Board=3x3, King=1x2", func() {
			territory, err := NewKing(NewBoard(3, 3), 1, 2).GetTerritory()
			So(err, ShouldBeNil)
			So(territory, ShouldResemble, []bool{
				false, false, false,
				true, true, true,
				true, true, true,
			})
		})
		Convey("Board=3x3, King=0x2", func() {
			territory, err := NewKing(NewBoard(3, 3), 0, 2).GetTerritory()
			So(err, ShouldBeNil)
			So(territory, ShouldResemble, []bool{
				false, false, false,
				true, true, false,
				true, true, false,
			})
		})
		Convey("Board=3x3, King=0x1", func() {
			territory, err := NewKing(NewBoard(3, 3), 0, 1).GetTerritory()
			So(err, ShouldBeNil)
			So(territory, ShouldResemble, []bool{
				true, true, false,
				true, true, false,
				true, true, false,
			})
		})
	})
}

func TestQueen(t *testing.T) {
	Convey("Testing Queen", t, func() {
		Convey("Board=3x3, Queen=1x1", func() {
			territory, err := NewQueen(NewBoard(3, 3), 1, 1).GetTerritory()
			So(err, ShouldBeNil)
			So(territory, ShouldResemble, []bool{
				true, true, true,
				true, true, true,
				true, true, true,
			})
		})
		Convey("Board=3x3, Queen=0x0", func() {
			territory, err := NewQueen(NewBoard(3, 3), 0, 0).GetTerritory()
			So(err, ShouldBeNil)
			So(territory, ShouldResemble, []bool{
				true, true, true,
				true, true, false,
				true, false, true,
			})
		})
		Convey("Board=3x3, Queen=1x0", func() {
			territory, err := NewQueen(NewBoard(3, 3), 1, 0).GetTerritory()
			So(err, ShouldBeNil)
			So(territory, ShouldResemble, []bool{
				true, true, true,
				true, true, true,
				false, true, false,
			})
		})
		Convey("Board=3x3, Queen=2x0", func() {
			territory, err := NewQueen(NewBoard(3, 3), 2, 0).GetTerritory()
			So(err, ShouldBeNil)
			So(territory, ShouldResemble, []bool{
				true, true, true,
				false, true, true,
				true, false, true,
			})
		})
		Convey("Board=3x3, Queen=2x1", func() {
			territory, err := NewQueen(NewBoard(3, 3), 2, 1).GetTerritory()
			So(err, ShouldBeNil)
			So(territory, ShouldResemble, []bool{
				false, true, true,
				true, true, true,
				false, true, true,
			})
		})
		Convey("Board=3x3, Queen=2x2", func() {
			territory, err := NewQueen(NewBoard(3, 3), 2, 2).GetTerritory()
			So(err, ShouldBeNil)
			So(territory, ShouldResemble, []bool{
				true, false, true,
				false, true, true,
				true, true, true,
			})
		})
		Convey("Board=3x3, Queen=1x2", func() {
			territory, err := NewQueen(NewBoard(3, 3), 1, 2).GetTerritory()
			So(err, ShouldBeNil)
			So(territory, ShouldResemble, []bool{
				false, true, false,
				true, true, true,
				true, true, true,
			})
		})
		Convey("Board=3x3, Queen=0x2", func() {
			territory, err := NewQueen(NewBoard(3, 3), 0, 2).GetTerritory()
			So(err, ShouldBeNil)
			So(territory, ShouldResemble, []bool{
				true, false, true,
				true, true, false,
				true, true, true,
			})
		})
		Convey("Board=3x3, Queen=0x1", func() {
			territory, err := NewQueen(NewBoard(3, 3), 0, 1).GetTerritory()
			So(err, ShouldBeNil)
			So(territory, ShouldResemble, []bool{
				true, true, false,
				true, true, true,
				true, true, false,
			})
		})
	})
}

func TestRook(t *testing.T) {
	Convey("Testing Rook", t, func() {
		Convey("Board=3x3, Rook=1x1", func() {
			territory, err := NewRook(NewBoard(3, 3), 1, 1).GetTerritory()
			So(err, ShouldBeNil)
			So(territory, ShouldResemble, []bool{
				false, true, false,
				true, true, true,
				false, true, false,
			})
		})
		Convey("Board=3x3, Rook=0x0", func() {
			territory, err := NewRook(NewBoard(3, 3), 0, 0).GetTerritory()
			So(err, ShouldBeNil)
			So(territory, ShouldResemble, []bool{
				true, true, true,
				true, false, false,
				true, false, false,
			})
		})
		Convey("Board=3x3, Rook=1x0", func() {
			territory, err := NewRook(NewBoard(3, 3), 1, 0).GetTerritory()
			So(err, ShouldBeNil)
			So(territory, ShouldResemble, []bool{
				true, true, true,
				false, true, false,
				false, true, false,
			})
		})
		Convey("Board=3x3, Rook=2x0", func() {
			territory, err := NewRook(NewBoard(3, 3), 2, 0).GetTerritory()
			So(err, ShouldBeNil)
			So(territory, ShouldResemble, []bool{
				true, true, true,
				false, false, true,
				false, false, true,
			})
		})
		Convey("Board=3x3, Rook=2x1", func() {
			territory, err := NewRook(NewBoard(3, 3), 2, 1).GetTerritory()
			So(err, ShouldBeNil)
			So(territory, ShouldResemble, []bool{
				false, false, true,
				true, true, true,
				false, false, true,
			})
		})
		Convey("Board=3x3, Rook=2x2", func() {
			territory, err := NewRook(NewBoard(3, 3), 2, 2).GetTerritory()
			So(err, ShouldBeNil)
			So(territory, ShouldResemble, []bool{
				false, false, true,
				false, false, true,
				true, true, true,
			})
		})
		Convey("Board=3x3, Rook=1x2", func() {
			territory, err := NewRook(NewBoard(3, 3), 1, 2).GetTerritory()
			So(err, ShouldBeNil)
			So(territory, ShouldResemble, []bool{
				false, true, false,
				false, true, false,
				true, true, true,
			})
		})
		Convey("Board=3x3, Rook=0x2", func() {
			territory, err := NewRook(NewBoard(3, 3), 0, 2).GetTerritory()
			So(err, ShouldBeNil)
			So(territory, ShouldResemble, []bool{
				true, false, false,
				true, false, false,
				true, true, true,
			})
		})
		Convey("Board=3x3, Rook=0x1", func() {
			territory, err := NewRook(NewBoard(3, 3), 0, 1).GetTerritory()
			So(err, ShouldBeNil)
			So(territory, ShouldResemble, []bool{
				true, false, false,
				true, true, true,
				true, false, false,
			})
		})
	})
}

func TestBishop(t *testing.T) {
	Convey("Testing Bishop", t, func() {
		Convey("Board=3x3, Bishop=1x1", func() {
			territory, err := NewBishop(NewBoard(3, 3), 1, 1).GetTerritory()
			So(err, ShouldBeNil)
			So(territory, ShouldResemble, []bool{
				true, false, true,
				false, true, false,
				true, false, true,
			})
		})
		Convey("Board=3x3, Bishop=0x0", func() {
			territory, err := NewBishop(NewBoard(3, 3), 0, 0).GetTerritory()
			So(err, ShouldBeNil)
			So(territory, ShouldResemble, []bool{
				true, false, false,
				false, true, false,
				false, false, true,
			})
		})
		Convey("Board=3x3, Bishop=1x0", func() {
			territory, err := NewBishop(NewBoard(3, 3), 1, 0).GetTerritory()
			So(err, ShouldBeNil)
			So(territory, ShouldResemble, []bool{
				false, true, false,
				true, false, true,
				false, false, false,
			})
		})
		Convey("Board=3x3, Bishop=2x0", func() {
			territory, err := NewBishop(NewBoard(3, 3), 2, 0).GetTerritory()
			So(err, ShouldBeNil)
			So(territory, ShouldResemble, []bool{
				false, false, true,
				false, true, false,
				true, false, false,
			})
		})
		Convey("Board=3x3, Bishop=2x1", func() {
			territory, err := NewBishop(NewBoard(3, 3), 2, 1).GetTerritory()
			So(err, ShouldBeNil)
			So(territory, ShouldResemble, []bool{
				false, true, false,
				false, false, true,
				false, true, false,
			})
		})
		Convey("Board=3x3, Bishop=2x2", func() {
			territory, err := NewBishop(NewBoard(3, 3), 2, 2).GetTerritory()
			So(err, ShouldBeNil)
			So(territory, ShouldResemble, []bool{
				true, false, false,
				false, true, false,
				false, false, true,
			})
		})
		Convey("Board=3x3, Bishop=1x2", func() {
			territory, err := NewBishop(NewBoard(3, 3), 1, 2).GetTerritory()
			So(err, ShouldBeNil)
			So(territory, ShouldResemble, []bool{
				false, false, false,
				true, false, true,
				false, true, false,
			})
		})
		Convey("Board=3x3, Bishop=0x2", func() {
			territory, err := NewBishop(NewBoard(3, 3), 0, 2).GetTerritory()
			So(err, ShouldBeNil)
			So(territory, ShouldResemble, []bool{
				false, false, true,
				false, true, false,
				true, false, false,
			})
		})
		Convey("Board=3x3, Bishop=0x1", func() {
			territory, err := NewBishop(NewBoard(3, 3), 0, 1).GetTerritory()
			So(err, ShouldBeNil)
			So(territory, ShouldResemble, []bool{
				false, true, false,
				true, false, false,
				false, true, false,
			})
		})
	})
}

func TestKnight(t *testing.T) {
	Convey("Testing Knight", t, func() {
		Convey("Board=3x3, Knight=1x1", func() {
			territory, err := NewKnight(NewBoard(3, 3), 1, 1).GetTerritory()
			So(err, ShouldBeNil)
			So(territory, ShouldResemble, []bool{
				false, false, false,
				false, true, false,
				false, false, false,
			})
		})
		Convey("Board=3x3, Knight=0x0", func() {
			territory, err := NewKnight(NewBoard(3, 3), 0, 0).GetTerritory()
			So(err, ShouldBeNil)
			So(territory, ShouldResemble, []bool{
				true, false, false,
				false, false, true,
				false, true, false,
			})
		})
		Convey("Board=3x3, Knight=1x0", func() {
			territory, err := NewKnight(NewBoard(3, 3), 1, 0).GetTerritory()
			So(err, ShouldBeNil)
			So(territory, ShouldResemble, []bool{
				false, true, false,
				false, false, false,
				true, false, true,
			})
		})
		Convey("Board=3x3, Knight=2x0", func() {
			territory, err := NewKnight(NewBoard(3, 3), 2, 0).GetTerritory()
			So(err, ShouldBeNil)
			So(territory, ShouldResemble, []bool{
				false, false, true,
				true, false, false,
				false, true, false,
			})
		})
		Convey("Board=3x3, Knight=2x1", func() {
			territory, err := NewKnight(NewBoard(3, 3), 2, 1).GetTerritory()
			So(err, ShouldBeNil)
			So(territory, ShouldResemble, []bool{
				true, false, false,
				false, false, true,
				true, false, false,
			})
		})
		Convey("Board=3x3, Knight=2x2", func() {
			territory, err := NewKnight(NewBoard(3, 3), 2, 2).GetTerritory()
			So(err, ShouldBeNil)
			So(territory, ShouldResemble, []bool{
				false, true, false,
				true, false, false,
				false, false, true,
			})
		})
		Convey("Board=3x3, Knight=1x2", func() {
			territory, err := NewKnight(NewBoard(3, 3), 1, 2).GetTerritory()
			So(err, ShouldBeNil)
			So(territory, ShouldResemble, []bool{
				true, false, true,
				false, false, false,
				false, true, false,
			})
		})
		Convey("Board=3x3, Knight=0x2", func() {
			territory, err := NewKnight(NewBoard(3, 3), 0, 2).GetTerritory()
			So(err, ShouldBeNil)
			So(territory, ShouldResemble, []bool{
				false, true, false,
				false, false, true,
				true, false, false,
			})
		})
		Convey("Board=3x3, Knight=0x1", func() {
			territory, err := NewKnight(NewBoard(3, 3), 0, 1).GetTerritory()
			So(err, ShouldBeNil)
			So(territory, ShouldResemble, []bool{
				false, false, true,
				true, false, false,
				false, false, true,
			})
		})
	})
}
