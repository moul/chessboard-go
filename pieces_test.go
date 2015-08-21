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
