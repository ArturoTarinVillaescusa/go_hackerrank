package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Complete the jumpingOnClouds function below.
func jumpingOnClouds(c []int32) int32 {
	var result int
	steps := -1
	for _, x := range c {
		if x == 0 {
			steps++
		} else {
			result += (steps / 2) + (steps % 2) + 1
			steps = -1
		}
	}
	result += (steps / 2) + (steps % 2)
	return int32(result)
}

// 7
// 0 0 1 0 0 1 0
func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

	stdout, err := os.Create("/tmp/salida")
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024 * 1024)

	nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	n := int32(nTemp)

	cTemp := strings.Split(readLine(reader), " ")

	var c []int32

	for i := 0; i < int(n); i++ {
		cItemTemp, err := strconv.ParseInt(cTemp[i], 10, 64)
		checkError(err)
		cItem := int32(cItemTemp)
		c = append(c, cItem)
	}

	result := jumpingOnClouds(c)

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

