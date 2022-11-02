package main

import "testing"

func TestPlus1(t *testing.T) {
	r := 1 + 2
	expect := 3
	if expect != r {
		t.Errorf("expect is %d but actual is %d", expect, r)
	}
}

func TestPlus2(t *testing.T) {
	r := 1 + 2 + 1
	expect := 4
	if r != expect {
		t.Errorf("expect is %d but actual is %d", expect, r)
	}
}

func TestPlus3(t *testing.T) {
	r := 1 + 2 + 1 + 2
	expect := 6
	if r != expect {
		t.Errorf("expect is %d but actual is %d", expect, r)
	}
}
