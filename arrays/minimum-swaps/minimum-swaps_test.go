package main

import "testing"

var testcases = []struct{
	arr []int32;
	out int32;
}{
	{[]int32{4, 3, 1, 2}, 3},
	{[]int32{1, 3, 5, 2, 4, 6, 7}, 3},
	{[]int32{1, 3, 5, 2, 4, 6, 7}, 7},
}

func TestMinimumSwaps(t *testing.T) {
	t.Parallel()

	for _, c:= range testcases {
		out := minimumSwaps(c.arr)

		if out != c.out {
			t.Errorf("Fail!!!: The expected swaps for %v was %v, but we got %v", c.arr, c.out, out)
		} else {
			t.Logf("Success!!!: We've got the expected swaps for %v: %v", c.arr, out)
		}
	}
}

func BenchmarkMinimumSwaps(b *testing.B) {
	println(b.N)

	for i := 0; i < b.N; i++ {
		for _, c := range testcases {
			minimumSwaps(c.arr)
		}
	}
}