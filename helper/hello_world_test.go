package helper

import "testing"

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("fuad")

	if result != "Hello fuad" {
		t.Fatal("result must be hello fuad")
	}
}
