package main

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestMainGoImport(t *testing.T) {
	Convey("GoImport", t, func() {
		var (
			file  = ""
			bytes = []byte("")
		)
		Convey("When everything goes positive", func() {
			res, err := GoImport(file, bytes)
			Convey("Then err should be nil.res should not be nil.", func() {
				So(err, ShouldBeNil)
				So(res, ShouldNotBeNil)
			})
		})
	})
}

func TestMainConvertMethod(t *testing.T) {
	Convey("ConvertMethod", t, func() {
		var (
			path = ""
		)
		Convey("When everything goes positive", func() {
			method := ConvertMethod(path)
			Convey("Then method should not be nil.", func() {
				So(method, ShouldNotBeNil)
			})
		})
	})
}

func TestMainConvertHump(t *testing.T) {
	Convey("ConvertHump", t, func() {
		var (
			words = ""
		)
		Convey("When everything goes positive", func() {
			p1 := ConvertHump(words)
			Convey("Then p1 should not be nil.", func() {
				So(p1, ShouldNotBeNil)
			})
		})
	})
}
