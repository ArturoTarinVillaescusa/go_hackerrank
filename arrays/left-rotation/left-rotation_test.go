package main

import (
	"testing"
)

var testcases = []struct {
	a   []int32
	d   int32
	out []int32
}{
	{[]int32{1, 2, 3, 4, 5}, 4, []int32{5, 1, 2, 3, 4}},
}

func Equal(a, b []int32) bool {
	if len(a) != len(b) {
		return false
	}

	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
func TestLeftRotation(t *testing.T) {
	t.Parallel()

	for _, c := range testcases {
		out := rotLeft(c.a, c.d)

		if Equal(out, c.out) {
			t.Logf("Success!!!: %v was expected for (%v, %v), and as expected, %v was returned", c.out, c.a, c.d, out)
		} else {
			t.Errorf("Fail!!!: %v was expected for (%v, %v), but %v was returned", c.out, c.a, c.d, out)
		}

	}
}

func BenchmarkLeftRotation(b *testing.B) {
	println(b.N)

	for i := 0; i < b.N; i++ {

		for _, c := range testcases {
			rotLeft(c.a, c.d)
		}
	}

}
