package main

import (
	"ascii/src"
	"os"
)

func main() {
	if len(os.Args) > 1 {
		src.Input()
	}
}
