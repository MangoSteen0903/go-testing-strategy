package main

import "testing"

func Test_isPrime(t *testing.T) {
	ok, msg := isPrime(0)

	if ok {
		t.Errorf("with %d as parameter, expected false but got true.", 0)
	}

	if msg != "0 is not prime number. prime number should be bigger than 1." {
		t.Error("Message is incorrect.")
	}
}
