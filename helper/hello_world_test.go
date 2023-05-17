package helper

import "testing"

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("fuad")

	if result != "Hello fuad" {
		panic("result not hello fuad")
	}
}
