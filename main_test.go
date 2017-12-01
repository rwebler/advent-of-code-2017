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

func TestReverseCaptchaHalfway(t *testing.T) {
	/*
		1212 produces 6: the list contains 4 items, and all four digits match the digit 2 items ahead.
		1221 produces 0, because every comparison is between a 1 and a 2.
		123425 produces 4, because both 2s match each other, but no other digit has a match.
		123123 produces 12.
		12131415 produces 4.
	*/
	o1 := ReverseCaptchaHalfway("1212")
	if o1 != 6 {
		t.Errorf("Expected output of 6 but got %d", o1)
	}
	o2 := ReverseCaptchaHalfway("1221")
	if o2 != 0 {
		t.Errorf("Expected output of 0 but got %d", o2)
	}
	o3 := ReverseCaptchaHalfway("123425")
	if o3 != 4 {
		t.Errorf("Expected output of 4 but got %d", o3)
	}
	o4 := ReverseCaptchaHalfway("123123")
	if o4 != 12 {
		t.Errorf("Expected output of 12 but got %d", o4)
	}
	o5 := ReverseCaptchaHalfway("12131415")
	if o5 != 4 {
		t.Errorf("Expected output of 4 but got %d", o5)
	}
}
