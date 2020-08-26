package main

import "testing"

var testcases = []struct {
	prices []int32
	k      int32
	out    int32
}{
	{[]int32{1, 12, 5, 111, 200, 1000, 10}, 50, 4},
	{[]int32{1, 2, 3, 4}, 7, 3},
	{[]int32{3, 7, 2, 9, 4}, 15, 3},
}

func TestMarkAndToys(t *testing.T) {
	t.Parallel()

	for _, c := range testcases {
		out := maximumToys(c.prices, c.k)
		t.Logf("%v", out)

		if c.out != out {
			t.Errorf("Fail!!!: %v expected for (%v, %v) but %v returned", c.out, c.k, c.prices, out)
		} else {
			t.Logf("Success!!!: %v expected for (%v, %v) and %v returned", c.out, c.k, c.prices, out)
		}

	}
}
