package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func multiplesOf3And5TimingOut(n int) {
	sum := 0

	for i := 1; i < n; i++ {
		if i%3 == 0 {
			//fmt.Println(i)
			sum += i
			continue
		}
		if i%5 == 0 {
			//fmt.Println(i)
			sum += i
			continue
		}
	}

	fmt.Println(sum)
}

func addition(first, step, count int) int {
	var last int = first + step*(count-1)
	return (first + last) * count / 2
}

func multiplesOf3And5(n int) int {
	x := 3
	y := 5
	var sumX int = addition(0, x, (n-1)/x+1)
	var sumY int = addition(0, y, (n-1)/y+1)
	var mcm int = minimoComunMultiplo(x, y)
	var sumBoth int = addition(0, mcm, (n-1)/mcm+1)
	result := sumX + sumY - sumBoth
	fmt.Println(result)
	return result
}

func minimoComunMultiplo(a, b int) int {
	return a * b / maximoComunDivisor(a, b)
}

func maximoComunDivisor(a, b int) int {
	if b == 0 {
		return a
	} else {
		return maximoComunDivisor(b, a%b)
	}
}

func main() {
	//Enter your code here. Read input from STDIN. Print output to STDOUT

	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	qTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	q := int32(qTemp)

	for qItr := 0; qItr < int(q); qItr++ {
		nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
		checkError(err)
		n := int(nTemp)

		// 2
		// 10
		// 100

		multiplesOf3And5(n)
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
