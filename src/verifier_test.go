package main

import (
	"testing"
)

func TestReadFromFile(t *testing.T) {
	_, err := readFromFile("../test_cases/test_false.json")
	if err != nil {
		t.Fatal("Error reading from file:", err)
	}
}

func TestValidateJsonStr(t *testing.T) {
	testWrongStr, _ := readFromFile("../test_cases/test_wrong.json")
	_, err := validateJsonStr(testWrongStr)
	if err == nil {
		t.Errorf("Expected invalid json to generate error")
	}

	testFalseStr, _ := readFromFile("../test_cases/test_false.json")
	testFalse, _ := validateJsonStr(testFalseStr)

	if testFalse {
		t.Errorf("Expected: false, got: true")
	}

	testTrueStr, _ := readFromFile("../test_cases/test_true.json")
	testTrue, _ := validateJsonStr(testTrueStr)

	if !testTrue {
		t.Errorf("Expected: true, got: false")
	}

}
