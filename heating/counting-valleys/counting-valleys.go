package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Complete the countingValleys function below.
func countingValleys(n int32, s string) int32 {
	result := 0;
	level := 0;
	inSeaLevel := false
	for _, c := range s {
		switch c {
		case 'D':
			level--
		case 'U':
			level++
		}
		if inSeaLevel && level == 0 {
			result++
			inSeaLevel = false
		} else if !inSeaLevel && level < 0 {
			inSeaLevel = true
		}
	}
	return int32(result)
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

	stdout, err := os.Create("/tmp/salida")
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024 * 1024)

	nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	n := int32(nTemp)

	s := readLine(reader)

	result := countingValleys(n, s)

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