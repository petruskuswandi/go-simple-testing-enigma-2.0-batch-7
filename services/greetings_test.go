package services

import "testing"

func TestGreeting_Success(t *testing.T) {
	expectedName := "Ane"
	expected := "Hola!"
	actual := Greeting(expectedName)

	if actual != expected {
		t.Errorf("Expected greeting: %s, but actual got: %s", expected, actual)
	}
}
func TestGreeting_Fail(t *testing.T) {
	expectedName := "Budi"
	expected := "Hola!!"
	actual := Greeting(expectedName)

	if actual == expected {
		t.Errorf("Expected different from greeting: %s", expected)
	}
}

// go test -v -coverprofile=cover.out -> untuk melihat hasil coverage (console)
// go tool cover -html=cover.out -> untuk melihat hasil coverage (html/web)
// go test ./... -v => untuk mencari keseluruhan baik itu folder atau sub folder yang memiliki file_test.go
