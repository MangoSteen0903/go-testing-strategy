package main

import "testing"

type primeTest struct {
	name        string
	testNum     int
	expected    bool
	expectedMsg string
}

func Test_isPrime(t *testing.T) {
	tests := []primeTest{
		{
			"prime", 7, true, "7 is a prime number.",
		},
		{
			"not prime", 8, false, "8 is not a prime number.",
		},
		{
			"zero", 0, false, "0 is not prime number. prime number should be bigger than 1.",
		},
		{
			"negative", -1, false, "-1 : should not be negative number. please try again.",
		},
	}

	for _, test := range tests {
		result, msg := isPrime(test.testNum)

		if test.expected && !result {
			t.Errorf("%s: expected true but got false.", test.name)
		}

		if !test.expected && result {
			t.Errorf("%s: expected false but got true.", test.name)
		}

		if test.expectedMsg != msg {
			t.Errorf("%s: expected %s but got %s", test.name, test.expectedMsg, msg)
		}
	}
}
