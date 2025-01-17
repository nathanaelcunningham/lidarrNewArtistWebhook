package main

import "testing"

func TestMain_sanitizeFolderName(t *testing.T) {
	artistName := "J. Cole"

	expected := "J Cole"

	actual := sanitizeFolderName(artistName)

	if actual != expected {
		t.Errorf("Expected %s, but got %s", expected, actual)
	}
}
