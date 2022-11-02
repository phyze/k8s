package main

import "testing"

func TestSub1(t *testing.T) {
	r := 8 - 1
	expect := 7
	if expect != r {
		t.Errorf("expect is %d but actual is %d", expect, r)
	}
}

func TestSub2(t *testing.T) {
	r := 8 - 2
	expect := 6
	if r != expect {
		t.Errorf("expect is %d but actual is %d", expect, r)
	}
}

func TestSub3(t *testing.T) {
	r := 8 - 3
	expect := 5
	if r != expect {
		t.Errorf("expect is %d but actual is %d", expect, r)
	}
}
