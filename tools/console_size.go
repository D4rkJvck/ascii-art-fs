package tools

import (
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func ConsoleSize() int {
	cmd := exec.Command("tput", "cols")
	cmd.Stdin = os.Stdin
	size, err := cmd.Output()
	if err != nil {
		log.Fatalf("%v", err)
	}
	width, _ := strconv.Atoi(strings.TrimSpace(string(size)))
	return width
}
