package main

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestMaingenTest(t *testing.T) {
	Convey("genTest", t, func() {
		var (
			parses = []*parse{}
		)
		Convey("When everything goes positive", func() {
			err := genTest(parses)
			Convey("Then err should be nil.", func() {
				So(err, ShouldBeNil)
			})
		})
	})
}

func TestMaingenUTTest(t *testing.T) {
	Convey("genUTTest", t, func() {
		Convey("When everything goes positive", func() {
			err := s.genUTTest()
			Convey("Then err should be nil.", func() {
				So(err, ShouldBeNil)
			})
		})
	})
}

func TestMaingenTestMain(t *testing.T) {
	Convey("genTestMain", t, func() {
		Convey("When everything goes positive", func() {
			err := s.genTestMain()
			Convey("Then err should be nil.", func() {
				So(err, ShouldBeNil)
			})
		})
	})
}

func TestMaingenInterface(t *testing.T) {
	Convey("genInterface", t, func() {
		var (
			parses = []*parse{}
		)
		Convey("When everything goes positive", func() {
			err := genInterface(parses)
			Convey("Then err should be nil.", func() {
				So(err, ShouldBeNil)
			})
		})
	})
}

func TestMaingenMock(t *testing.T) {
	Convey("genMock", t, func() {
		var (
			files = ""
		)
		Convey("When everything goes positive", func() {
			err := genMock(files)
			Convey("Then err should be nil.", func() {
				So(err, ShouldBeNil)
			})
		})
	})
}

func TestMaingenMonkey(t *testing.T) {
	Convey("genMonkey", t, func() {
		var (
			parses = []*parse{}
		)
		Convey("When everything goes positive", func() {
			err := genMonkey(parses)
			Convey("Then err should be nil.", func() {
				So(err, ShouldBeNil)
			})
		})
	})
}
