package sock_merchant

import (
	"testing"
)

var testcases = []struct {
	in  []int32
	out int32
}{
	{[]int32{10, 20, 20, 10, 10, 30, 50, 10, 20}, 3},
	{[]int32{10, 20, 20, 10, 7, 7, 7, 27, 10, 30, 50, 10, 20}, 4},
	{[]int32{1, 20, 2, 1, 2, 90, 3, 3, 10, 20}, 4},
	{[]int32{1, 20, 2, 1, 2, 3, 3, 3, 3, 3, 3, 3, 3, 10, 20}, 7},
}

func TestSockMerchant(t *testing.T) {
	t.Parallel()
	for _, c := range testcases {
		out := sockMerchant(int32(len(c.in)), c.in)

		if out != c.out {
			t.Errorf("Fail!!!:  %v was expected for %v, but %v returned", c.out, c.in, out)
		} else {
			t.Logf("Success!!!: %v was expected for %v, and %v returned", c.out, c.in, out)

		}
	}
}

func BenchmarkSockMerchant(b *testing.B) {
	println(b.N)
	for i := 0; i < b.N; i++ {
		for _, c := range testcases {
			sockMerchant(int32(len(c.in)), c.in)
		}
	}
}
