package main

import (
	"go/ast"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestMainparseArgs(t *testing.T) {
	Convey("parseArgs", t, func() {
		var (
			args  = []string{}
			res   = &[]string{}
			index = int(0)
		)
		Convey("When everything goes positive", func() {
			err := parseArgs(args, res, index)
			Convey("Then err should be nil.", func() {
				So(err, ShouldBeNil)
			})
		})
	})
}

func TestMainparseFile(t *testing.T) {
	Convey("parseFile", t, func() {
		var (
			files = ""
		)
		Convey("When everything goes positive", func() {
			parses, err := parseFile(files)
			Convey("Then err should be nil.parses should not be nil.", func() {
				So(err, ShouldBeNil)
				So(parses, ShouldNotBeNil)
			})
		})
	})
}

func TestMainparserParams(t *testing.T) {
	Convey("parserParams", t, func() {
		var (
			fields = []*ast.Field{}
		)
		Convey("When everything goes positive", func() {
			params := parserParams(fields)
			Convey("Then params should not be nil.", func() {
				So(params, ShouldNotBeNil)
			})
		})
	})
}

func TestMainparseType(t *testing.T) {
	Convey("parseType", t, func() {
		var (
			expr ast.Expr
		)
		Convey("When everything goes positive", func() {
			p1 := parseType(expr)
			Convey("Then p1 should not be nil.", func() {
				So(p1, ShouldNotBeNil)
			})
		})
	})
}

func TestMainparseFuncType(t *testing.T) {
	Convey("parseFuncType", t, func() {
		var (
			temp = ""
			data = &ast.FieldList{}
		)
		Convey("When everything goes positive", func() {
			p1 := parseFuncType(temp, data)
			Convey("Then p1 should not be nil.", func() {
				So(p1, ShouldNotBeNil)
			})
		})
	})
}

func TestMainparseImports(t *testing.T) {
	Convey("parseImports", t, func() {
		var (
			specs = []ast.Spec{}
		)
		Convey("When everything goes positive", func() {
			params := parseImports(specs)
			Convey("Then params should not be nil.", func() {
				So(params, ShouldNotBeNil)
			})
		})
	})
}
