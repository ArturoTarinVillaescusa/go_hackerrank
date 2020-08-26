package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Complete the hourglassSum function below.
func hourglassSum(arr [][]int32) int32 {
	var maxVal int32 = -100000000
	for x := 0; x < len(arr)-2; x++ {
		for y := 0; y < len(arr[0])-2; y++ {
			var sum int32 = 0
			sum += arr[x][y]
			sum += arr[x][y+1]
			sum += arr[x+2][y]
			sum += arr[x+1][y+1]
			sum += arr[x][y+2]
			sum += arr[x+2][y+1]
			sum += arr[x+2][y+2]
			if sum > maxVal {
				maxVal = sum
			}
		}
	}
	return maxVal
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create("/tmp/salida")
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	var arr [][]int32
	for i := 0; i < 6; i++ {
		arrRowTemp := strings.Split(readLine(reader), " ")

		var arrRow []int32
		for _, arrRowItem := range arrRowTemp {
			arrItemTemp, err := strconv.ParseInt(arrRowItem, 10, 64)
			checkError(err)
			arrItem := int32(arrItemTemp)
			arrRow = append(arrRow, arrItem)
		}

		if len(arrRow) != int(6) {
			panic("Bad input")
		}

		arr = append(arr, arrRow)
	}

	result := hourglassSum(arr)

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
