package chessboard

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestBoard(t *testing.T) {
	Convey("Testing Board", t, func() {
		Convey("Board.GetPositions()", func() {
			board := NewBoard(3, 3)
			expected := []Coordinates{
				{0, 0}, {1, 0}, {2, 0},
				{0, 1}, {1, 1}, {2, 1},
				{0, 2}, {1, 2}, {2, 2},
			}
			So(board.GetPositions(), ShouldResemble, expected)

			board = NewBoard(4, 3)
			expected = []Coordinates{
				{0, 0}, {1, 0}, {2, 0}, {3, 0},
				{0, 1}, {1, 1}, {2, 1}, {3, 1},
				{0, 2}, {1, 2}, {2, 2}, {3, 2},
			}
			So(board.GetPositions(), ShouldResemble, expected)
		})

		Convey("Board.CoordinatesToIndex()", func() {
			board := NewBoard(3, 3)

			Convey("valid data", func() {
				index, err := board.CoordinatesToIndex(Coordinates{0, 0})
				So(err, ShouldBeNil)
				So(index, ShouldEqual, 0)

				index, err = board.CoordinatesToIndex(Coordinates{1, 1})
				So(err, ShouldBeNil)
				So(index, ShouldEqual, 4)

				index, err = board.CoordinatesToIndex(Coordinates{2, 2})
				So(err, ShouldBeNil)
				So(index, ShouldEqual, 8)
			})

			Convey("invalid data", func() {
				index, err := board.CoordinatesToIndex(Coordinates{-1, 0})
				So(err, ShouldNotBeNil)
				So(index, ShouldEqual, -1)

				index, err = board.CoordinatesToIndex(Coordinates{0, -1})
				So(err, ShouldNotBeNil)
				So(index, ShouldEqual, -1)

				index, err = board.CoordinatesToIndex(Coordinates{3, 0})
				So(err, ShouldNotBeNil)
				So(index, ShouldEqual, -1)

				index, err = board.CoordinatesToIndex(Coordinates{0, 3})
				So(err, ShouldNotBeNil)
				So(index, ShouldEqual, -1)
			})
		})

		Convey("Board.IndexToCoordinates()", func() {
			Convey("valid data (3x3)", func() {
				board := NewBoard(3, 3)
				coordinates, err := board.IndexToCoordinates(0)
				So(err, ShouldBeNil)
				So(coordinates, ShouldResemble, Coordinates{0, 0})

				coordinates, err = board.IndexToCoordinates(1)
				So(err, ShouldBeNil)
				So(coordinates, ShouldResemble, Coordinates{1, 0})

				coordinates, err = board.IndexToCoordinates(2)
				So(err, ShouldBeNil)
				So(coordinates, ShouldResemble, Coordinates{2, 0})

				coordinates, err = board.IndexToCoordinates(3)
				So(err, ShouldBeNil)
				So(coordinates, ShouldResemble, Coordinates{0, 1})

				coordinates, err = board.IndexToCoordinates(4)
				So(err, ShouldBeNil)
				So(coordinates, ShouldResemble, Coordinates{1, 1})

				coordinates, err = board.IndexToCoordinates(5)
				So(err, ShouldBeNil)
				So(coordinates, ShouldResemble, Coordinates{2, 1})

				coordinates, err = board.IndexToCoordinates(6)
				So(err, ShouldBeNil)
				So(coordinates, ShouldResemble, Coordinates{0, 2})

				coordinates, err = board.IndexToCoordinates(7)
				So(err, ShouldBeNil)
				So(coordinates, ShouldResemble, Coordinates{1, 2})

				coordinates, err = board.IndexToCoordinates(8)
				So(err, ShouldBeNil)
				So(coordinates, ShouldResemble, Coordinates{2, 2})
			})

			Convey("valid data - wide (1x4)", func() {
				board := NewBoard(1, 4)
				coordinates, err := board.IndexToCoordinates(0)
				So(err, ShouldBeNil)
				So(coordinates, ShouldResemble, Coordinates{0, 0})

				coordinates, err = board.IndexToCoordinates(1)
				So(err, ShouldBeNil)
				So(coordinates, ShouldResemble, Coordinates{0, 1})

				coordinates, err = board.IndexToCoordinates(2)
				So(err, ShouldBeNil)
				So(coordinates, ShouldResemble, Coordinates{0, 2})

				coordinates, err = board.IndexToCoordinates(3)
				So(err, ShouldBeNil)
				So(coordinates, ShouldResemble, Coordinates{0, 3})
			})

			Convey("valid data - long (4x1)", func() {
				board := NewBoard(4, 1)
				coordinates, err := board.IndexToCoordinates(0)
				So(err, ShouldBeNil)
				So(coordinates, ShouldResemble, Coordinates{0, 0})

				coordinates, err = board.IndexToCoordinates(1)
				So(err, ShouldBeNil)
				So(coordinates, ShouldResemble, Coordinates{1, 0})

				coordinates, err = board.IndexToCoordinates(2)
				So(err, ShouldBeNil)
				So(coordinates, ShouldResemble, Coordinates{2, 0})

				coordinates, err = board.IndexToCoordinates(3)
				So(err, ShouldBeNil)
				So(coordinates, ShouldResemble, Coordinates{3, 0})
			})

			Convey("invalid data", func() {
				board := NewBoard(3, 3)
				coordinates, err := board.IndexToCoordinates(-1)
				So(err, ShouldNotBeNil)
				So(coordinates, ShouldResemble, Coordinates{-1, -1})

				coordinates, err = board.IndexToCoordinates(9)
				So(err, ShouldNotBeNil)
				So(coordinates, ShouldResemble, Coordinates{-1, -1})
			})
		})
	})
}
