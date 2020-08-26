package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func fibbonacciWord(s1 string, s2 string, n int64) string {
	var tmp string
	var i int64

	for i = 2; i <= n; i++ {
		tmp = s1
		s1 += s2
		s2 = tmp
	}

	return s1
}

func main() {
	//Enter your code here. Read input from STDIN. Print output to STDOUT

	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	qTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	q := int32(qTemp)

	for qItr := 0; qItr < int(q); qItr++ {
		lineTemp := strings.Split(readLine(reader), " ")

		s1 := lineTemp[0]
		s2 := lineTemp[1]
		nTemp, err := strconv.ParseInt(lineTemp[2], 10, 64)
		checkError(err)
		n := int64(nTemp)

		fmt.Printf("%v\n", s1)
		fmt.Printf("%v\n", s2)
		fmt.Printf("%v\n", n)

		// 2
		// 1415926535 8979323846 35
		// 1415926535897932384626433832795028841971693993751058209749445923078164062862089986280348253421170679 8214808651328230664709384460955058223172535940812848111745028410270193852110555964462294895493038196 104683731294243150

		fibbonacciWord(s1, s2, n)
	}

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
