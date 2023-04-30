package main

import "os"

func main() {
	if len(os.Args) > 1 && os.Args[1] != "" {
		lines := fileSystem()       //---			---> Collect Lines from [BANNER] File (os.Args[2])
		words := Format(os.Args[1]) //---			---> Manage Printability & "\n" &  Emptiness
		wordArt(words, lines)       //---			---> Print Letter of Argument from File Lines
	}
}
