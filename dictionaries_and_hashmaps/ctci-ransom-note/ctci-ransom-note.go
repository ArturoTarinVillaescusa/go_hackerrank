package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Complete the checkMagazine function below.
func checkMagazine(magazine []string, note []string) bool {
	re := true

	magazineOccurences := make(map[string]int)

	for i := 0; i < len(magazine); i++ {
		position := magazine[i]
		magazineOccurences[position]++
	}

	for n := range note {
		noteFound := false
		for m := range magazineOccurences {
			if note[n] == m {
				if magazineOccurences[m] > 0 {
					magazineOccurences[m]--
					noteFound = true
				} else {
					re = false
				}
				break
			}
		}
		if noteFound == false {
			re = false
			break
		}
	}

	if re {
		fmt.Printf("Yes");
	} else {
		fmt.Printf("No");
	}

	return re

	/*
	magazineOccurences := make(map[string]int)
	noteOccurences := make(map[string]int)

	re := true

	for i := 0; i < len(magazine); i++ {
		position := magazine[i]
		magazineOccurences[position]++
	}


	for i := 0; i < len(note); i++ {
		position := note[i]
		noteOccurences[position]++
	}

	re = true
	for n := range noteOccurences {
		for m := range magazineOccurences {
			if n == m {
				if magazineOccurences[m] < 1 {
					re = false
				} else {
					magazineOccurences[m]--
					noteOccurences[n]--
				}
			}
		}
	}

	for n:= range noteOccurences {
		if noteOccurences[n] > 0 {
			re = false
		}
	}

	if re {
		fmt.Printf("Yes");
	} else {
		fmt.Printf("No");
	}

	return re
	*/

	/*
	for itN := range note {
		bFind := false
		for itM := range magazine {
			if itN == itM {
				bFind = true
				break
			}
			itM++
		}

		if bFind == false {
			fmt.Printf("No")
			return;
		}
	}
	fmt.Printf("Yes")
	return;
*/
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

	mn := strings.Split(readLine(reader), " ")

	mTemp, err := strconv.ParseInt(mn[0], 10, 64)
	checkError(err)
	m := int32(mTemp)

	nTemp, err := strconv.ParseInt(mn[1], 10, 64)
	checkError(err)
	n := int32(nTemp)

	magazineTemp := strings.Split(readLine(reader), " ")

	var magazine []string

	for i := 0; i < int(m); i++ {
		magazineItem := magazineTemp[i]
		magazine = append(magazine, magazineItem)
	}

	noteTemp := strings.Split(readLine(reader), " ")

	var note []string

	for i := 0; i < int(n); i++ {
		noteItem := noteTemp[i]
		note = append(note, noteItem)
	}

	checkMagazine(magazine, note)
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

