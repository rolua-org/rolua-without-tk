package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Usage: rolua <file> [args]")
		return
	}

	if err := execute(os.Args[1], os.Args[2:]); err != nil {
		panic(err)
	}
}
