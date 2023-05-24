package tools

import (
	"bufio"
	"log"
	"os"
)

func FileScan(s string) []string {
	var lines []string
	//---> Open File
	content, err := os.Open(s)
	if err != nil {
		log.Fatalf("\n%v", err)
	}
	defer content.Close()
	//---> Scan File
	scanner := bufio.NewScanner(content)
	for scanner.Scan() {
		lines = append(lines, scanner.Text()) //---> Collect File Lines
	}
	if len(lines) != 855 {
		log.Fatal("\033[31m\nERROR: \033[0mInvalid File")
	}
	return lines
}
