package main

import "testing"

var testcases = []struct {
	in []int32;
	out int32;
}{
	{[]int32{0, 0, 1, 0, 0, 1, 0}, 4},
	{[]int32{0, 0, 1, 0, 1, 1, 1, 0, 0, 1, 0}, 5},
}

func TestJumpingOnClouds(t *testing.T){
	t.Parallel()
	for _, c:= range testcases {
		out := jumpingOnClouds(c.in)

		if (c.out == out) {
			t.Logf("Success!!!: %v was returned and %v was expected for %v", out, c.out, c.in)
		} else {
			t.Errorf("Fail!!!: %v was returned, but %v was expected for %v", out, c.out, c.in)
		}
	}
}

func BenchmarkJumpingOnClouds(b *testing.B) {
	println(b.N)
	for i := 0; i < b.N; i++ {
		for _, c := range testcases {
			jumpingOnClouds(c.in)
		}
	}
}
