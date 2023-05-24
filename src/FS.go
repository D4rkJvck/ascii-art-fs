package src

import (
	"ascii/tools"
	"fmt"
	"os"
)

func Ascii_Art_FS() {
	lines := tools.FileSystem(os.Args, 2, 3)
	words := tools.StringFormat(os.Args[1])
	//---> Print Argument's Words' Letters using File Pattern
	for _, word := range words {
		if word == "" {
			fmt.Println()
		} else {
			for i := 1; i < 9; i++ {
				for _, letter := range word {
					fmt.Print(lines[int(letter-32)*9+i])
				}
				fmt.Println()
			}
		}
	}
}
