package utils_test

import (
	"assignment-2/utils"
	"testing"
)

func TestToJsonMap(t *testing.T) {
	type testStruct struct {
		A float64 `json:"a"`
		B string  `json:"b"`
		C bool    `json:"c"`
	}

	testData := testStruct{
		A: 1,
		B: "test",
		C: true,
	}

	jsonMap, err := utils.ToJsonMap(testData)
	if err != nil {
		t.Errorf("Error converting struct to map: %v", err)
	}

	if jsonMap["a"] != testData.A {
		t.Errorf("Expected A to be %v, got %v", testData.A, jsonMap["a"])
	}

	if jsonMap["b"] != testData.B {
		t.Errorf("Expected B to be %v, got %v", testData.B, jsonMap["b"])
	}

	if jsonMap["c"] != testData.C {
		t.Errorf("Expected C to be %v, got %v", testData.C, jsonMap["c"])
	}
}
