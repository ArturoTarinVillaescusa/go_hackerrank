package main

import "testing"

var testcases = []struct{
	s string;
	out int32;
}{
	{"cdcd", 5},
	{"ifailuhkqq", 3},
	{"kkkk", 10},
	{"abba", 4},
	{"abcd", 0},
}

func TestSherlockAndAnagrams(t *testing.T) {
	t.Parallel()

	for _, c := range testcases {
		out := sherlockAndAnagrams(c.s)

		if out != c.out {
			t.Errorf("Fail!!!: %v expected for %v but %v returned", c.out, c.s, out)
		} else {
			t.Logf("Success!!!: %v expected for %v and %v returned", c.out, c.s, out)
		}
	}
}


func BenchmarkSherlockAndAnagrams(b *testing.B) {
	println(b.N)
	for i := 0; i < b.N; i++ {
		for _, c := range testcases {
			sherlockAndAnagrams(c.s)
		}
	}
}

