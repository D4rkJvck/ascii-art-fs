package tools

import (
	"log"
	"strconv"
	"strings"
)

func ColorChoiceAnsi(s string) int {
	switch s {
	case "black":
		return 30
	case "red":
		return 31
	case "green":
		return 32
	case "yellow":
		return 33
	case "blue":
		return 34
	case "magenta":
		return 35
	case "cyan":
		return 36
	case "white":
		return 37
	case "gray":
		return 90
	case "lightgreen":
		return 92
	case "lightblue":
		return 94
	case "pink":
		return 95
	default:
		return 0
	}
}

func ColorChoiceRgb(s string) (r, g, b int) {
	// Supprimer les caractères inutiles autour de la chaîne RGB
	s = strings.TrimSpace(s)
	s = strings.TrimPrefix(s, "rgb(")
	s = strings.TrimSuffix(s, ")")

	// Séparer les composantes R, G et B
	components := strings.Split(s, ";")
	if len(components) != 3 {
		log.Fatal("\033[31m\nERROR: \033[0mInvalid RGB Color")
	}

	r, err := strconv.Atoi(components[0])
	if err != nil || (r < 0 || r > 255) {
		log.Fatal("\033[31m\nERROR: \033[0mInvalid RGB Color")
	}

	g, err = strconv.Atoi(components[1])
	if err != nil ||(g < 0 || g > 255) {
		log.Fatal("\033[31m\nERROR: \033[0mInvalid RGB Color")
	}

	b, err = strconv.Atoi(components[2])
	if err != nil || (b < 0 || b > 255) {
		log.Fatal("\033[31m\nERROR: \033[0mInvalid RGB Color")
	}

	return
}
