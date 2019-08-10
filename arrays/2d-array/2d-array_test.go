package main

import "testing"

var testcases = []struct{
	arr [][]int32;
	out int32;
} {
	{
		[][]int32{
			{-1, -1,  0, -9, -2, -2},
			{-2, -1, -6, -8, -2, -5},
			{-1, -1, -1, -2, -3, -4},
			{-1, -9, -2, -4, -4, -5},
			{-7, -3, -3, -2, -9, -9},
			{-1, -3, -1, -2, -4, -5},
		},
		-6},
	{
		[][]int32{
			{-9, -9, -9,  1, 1, 1},
			{0, -9,  0,  4, 3, 2},
			{-9, -9, -9,  1, 2, 3},
			{0,  0,  8,  6, 6, 0},
			{0,  0, 0, -2, 0, 0},
			{0,  0,  1,  2, 4, 0},
		},
		28},
	{[][]int32{
		{1,  2,   3,  4,  5,  6},
		{7,  8,   9, 10, 11, 12},
		{13, 14, 15, 16, 17, 18},
		{19, 20, 21, 22, 23, 24},
		{25, 26, 27, 28, 29, 30},
		{31, 32, 33, 34, 35, 36},
	},
	203},
	{[][]int32{
		{1, 1, 1, 0, 0, 0},
		{0, 1, 0, 0, 0, 0},
		{1, 1, 1, 0, 0, 0},
		{0, 0, 2, 4, 4, 0},
		{0, 0, 0, 2, 0, 0},
		{0, 0, 1, 2, 4, 0},
		},
	 19},
	{[][]int32{
		{1, 1, 1, 0, 0, 0},
		{0, 1, 0, 0, 0, 0},
		{1, 1, 1, 0, 0, 0},
		{0, 0, 2, 4, 4, 0},
		{0, 0, 0, 2, 0, 0},
		{0, 0, 1, 2, 4, 0},
	},
		7},
}

func Test2dArray(t *testing.T) {
	t.Parallel()
	for _, c:= range testcases {
		out := hourglassSum(c.arr)
		if out != c.out {
			t.Errorf("Fail!!!: %v was returned for (%v), but %v was expected", out, c.arr, c.out)
		} else {
			t.Logf("Success!!!: %v was returned for (%v) and %v was expected", out, c.arr, c.out)
		}
	}
}

func Benchmark2dArray(b *testing.B) {
	println(b.N)
	for i := 0; i < b.N; i++ {
		for _, c:= range testcases {
			hourglassSum(c.arr)
		}
	}
}
