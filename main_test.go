package main

import (
	"testing"
)

func TestReverseCaptcha(t *testing.T) {
	o1 := ReverseCaptcha("1122")
	if o1 != 3 {
		t.Errorf("Expected output of 3 but got %d", o1)
	}
	o2 := ReverseCaptcha("1111")
	if o2 != 4 {
		t.Errorf("Expected output of 4 but got %d", o2)
	}
	o3 := ReverseCaptcha("1234")
	if o3 != 0 {
		t.Errorf("Expected output of 0 but got %d", o3)
	}
	o4 := ReverseCaptcha("91212129")
	if o4 != 9 {
		t.Errorf("Expected output of 9 but got %d", o4)
	}
}
