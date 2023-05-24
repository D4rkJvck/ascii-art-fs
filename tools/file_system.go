package tools

import (
	"log"
)

func FileSystem(t []string, min, max int) []string {
	//---> Check Usage Format
	if len(t) < min || len(t) > max {
		log.Fatal("\nUsage: go run . [OPTION] [STRING] [BANNER]\n\nEX: go run . something standard")
	}
	//---> Default File
	if len(t) == min {
		return FileScan("templates/standard.txt")
	}
	//---> Check [BANNER]
	if t[max-1] != "standard" && t[max-1] != "shadow" && t[max-1] != "thinkertoy" && t[max-1] != "graffiti" && t[max-1] != "cards" {
		log.Fatal("\033[31m\nERROR: \033[0mInvalid [BANNER]")
	}
	//---> Scan File using Given File Name in Argument
	return FileScan("templates/" + t[max-1] + ".txt")
}
