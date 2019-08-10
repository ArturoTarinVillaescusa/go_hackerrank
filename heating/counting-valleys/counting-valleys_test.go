package main

import (
	"testing"
)

var testcases = []struct{
	in string;
	out int32;
}{
	{ "U D D U U U D U U U U", 1},
	{"U D D D U D U U U U U U U U D D D D D D D D D D U D D D U D U U U U U U U U D D D D D D D D D D U D D D U D U U U U U U U U D D D D D D D D D D",
	 2},
}
func TestCountingValleys(t *testing.T){
	t.Parallel()

	for _, c := range testcases {
		out := countingValleys(int32(len(c.in)), c.in)
		if out == c.out {
			t.Logf("Success!!!: expected %v and received %v for %v", c.out, out, c.in)
		} else {
			t.Errorf("Error!!!: expected %v and received %v for %v", c.out, out, c.in)
		}
	}
}

func BenchmarkCountingValleys(b *testing.B) {
	println(b.N)
	for i := 0; i < b.N; i++ {
		for _, c := range testcases {
			countingValleys(int32(len(c.in)), c.in)
		}
	}
}