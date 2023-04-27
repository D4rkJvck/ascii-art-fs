package main

import (
	"os"
)

func fileSystem() []string {
	if len(os.Args) < 3 { //---			---> Default Banner
		return fileScan("standard.txt")
	} else if len(os.Args) > 3 { //---			---> Should be: "go run . [STRING] [BANNER]"
		os.Stderr.WriteString("Usage: go run . [STRING] [BANNER]\n\nEX: go run . something standard")
		os.Exit(0)
	} else if os.Args[2] != "standard" && os.Args[2] != "shadow" && os.Args[2] != "thinkertoy" { //---> Valid Banner
		os.Stderr.WriteString("\033[31mERROR --> \033[30mInvalid [BANNER]\n\033[32m[BANNER] --> \033[34mstandard || shadow || thinkertoy\033[0m")
		os.Exit(0)
	}
	return fileScan(os.Args[2] + ".txt") //---> User just need to put Banner Name
}
