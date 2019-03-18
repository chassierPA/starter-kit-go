package loader

import (
	"testing"
)

//TestJSON struct to unit test
type TestJSON struct {
	Test string `json:"test"`
}

func TestLoaderValidFile(t *testing.T) {
	pathFile := "./unit_test.json"
	var test TestJSON

	err := LoadJSON(pathFile, &test)
	if err != nil {
		t.Errorf("ValidFile return error: %s", err)
	}
}

func TestLoaderInvalidFile(t *testing.T) {
	pathFile := "invalid.json"
	var test TestJSON

	err := LoadJSON(pathFile, &test)
	if err == nil {
		t.Errorf("invalid path file doesn't return an error !")
	}
}
