package src

import (
	"ascii/tools"
	"log"
	"os"
	"strings"
)

func Ascii_Art_OutPut(file string) {
	if !strings.HasSuffix(file, ".txt") { //---> File Format Condition
		log.Fatal("\nUsage: go run . [OPTION] [STRING] [BANNER]\n\nEX: go run . --output=<fileName\033[31m.txt\033[0m> something standard")
	}
	lines := tools.FileSystem(os.Args, 3, 4)
	words := tools.StringFormat(os.Args[2])
	var out string
	//---> Collect Argument's Words' Letters using File Pattern
	for _, word := range words {
		if word == "" {
			out += "\n"
		} else {
			for i := 1; i < 9; i++ {
				for _, letter := range word {
					out += lines[int(letter-32)*9+i]
				}
				out += "\n"
			}
		}
	}
	os.WriteFile("data/"+file, []byte(out), 0700) //---> Place Collected Letters' Art in File
}
