package main

import "testing"

func TestHello(t *testing.T) {
	var result = helloTest()
	var expect = "Hello World!"

	if result != expect {
		t.Errorf("expect %v and got %v", result, expect)
	}
}
