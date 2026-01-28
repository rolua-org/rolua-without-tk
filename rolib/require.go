package rolib

import (
	"rolua/rolua"

	rt "github.com/arnodel/golua/runtime"
	"github.com/traefik/yaegi/interp"
)

func require(t *rt.Thread, gc *rt.GoCont) (rt.Cont, error) {
	path, err := gc.StringArg(0)
	if err != nil {
		return nil, err
	}

	i := interp.New(interp.Options{})
	i.Use(rolua.Symbols)

	_, goErr := i.EvalPath(path)
	if goErr != nil {
		panic(goErr)
	}

	v, goErr := i.Eval("rolib.Init()")
	if goErr != nil {
		panic(goErr)
	}

	pkg := v.Interface().(*rt.Table)
	return gc.PushingNext(t.Runtime, rt.TableValue(pkg)), nil
}
