package main

import "os"

func main() {

	if len(os.Args) > 1 && os.Args[1] != "" {
		words := Format(os.Args[1]) //---			---> Manage Printability & "\n" &  Emptiness
		lines := fileSystem()       //---			---> Collect Lines from File given as Second Argument
		wordArt(words, lines)       //---			---> Print Letter of Argument from File Lines
	}
}
