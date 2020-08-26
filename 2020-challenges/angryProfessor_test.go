package main

import "testing"

var testcases = []struct {
	k int32
	a []int32
	out string
}{

	{3,
	[]int32{-1, -3, 4, 2},
	"YES"},
	{2,
	[]int32{0, -1, 2, 1},
	"NO",
	},
	{2,
	[]int32{-93, -86, 49, -62, -90, -63, 40, 72, 11, 67},
	"NO"},
	{4,
		[]int32{-93, -86, 49, -62, -90, -63, 40, 72, 11, 67},
		"NO"},
	{10,
		[]int32{23, -35, -2, 58, -67, -56, -42, -73, -19, 37},
		"YES"},
	{9,
		[]int32{13, 91, 56, -62, 96, -5, -84, -36, -46, -13},
		"YES"},
	{8,
		[]int32{45, 67, 64, -50, -8, 78, 84, -51, 99, 60},
		"YES"},
	{7,
		[]int32{26, 94, -95, 34, 67, -97, 17, 52, 1, 86},
		"YES"},
	{2,
		[]int32{19, 29, -17, -71, -75, -27, -56, -53, 65, 83},
		"NO"},
	{10,
		[]int32{-32, -3, -70, 8, -40, -96, -18, -46, -21, -79},
		"YES"},
	{9,
		[]int32{-50, 0, 64, 14, -56, -91, -65, -36, 51, -28},
		"YES"},
	{6,
		[]int32{-58, -29, -35, -18, 43, -28, -76, 43, -13, 6},
		"NO"},
	{1,
		[]int32{88, -17, -96, 43, 83, 99, 25, 90, -39, 86},
		"NO"},
}

func Test(t *testing.T)  {
	t.Parallel()
	for _, c := range testcases {
		out := angryProfessor(c.k, c.a)

		if out != c.out {
			t.Errorf("Fail!!!. %v was expected for k=%v and a=%v, but %v was returned.", c.out, c.k, c.a, c.out)
		} else {
			t.Logf("Success!!!. %v was expected for k=%v and a=%v, and %v was returned.", c.out, c.k, c.a, c.out)
		}
	}
}

func BenchmarkAngryProfessor(b *testing.B) {
	println(b.N)
	for i := 0; i < b.N; i++ {
		for _, c := range testcases {
			angryProfessor(c.k, c.a)
		}
	}
}
