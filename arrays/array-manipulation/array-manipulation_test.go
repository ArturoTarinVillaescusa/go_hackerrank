package main

import "testing"

var testcases = []struct{
	n int32;
	queries [][]int32;
	out int64;
}{
	{5, [][]int32{
		{1, 2, 100},
		{2, 5, 100},
		{3, 4, 100},
	}, 200},
	{3, [][]int32{
		{1, 5, 3},
		{4, 8, 7},
		{6, 9, 1},
	}, 10},
	{3, [][]int32{
		{2, 6, 8},
		{3, 5, 7},
		{1, 8, 1},
		{5, 9, 15},
	}, 31},
	{3, [][]int32{
		{2250, 2540, 180674},
		{2459, 3892, 434122},
		{2321, 3289, 363503},
		{1975, 2754, 374161},
		{3283, 3759, 45954},
		{711,3596, 978769},
		{1468, 3899, 109177},
		{1044, 2661, 538264},
		{1182, 1224, 795164},
		{3664, 3745, 221217},
		{1994, 2209, 471756},
		{2237, 3236, 358214},
		{1526, 1710, 220886},
		{631, 1214, 918861},
		{265, 3727, 734619},
		{975, 1681, 154566},
		{802, 2442, 530465},
		{1454, 1980, 443860},
		{596, 2677, 745394},
	}, 5396102},


}

func TestArrayManipulation(t *testing.T) {
	t.Parallel()

	for _, c := range testcases {
		out := arrayManipulation(c.n, c.queries)

		if (out != c.out) {
			t.Errorf("Fail!!!: %v was expected for %v, but we obtained %v", c.out, c.queries, out)
		} else {
			t.Logf("Success!!!: The function returned %v for %v", c.out, c.queries)
		}
	}
}

func BenchmarkArrayManipulation(b *testing.B) {
	println(b.N)

	for i := 0; i < b.N; i++ {
		for _, c := range testcases {
			arrayManipulation(c.n, c.queries)
		}
	}
}
