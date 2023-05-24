package tools

func ColorChoice(s string) int {
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
