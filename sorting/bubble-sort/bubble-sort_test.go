package main

import "testing"

var testcases = []struct{
	arr []int32;
	out []int;
} {
	{[]int32{3, 2, 1}, []int{3, 1, 3}},
	{[]int32{1, 2, 3}, []int{0, 1, 3}},
}

func TestBubbleSort(t *testing.T) {
	t.Parallel()

	for _, c := range testcases {
		out := countSwaps(c.arr)
		t.Logf("%v", out)

		if c.out[0] != out[0] || c.out[1] != out[1] || c.out[2] != out[2]   {
			t.Errorf("Fail!!!: %v expected for %v but %v returned", c.out, c.arr, out)
		} else {
			t.Logf("Success!!!: %v expected for %v and %v returned", c.out, c.arr, out)
		}

	}
}
