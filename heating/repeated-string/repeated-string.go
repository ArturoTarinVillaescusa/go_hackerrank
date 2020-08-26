package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Complete the repeatedString function below.
func repeatedString(s string, n int64) int64 {
	wholeString := n / int64(len(s))
	remainingString := n % int64(len(s))
	var a_counter, remaining_a_counter int64
	for i, r := range s {
		if r == 'a' {
			a_counter++
			if int64(i) < remainingString {
				remaining_a_counter++
			}
		}
	}
	result := wholeString*a_counter + remaining_a_counter
	return result
}

// aba
// 10
func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create("/tmp/salida")
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	s := readLine(reader)

	n, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)

	result := repeatedString(s, n)

	println(result)
	fmt.Fprintf(writer, "%d\n", result)

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
