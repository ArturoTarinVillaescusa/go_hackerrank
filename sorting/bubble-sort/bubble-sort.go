package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Complete the countSwaps function below.
func countSwaps(arr []int32) []int {
	var out []int
	var i = 0
	var swaps = 0

	if notSorted(arr) {
		for i=0; i< len(arr) - 1; i++ {
			for j:=i+1; j < len(arr); j++ {
				if (arr[i] > arr[j]) {
					arr[i], arr[j] = arr[j], arr[i]
					swaps++
				}
			}
		}
	}

	// fmt.Println(x)
	fmt.Println("Array is sorted in ", swaps, " swaps.")
	fmt.Println("First Element: ", arr[0])
	fmt.Print("Last Element: ", arr[len(arr) - 1])

	out = []int{swaps, int(arr[0]), int(arr[len(arr) - 1])}

	return out
}

func notSorted(arr []int32) bool {
	for i := 0; i < len(arr) -1; i++ {
		if arr[i] > arr [i + 1] {
			return true
		}
	}

	return false
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

	nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	n := int32(nTemp)

	aTemp := strings.Split(readLine(reader), " ")

	var a []int32

	for i := 0; i < int(n); i++ {
		aItemTemp, err := strconv.ParseInt(aTemp[i], 10, 64)
		checkError(err)
		aItem := int32(aItemTemp)
		a = append(a, aItem)
	}

	countSwaps(a)
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
