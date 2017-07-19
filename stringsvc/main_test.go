package main

import "testing"

var sut = stringService{}

func TestUppercase(t *testing.T) {
	ssUpper, err := sut.Uppercase(nil, "foo")
	if err != nil {
		t.Error("Unable to uppercase string")
	}
	if ssUpper != "FOO" {
		t.Errorf("This isn't uppercase foo: %s", ssUpper)
	}
}

func TestCount(t *testing.T) {
	count := sut.Count(nil, "foo")
	if count != 3 {
		t.Errorf("Count is %d", count)
	}
}

func Test
