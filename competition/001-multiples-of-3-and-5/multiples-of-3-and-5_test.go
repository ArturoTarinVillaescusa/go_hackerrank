package main

import "testing"

var testcases = []struct{
	n int;
	out int;
} {
	{10, 23},
	{100, 2318},
}

func TestMultiplesOf3And5(t *testing.T)  {
	t.Parallel()

	for _, c := range testcases {
		out := multiplesOf3And5(c.n)

		if out != c.out {
			t.Errorf("Fail!!!: %v was expected for %v, but you returned %v instead!!!", c.out, c.n, out)
		} else {
			t.Logf("Success!!!: you returned %v, exactly what it was expected for %v", out, c.n)
		}
	}
}

func BenchmarkMultiplesOf3And5(b *testing.B) {
	println(b.N)
	for i := 0; i < b.N; i++ {
		for _, c := range testcases {
			multiplesOf3And5(c.n)
		}
	}
}