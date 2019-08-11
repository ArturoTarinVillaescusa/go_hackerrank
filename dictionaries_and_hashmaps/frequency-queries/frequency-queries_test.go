package main

import "testing"

var testcases = []struct{
	queries [][]int32;
	out []int32;
}{
	{[][]int32{{3, 4}, {2, 1003}, {1, 16}, {3, 1}}, []int32{0, 1}},
	{[][]int32{{1, 3}, {2, 3}, {3, 2}, {1, 4}, {1, 5}, {1, 5}, {1, 4}, {3, 2}, {2, 4}, {3, 2}}, []int32{0, 1, 1}},
	{[][]int32{{1, 5}, {1, 6}, {3, 2}, {1, 10}, {1, 10}, {1, 6}, {2, 5}, {3, 2}}, []int32{0, 1}},
}

func TestFrequencyQueries(t *testing.T) {
	t.Parallel()

	for _, c := range testcases {
		out := freqQuery(c.queries)

		if c.out[0] != out[0] {
			t.Errorf("Fail!!!: %v expected for %v but %v returned", c.out, c.queries, out)
		} else {
			t.Logf("Success!!!: %v expected for %v and %v returned", c.out, c.queries, out)
		}
/*

		if out != c.out {
			t.Errorf("Fail!!!: %v expected for %v but %v returned", c.out, c.s, out)
		} else {
			t.Logf("Success!!!: %v expected for %v and %v returned", c.out, c.s, out)
		}
*/
	}
}


func BenchmarkFrequencyQueries(b *testing.B) {
	println(b.N)
	for i := 0; i < b.N; i++ {
		for _, c := range testcases {
			freqQuery(c.queries)
		}
	}
}

