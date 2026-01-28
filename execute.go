package main

import (
	"os"

	"github.com/arnodel/golua/lib"
	"github.com/arnodel/golua/lib/base"
	rt "github.com/arnodel/golua/runtime"
)

func execute(entry string, args []string) error {
	r := rt.New(nil)
	base.Load(r)
	cleanupLib := lib.LoadLibs(r, libs...)

	defer cleanupLib()
	defer r.Close(nil)

	chunk, err := os.ReadFile(entry)
	if err != nil {
		return err
	}

	var argVals []rt.Value
	if len(args) > 0 {
		argTable := rt.NewTable()
		argVals = make([]rt.Value, len(args))

		for idx, arg := range args {
			argVal := rt.StringValue(arg)
			r.SetTable(argTable, rt.IntValue(int64(idx+1)), argVal)
			argVals[idx] = argVal
		}

		r.SetTable(r.GlobalEnv(), rt.StringValue("arg"), rt.TableValue(argTable))
	}

	clos, err := r.LoadFromSourceOrCode(entry, chunk, "bt", rt.TableValue(r.GlobalEnv()), true)
	if err != nil {
		return err
	}

	cerr := rt.Call(r.MainThread(), rt.FunctionValue(clos), argVals, rt.NewTerminationWith(nil, 0, false))
	if cerr != nil {
		return cerr
	}

	return nil
}
