package main

import (
	"rolua/rolib"

	"github.com/arnodel/golua/lib/base"
	"github.com/arnodel/golua/lib/coroutine"
	"github.com/arnodel/golua/lib/debuglib"
	"github.com/arnodel/golua/lib/golib"
	"github.com/arnodel/golua/lib/iolib"
	"github.com/arnodel/golua/lib/mathlib"
	"github.com/arnodel/golua/lib/oslib"
	"github.com/arnodel/golua/lib/packagelib"
	"github.com/arnodel/golua/lib/runtimelib"
	"github.com/arnodel/golua/lib/stringlib"
	"github.com/arnodel/golua/lib/tablelib"
	"github.com/arnodel/golua/lib/utf8lib"
)

var libs = []packagelib.Loader{
	base.LibLoader,
	packagelib.LibLoader,
	coroutine.LibLoader,
	stringlib.LibLoader,
	tablelib.LibLoader,
	mathlib.LibLoader,
	iolib.LibLoader,
	utf8lib.LibLoader,
	oslib.LibLoader,
	debuglib.LibLoader,
	runtimelib.LibLoader,
	golib.LibLoader,

	rolib.LibLoader,
}
