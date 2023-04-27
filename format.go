package main

import (
	"os"
	"strings"
)

func Format(s string) []string {
	Printability(s) //---			---> Check if there is any Non-Printable Element
	var words []string
	if s == "\\n" {
		return []string{""} //---			---> "\n" as Argument
	} else if strings.Contains(s, "\\n") { 
		words = strings.Split(s, "\\n") //---			---> "\n" as Contained in Argument
	} else {
		words = append(words, os.Args[1]) //---			---> "\n" as Absent from Argument
	}
	return Emptiness(words) //---			---> Check if there is only Empty String in Table
}

func Printability(s string) {
	r := []rune(s)
	if len(r) != len(s) { //---			---> Given that Non-Printable is represented by two characters 
		os.Stderr.WriteString("\033[31mERROR --> \033[30mNon-Printable Detected\033[0m")
		os.Exit(0)
	}
}

func Emptiness(t []string) []string {
	for _, v := range t {
		if v != "" { //---			---> Non-Empty string = Non-Empty Table
			return t
		}
	}
	return t[1:]
}
