package tools

import (
	"log"
	"strings"
)

func StringFormat(s string) []string {
	//---> Check Printability
	r := []rune(s)
	if len(r) != len(s) {
		log.Fatal("\033[31m\nERROR: \033[0mNon-Printable Detected")
	}
	words := strings.Split(s, "\\n") //---> Manage "/n"
	//---> Check Emptiness
	for _, void := range words {
		if void != "" {
			return words
		}
	}
	return words[1:]
}
