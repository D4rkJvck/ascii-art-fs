package main

import (
	"ascii/tools"
	"reflect"
	"testing"
)

func Test_Black(t *testing.T) {
	input := "black"
	expect := 30
	output := tools.ColorChoiceAnsi(input)
	if !reflect.DeepEqual(output, expect) {
		t.Errorf("\nOutput: %v\nExpected: %v", output, expect)
	}
}

func Test_Red(t *testing.T) {
	input := "red"
	expect := 31
	output := tools.ColorChoiceAnsi(input)
	if !reflect.DeepEqual(output, expect) {
		t.Errorf("\nOutput: %v\nExpected: %v", output, expect)
	}
}

func Test_Green(t *testing.T) {
	input := "green"
	expect := 32
	output := tools.ColorChoiceAnsi(input)
	if !reflect.DeepEqual(output, expect) {
		t.Errorf("\nOutput: %v\nExpected: %v", output, expect)
	}
}

func Test_Yellow(t *testing.T) {
	input := "yellow"
	expect := 33
	output := tools.ColorChoiceAnsi(input)
	if !reflect.DeepEqual(output, expect) {
		t.Errorf("\nOutput: %v\nExpected: %v", output, expect)
	}
}

func Test_Blue(t *testing.T) {
	input := "blue"
	expect := 34
	output := tools.ColorChoiceAnsi(input)
	if !reflect.DeepEqual(output, expect) {
		t.Errorf("\nOutput: %v\nExpected: %v", output, expect)
	}
}

func Test_Magenta(t *testing.T) {
	input := "magenta"
	expect := 35
	output := tools.ColorChoiceAnsi(input)
	if !reflect.DeepEqual(output, expect) {
		t.Errorf("\nOutput: %v\nExpected: %v", output, expect)
	}
}

func Test_Cyan(t *testing.T) {
	input := "cyan"
	expect := 36
	output := tools.ColorChoiceAnsi(input)
	if !reflect.DeepEqual(output, expect) {
		t.Errorf("\nOutput: %v\nExpected: %v", output, expect)
	}
}

func Test_White(t *testing.T) {
	input := "white"
	expect := 37
	output := tools.ColorChoiceAnsi(input)
	if !reflect.DeepEqual(output, expect) {
		t.Errorf("\nOutput: %v\nExpected: %v", output, expect)
	}
}

func Test_Gray(t *testing.T) {
	input := "gray"
	expect := 90
	output := tools.ColorChoiceAnsi(input)
	if !reflect.DeepEqual(output, expect) {
		t.Errorf("\nOutput: %v\nExpected: %v", output, expect)
	}
}

func Test_LightGreen(t *testing.T) {
	input := "lightgreen"
	expect := 92
	output := tools.ColorChoiceAnsi(input)
	if !reflect.DeepEqual(output, expect) {
		t.Errorf("\nOutput: %v\nExpected: %v", output, expect)
	}
}

func Test_LightBlue(t *testing.T) {
	input := "lightblue"
	expect := 94
	output := tools.ColorChoiceAnsi(input)
	if !reflect.DeepEqual(output, expect) {
		t.Errorf("\nOutput: %v\nExpected: %v", output, expect)
	}
}

func Test_Pink(t *testing.T) {
	input := "pink"
	expect := 95
	output := tools.ColorChoiceAnsi(input)
	if !reflect.DeepEqual(output, expect) {
		t.Errorf("\nOutput: %v\nExpected: %v", output, expect)
	}
}

func Test_Empty(t *testing.T) {
	input := ""
	expect := 0
	output := tools.ColorChoiceAnsi(input)
	if !reflect.DeepEqual(output, expect) {
		t.Errorf("\nOutput: %v\nExpected: %v", output, expect)
	}
}
