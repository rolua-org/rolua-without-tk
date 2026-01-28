package rolib

import (
	"github.com/arnodel/golua/lib/packagelib"
	rt "github.com/arnodel/golua/runtime"
)

var LibLoader = packagelib.Loader{
	Name: "rolib",
	Load: load,
}

func load(r *rt.Runtime) (rt.Value, func()) {
	pkg := rt.NewTable()

	r.SetEnvGoFunc(pkg, "require", require, 1, false)

	return rt.TableValue(pkg), nil
}
