package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

var scanner *bufio.Scanner

func main() {

	in1 := readInt(scanner)
	in2 := readInts(scanner)
	var res int

	fmt.Println(loop(in1, res, in2))
}

func loop(n, res int, ints []int) int {
	flg := true
	for i := 0; i < n; i++ {
		if ints[i]%2 == 0 {
			ints[i] = ints[i] / 2
		} else {
			flg = false
			break
		}
	}
	if flg {
		res++
		return loop(n, res, ints)
	} else {
		return res
	}
}

// --- init

func init() {
	var reader io.Reader

	if len(os.Args) > 1 {
		file, err := os.Open(os.Args[1])
		if err != nil {
			fmt.Println("ファイルを開けません:", err)
			os.Exit(1)
		}
		defer file.Close()
		reader = file
	} else {
		reader = os.Stdin
	}

	scanner = bufio.NewScanner(reader)
}

func readInt(scanner *bufio.Scanner) int {
	scanner.Scan()
	i, err := strconv.Atoi(scanner.Text())
	if err != nil {
		panic(err)
	}
	return i
}

func readString(scanner *bufio.Scanner) string {
	scanner.Scan()
	return scanner.Text()
}

func readInts(scanner *bufio.Scanner) []int {
	scanner.Scan()
	inputStr := scanner.Text()
	inputStrs := strings.Split(inputStr, " ")

	inputs := make([]int, len(inputStrs))
	for i, inputStr := range inputStrs {
		input, err := strconv.Atoi(inputStr)
		if err != nil {
			panic(err)
		}
		inputs[i] = input
	}

	return inputs
}
