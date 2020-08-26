package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Complete the angryProfessor function below.
func angryProfessor(k int32, a []int32) string {
	/*
		For example, there are n = 6 students who arrive at times a = [-1 -1 0 0 1 1].
		Four are there on time, and two arrive late.
		If there must be k = 4 for class to go on, in this case the class will continue.
		If there must be k = 5, then class is cancelled.
	*/

	count  := int32(0)
	for i := range a {
		if a[i] <= 0 {
			count++
		}
	}
	fmt.Printf("Verify wether the professor will be angry with the %v threshold and the %v arriving times ...\n", k, a)
	if count >= k {
		return "YES"
	} else {
		return "NO"
	}
}

func main() {
	fmt.Printf("Sample Input:\n" +
		"Number of test cases: 1\n" +
		"Number of students and cancelation threshold: 4 2\n" +
		"4 3-1 -3 4 2\n" +
		"0 -1 2 1\n\n" +
		"Sample Output:\n" +
		"YES\n" +
		"NO\n\n")

	fmt.Printf("Number of test cases: ")
	reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

	os.Setenv("OUTPUT_PATH", "/tmp/salida.txt")
	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024 * 1024)

	tTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	t := int32(tTemp)

	for tItr := 0; tItr < int(t); tItr++ {
		fmt.Printf("Number of students and cancelation threshold (write two numbers, n and k, separated by space): ")

		nk := strings.Split(readLine(reader), " ")

		nTemp, err := strconv.ParseInt(nk[0], 10, 64)
		checkError(err)
		n := int32(nTemp)

		kTemp, err := strconv.ParseInt(nk[1], 10, 64)
		checkError(err)
		k := int32(kTemp)

		fmt.Printf("Arriving times (write %v numbers separated : ", n)
		aTemp := strings.Split(readLine(reader), " ")

		var a []int32

		for i := 0; i < int(n); i++ {
			aItemTemp, err := strconv.ParseInt(aTemp[i], 10, 64)
			checkError(err)
			aItem := int32(aItemTemp)
			a = append(a, aItem)
		}

		result := angryProfessor(k, a)

		fmt.Fprintf(writer, "%s\n", result)
	}

	writer.Flush()
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
