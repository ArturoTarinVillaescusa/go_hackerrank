package main

import "testing"

var testcases = []struct{
	s1 string;
	s2 string;
	out string;
} {
	{"abba", "a", "YES"},
	{"abba", "c", "NO"},
}

func TestTwoStrings(t *testing.T)  {
	t.Parallel()

	for _, c := range testcases {
		out := twoStrings(c.s1, c.s2)

		if out != c.out {
			t.Errorf("Fail!!!: '%v' and '%v' doesn't contain common substrings, but your function told TRUE", c.s1, c.s2)
		} else {
			t.Logf("Success!!!: '%v' and '%v'  contain common substrings, and your function found it!", c.s1, c.s2)
		}
	}
}

func BenchmarkTwoStrings(b *testing.B) {
	println(b.N)
	for i := 0; i < b.N; i++ {
		for _, c := range testcases {
			twoStrings(c.s1, c.s2)
		}
	}
}