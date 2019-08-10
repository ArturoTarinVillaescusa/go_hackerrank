package main

import (
	"testing"
)

var testcases = []struct{
	q []int32;
}{
	{[]int32{2, 1, 5, 3, 4}},
	{[]int32{2, 5, 1, 3, 4}},
	{[]int32{5, 1, 2, 3, 7, 8, 6, 4}},
	{[]int32{1, 2, 5, 3, 7, 8, 6, 4}},
	{[]int32{1, 2, 5, 3, 4, 7, 8, 6}},
}

func TestNewYearChaos(t *testing.T){
	t.Parallel()

	for _,c := range testcases {
		out := minimumBribes(c.q)

		if out == -1 {
			t.Errorf("%v was too chaotic", c.q)
		} else {
			t.Logf("The number of bribes for %v was %v", c.q, out)
		}

	}
}

func BenchmarkNewYearChaos(b *testing.B) {
	println(b.N)

	for i :=0; i < b.N; i++ {
		for _, c := range testcases {
			minimumBribes(c.q)
		}
	}
}