package helloworld

import "testing"

func Test_add(t *testing.T) {
	r := add(1, 1)
	if r != 2 {
		t.Error("error")
	}
}
