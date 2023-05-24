package main

import (
	"ascii/tools"
	"reflect"
	"testing"
)

func Test_0(t *testing.T) {
	input := "file_scan_test0"
	expect := 855
	output := tools.FileScan("../lib/" + input + ".txt")
	if !reflect.DeepEqual(len(output), expect) {
		t.Errorf("\nOutput lenght: %v\nExpected: %v", len(output), expect)
	}
}

func Test_1(t *testing.T) {
	input := "file_scan_test1"
	expect := []string{"A", "B", "C", "D", "E", "F"}
	output := tools.FileScan("../lib/" + input + ".txt")[0:6]
	if !reflect.DeepEqual(output, expect) {
		t.Errorf("\nOutput: %v\nExpected: %v", output, expect)
	}
}
