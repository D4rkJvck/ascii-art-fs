package tools

import "errors"

func ColorFormat(t []string) ([]string, error) {
	err := errors.New("\nUsage: go run . [OPTION] [STRING]\n\nEX: go run . --color=<color> <letters to be colored> something")
	//---> Lenght Check
	if len(t) < 3 || len(t) > 4 {
		return nil, err
	}
	if len(t) == 4 {
		return StringFormat(t[3]), nil
	}
	return StringFormat(t[2]), nil
}
