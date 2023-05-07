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

	in := readString(scanner)
	var res int
	for _, v := range in {
		if string(v) == "1" {
			res++
		}
	}

	fmt.Println(res)
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
