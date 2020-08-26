package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

// Complete the sherlockAndAnagrams function below.
func sherlockAndAnagrams(s string) int32 {
	data := []rune(s)
	size := len(data)
	var res int32

	for i := 1; i < size; i++ {
		m := make(map[string]int32)

		// Split the string in substrings
		for j := 0; j < size-i+1; j++ {

			// Get the new substring
			subString := []rune(s[j : j+i])

			//order the substring
			sort.Slice(subString, func(i, j int) bool {
				return subString[i] < subString[j]
			})

			//add the substring in the map and update the value
			m[string(subString)]++

		}
		//read the map
		for _, v := range m {
			//get the number of occurency of the substring
			if v > 1 {
				// Get Triangular number (wikipedia)
				res += (v - 1) * v / 2
			}
		}

	}

	return res
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	// stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	stdout, err := os.Create("/tmp/salida")
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	qTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	q := int32(qTemp)

	for qItr := 0; qItr < int(q); qItr++ {
		s := readLine(reader)

		/*
			2
			abba
			abcd

			out:
			4
			0

			2
			ifailuhkqq
			kkkk

			out:
			3
			10

			1
			cdcd
			out:
			5
		*/
		result := sherlockAndAnagrams(s)

		fmt.Fprintf(writer, "%d\n", result)
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
