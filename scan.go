package main

import (
	"bufio"
	"os"
)

func fileScan(s string) []string {
	var lines []string
	content, err := os.Open(s) //---			---> Before Scan
	if err != nil {
		os.Stderr.WriteString(err.Error()) //---			---> Example: File Missing
		os.Exit(0)
	}
	scanner := bufio.NewScanner(content)
	for scanner.Scan() {
		lines = append(lines, scanner.Text()) //---			---> Retrieve File Content Line by Line
	}
	return lines
}
