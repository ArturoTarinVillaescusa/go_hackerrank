package main

import "testing"

var testcases = []struct{
	s string;
	n int64;
	out int64;
}{
	{"aba", 10, 7},
	{"abac", 100, 50},
	{"a", 1000000000000, 1000000000000},
}
func TestRepeatedString(t *testing.T) {
	t.Parallel()

	for _,c := range testcases {
		out := repeatedString(c.s, c.n)

		if out != c.out {
			t.Errorf("Fail!!!: %v was expected for ('%v', %v), but %v was returned instead", c.out, c.s, c.n, out)
		} else {
			t.Logf("Success!!!: %v was expected for ('%v', %v), and %v was returned", c.out, c.s, c.n, out)
		}
	}
}

func BenchmarkRepeatedString(b *testing.B) {
	println(b.N)
	for i :=0; i < b.N; i++ {
		for _, c:= range testcases {
			repeatedString(c.s, c.n)
		}
	}
}