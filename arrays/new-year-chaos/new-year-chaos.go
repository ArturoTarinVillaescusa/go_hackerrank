package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Complete the minimumBribes function below.
func minimumBribes(q []int32) int {
	cnt := 0;
	var temp int32
	for i := len(q) - 1; i >= 0; i-- {
		if(q[i] != int32(i + 1)) {
			if(((i - 1) >= 0) && q[i - 1] == int32(i + 1)) {
				cnt++;
				// swap q[i-1] and q[i]
				temp = q[i-1]
				q[i - 1] = q[i]
				q[i] = temp

			} else if(((i - 2) >= 0) && q[i - 2] == int32(i + 1)) {
				cnt += 2;
				q[i - 2] = q[i - 1];
				q[i - 1] = q[i];
				q[i] = int32(i + 1);
			} else {
				fmt.Println("Too chaotic");
				return -1;
			}
		}
	}
	fmt.Println(cnt);
	return cnt;
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

	tTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	t := int32(tTemp)

	for tItr := 0; tItr < int(t); tItr++ {
		nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
		checkError(err)
		n := int32(nTemp)

		qTemp := strings.Split(readLine(reader), " ")

		var q []int32

		for i := 0; i < int(n); i++ {
			qItemTemp, err := strconv.ParseInt(qTemp[i], 10, 64)
			checkError(err)
			qItem := int32(qItemTemp)
			q = append(q, qItem)
		}

		minimumBribes(q)
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

