package main

import "testing"

var testcases = []struct{
	arr []int64;
	r int64;
	out int64;
}{
	{[]int64{1, 4, 16, 64}, 4, 2},
	{[]int64{1, 2, 2, 4}, 2, 2},
	{[]int64{1, 3, 9, 9, 27, 81}, 3, 6},
	{[]int64{1, 5, 5, 25, 125}, 5, 4},
}

func TestSherlockAndAnagrams(t *testing.T) {
	t.Parallel()

	for _, c := range testcases {
		out := countTriplets(c.arr, c.r)

		if out != c.out {
			t.Errorf("Fail!!!: %v expected for %v but %v returned", c.out, c.arr, out)
		} else {
			t.Logf("Success!!!: %v expected for %v and %v returned", c.out, c.arr, out)
		}
	}
}


func BenchmarkSherlockAndAnagrams(b *testing.B) {
	println(b.N)
	for i := 0; i < b.N; i++ {
		for _, c := range testcases {
			countTriplets(c.arr, c.r)
		}
	}
}


