package main

import (
	"ascii/tools"
	"reflect"
	"testing"
)

func Test_StringFormat(t *testing.T) {
	input := "\\nHello\\n\\nThere\\n"
	expect := []string{"", "Hello", "", "There", ""}
	output := tools.StringFormat(input)
	if !reflect.DeepEqual(output, expect) {
		t.Errorf("\nOutput: %v\nExpected: %v", output, expect)
	}
}
