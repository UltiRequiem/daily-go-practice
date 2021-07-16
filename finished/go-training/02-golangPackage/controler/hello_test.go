package controller

import "testing"

func TestHelloWorld(t *testing.T) {
	hello := HelloWorld("Zero")
	if hello != "Hi, Zero" {
		t.Errorf("Testing fail")
	}

	hello = HelloWorld("Zero ")
	if hello != "Hi, Zero" {
		t.Errorf("Testing fail")
	}
}
