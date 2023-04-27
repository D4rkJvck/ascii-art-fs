package main

import "fmt"

func wordArt(t, l []string) {
	for _, word := range t {
		if word == "" { //---			---> From Argument Split
			fmt.Println("")
		} else {
			for i := 1; i < 9; i++ { //---			---> Line Loop
				for _, letter := range word { //---			---> Word Loop
					fmt.Print(l[lineAbove(letter)+i]) //---> Add each (i)th Line after Line ABOVE
				}
				fmt.Println() //---			---> Next Letter Line
			}
		}
	}
}

func lineAbove(r rune) int {
	return int(r-32) * 9
}