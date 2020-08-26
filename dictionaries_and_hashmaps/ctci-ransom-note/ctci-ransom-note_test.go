package main

import "testing"

var testcases = []struct {
	magazine []string
	note     []string
	out      bool
}{
	{[]string{"ive", "got", "a", "lovely", "bunch", "of", "coconuts"},
		[]string{"ive", "got", "some", "coconuts"},
		false},
	{[]string{"one", "two", "three"},
		[]string{"two"},
		true},
	{[]string{"two", "times", "three", "is", "not", "four"},
		[]string{"two", "times", "two", "is", "four"},
		false},
	{[]string{"give", "me", "one", "grand", "today", "night"},
		[]string{"give", "one", "grand", "today"},
		true},
}

func TestCtciRansomNote(t *testing.T) {
	t.Parallel()

	for _, c := range testcases {
		out := checkMagazine(c.magazine, c.note)

		if out != c.out {
			t.Errorf("Fail!!!: (%v, %v) expected %v, but returned %v", c.magazine, c.note, c.out, out)
		} else {
			t.Logf("Success!!!: (%v, %v) returned %v, exacly what it was expected", c.magazine, c.note, c.out)
		}
	}
}

func BenchmarkCtciRansomNote(b *testing.B) {
	println(b.N)
	for i := 0; i < b.N; i++ {
		for _, c := range testcases {
			checkMagazine(c.magazine, c.note)
		}
	}
}
