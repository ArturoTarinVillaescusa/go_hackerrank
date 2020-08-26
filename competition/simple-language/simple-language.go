package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the maximumProgramValue function below.
 */
func maximumProgramValue(n int32) int64 {
	/*
	 * Write your code here.
	 */

	var result int64 = 0

	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create("/tmp/salida")
	// stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	// 3
	// add -1
	// set -3
	// add -2

	for qItr := 0; qItr < int(n); qItr++ {
		lineTemp := strings.Split(readLine(reader), " ")

		op := lineTemp[0]
		valTemp, err := strconv.ParseInt(lineTemp[1], 10, 64)
		checkError(err)
		val := int64(valTemp)

		if op == "add" && val > 0 {
			result += val
		} else {
			result = val
		}
	}

	println(result)
	return result
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	// stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	stdout, err := os.Create("/tmp/salida")
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	n := int32(nTemp)

	result := maximumProgramValue(n)

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
