package src

import (
	"ascii/tools"
	"fmt"
	"log"
	"os"
	"strings"
)

func Ascii_Art_Color(color string) {
	if tools.ColorChoice(color) == 0 {
		log.Fatal("\033[31m\nERROR: \033[0mInvalid Color")
	}
	lines := tools.FileScan("templates/standard.txt")
	words, err_Format := tools.ColorFormat(os.Args)
	if err_Format != nil {
		log.Fatal(err_Format)
	}
	//---> Print Argument's Words' Letters using Scanned File
	for _, word := range words {
		if word == "" {
			fmt.Println()
		} else {
			for i := 1; i < 9; i++ {
				for _, letter := range word {
					if len(os.Args) == 4 {
						if strings.Contains(os.Args[2], string(letter)) {
							fmt.Printf("\033[%dm%v\033[0m", tools.ColorChoice(color), lines[int(letter-32)*9+i])
						} else {
							fmt.Print(lines[int(letter-32)*9+i])
						}
					} else {
						fmt.Printf("\033[%dm%v\033[0m", tools.ColorChoice(color), lines[int(letter-32)*9+i])
					}
				}
				fmt.Println()
			}
		}
	}
}
