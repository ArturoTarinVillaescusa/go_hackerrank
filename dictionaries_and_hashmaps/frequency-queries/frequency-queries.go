package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Complete the freqQuery function below.

const (
	add = iota + 1
	delete
	check
)

func freqQuery(queries [][]int32) []int32 {
	frequency := make(map[int32]int32)

	db := make(map[int32]int32)
	res := make([]int32, 0, 100)

	for _, v := range queries {
		operation, data := v[0], v[1]
		switch operation {
		case add:
			frequency[db[data]]--
			db[data]++
			frequency[db[data]]++
		case delete:
			if db[data] > 0 {
				frequency[db[data]]--
				db[data]--
				frequency[db[data]]++
			}
		case check:
			if frequency[data] > 0 {
				res = append(res, 1)
				break
			}
			res = append(res, 0)
		}
	}

	fmt.Println(res)
	return res
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	// stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	stdout, err := os.Create("/tmp/salida")
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	qTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	q := int32(qTemp)

	var queries [][]int32
	for i := 0; i < int(q); i++ {
		queriesRowTemp := strings.Split(strings.TrimRight(readLine(reader), " \t\r\n"), " ")

		var queriesRow []int32
		for _, queriesRowItem := range queriesRowTemp {
			queriesItemTemp, err := strconv.ParseInt(queriesRowItem, 10, 64)
			checkError(err)
			queriesItem := int32(queriesItemTemp)
			queriesRow = append(queriesRow, queriesItem)
		}

		if len(queriesRow) != 2 {
			panic("Bad input")
		}

		queries = append(queries, queriesRow)
	}

	ans := freqQuery(queries)

	for i, ansItem := range ans {
		fmt.Fprintf(writer, "%d", ansItem)

		if i != len(ans)-1 {
			fmt.Fprintf(writer, "\n")
		}
	}

	fmt.Fprintf(writer, "\n")

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
