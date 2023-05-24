package src

import (
	"ascii/tools"
	"fmt"
	"os"
)

func Ascii_Art_Justify(tipo string) {
	var art string
	lines := tools.FileSystem(os.Args, 3, 4)
	words := tools.StringFormat(os.Args[2])
	for _, word := range words {
		if word == "" {
			fmt.Println()
		} else {
			for i := 1; i < 9; i++ {
				for _, letter := range word {
					if tipo == "justify" && letter == 32 {
						art += "#"
					} else {
						art += lines[int(letter-32)*9+i]
					}
				}
				art = tools.Alignment(tipo, art)
				fmt.Println(art)
				art = ""
			}
		}
	}
}
