package src

import (
	"log"
	"os"
	"regexp"
)

func Input() {
	//---> Check Flag Pattern Matchbb
	pattern := regexp.MustCompile(`^--(.*)=(.*)`)
	if pattern.MatchString(os.Args[1]) {
		sub := pattern.FindStringSubmatch(os.Args[1])
		//---> Choose Options depending on Flags
		switch sub[1] {
		case "output":
			Ascii_Art_OutPut(sub[2])
		case "color":
			Ascii_Art_Color(sub[2])
		case "align":
			Ascii_Art_Justify(sub[2])
		default:
			log.Fatal("\033[31m\nERROR: \033[0mInvalid [OPTION]")
		}
	} else {
		Ascii_Art_FS() //---> No Flag Option
	}
}
