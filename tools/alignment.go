package tools

import (
	"log"
	"strings"
)

func Alignment(t, s string) string {
	width := ConsoleSize()
	size := len(s)
	void := width - size
	if void >= 0 {
		switch t {
		case "left":
			return s
		case "right":
			return strings.Repeat(" ", void) + s
		case "center":
			return strings.Repeat(" ", void/2) + s
		case "justify":
			tab := strings.Split(s, "#")
			void += len(tab) - 1
			size -= len(tab) - 1
			s = tab[0]
			if len(tab) > 1 {
				for _, str := range tab[1:] {
					if void%len(tab[1:]) != 0 {
						s += " "
						void--
					}
					s += strings.Repeat(" ", void/(len(tab[1:]))) + str
				}
			}
			return s
		default:
			log.Fatal("\nUsage: go run . [OPTION] [STRING] [BANNER]\n\nEX: go run . --align=\033[31mcenter\033[0m something standard")
		}
	} else {
		log.Fatal("\033[31m\nERROR: \033[0mCannot fit Terminal size")
	}
	return s
}
