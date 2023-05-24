package main

import (
	"ascii/tools"
	"os"
	"reflect"
	"testing"
)

func Test_1Arg(t *testing.T) {
	input := []string{os.Args[0]}
	expect := "\nUsage: go run . [OPTION] [STRING]\n\nEX: go run . --color=<color> <letters to be colored> something"
	output, err := tools.ColorFormat(input)
	if err == nil {
		t.Errorf("\nOutput: %v\nExpected: %v", output, expect)
	}
}

func Test_2Args(t *testing.T) {
	input := []string{os.Args[0], "--color=<color>"}
	expect := "\nUsage: go run . [OPTION] [STRING]\n\nEX: go run . --color=<color> <letters to be colored> something"
	output, err := tools.ColorFormat(input)
	if err == nil {
		t.Errorf("\nOutput: %v\nExpected: %v", output, expect)
	}
}

func Test_3Args(t *testing.T) {
	input := []string{os.Args[0], "--color=<color>", "anything\\nyou\\nwant"}
	expect := []string{"anything", "you", "want"}
	output, err := tools.ColorFormat(input)
	if err != nil {
		t.Errorf("\nOutput: %v\nExpected: %v", err, expect)
	}
	if !reflect.DeepEqual(output, expect) {
		t.Errorf("\nOutput: %v\nExpected: %v", output, expect)
	}
}

func Test_4Args(t *testing.T) {
	input := []string{os.Args[0], "--color=<color>", "<letters>", "anything\\nyou\\nwant"}
	expect := []string{"anything", "you", "want"}
	output, err := tools.ColorFormat(input)
	if err != nil {
		t.Errorf("\nOutput: %v\nExpected: %v", err, expect)
	}
	if !reflect.DeepEqual(output, expect) {
		t.Errorf("\nOutput: %v\nExpected: %v", output, expect)
	}
}

func Test_5Args(t *testing.T) {
	input := []string{os.Args[0], "--color=<color>", "<letters>", "hello", "graffiti"}
	expect := "\nUsage: go run . [OPTION] [STRING]\n\nEX: go run . --color=<color> <letters to be colored> something"
	output, err := tools.ColorFormat(input)
	if err == nil {
		t.Errorf("\nOutput: %v\nExpected: %v", output, expect)
	}
}
