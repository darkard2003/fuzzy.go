package fuzzygo

import "testing"

func TestHello(t *testing.T) {
	if SayHello() != "Hello, world." {
		t.Error("Expected Hello, world.")
	}
}
