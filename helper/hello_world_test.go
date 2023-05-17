package helper

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHelloWorldAssert(t *testing.T) {
	result := HelloWorld("fuad")
	assert.Equal(t, "Hello fuad", result, "result must be 'hello fuad'")
}

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("fuad")

	if result != "Hello fuad" {
		t.Fatal("result must be hello fuad")
	}
}
