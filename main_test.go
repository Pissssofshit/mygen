package main

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestMainmain(t *testing.T) {
	Convey("main", t, func() {
		Convey("When everything goes positive", func() {
			main()
			Convey("No return values", func() {
			})
		})
	})
}
