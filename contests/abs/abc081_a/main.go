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

/*
すぬけ君は 1,2,3 の番号がついた 3 つのマスからなるマス目を持っています。
各マスには 0 か 1 が書かれており、マス i には s (i) が書かれています。
すぬけ君は 1 が書かれたマスにビー玉を置きます。ビー玉が置かれるマスがいくつあるか求めてください。
*/

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

func readInt(scanner *bufio.Scanner) (int, error) {
	scanner.Scan()
	return strconv.Atoi(scanner.Text())
}

func readString(scanner *bufio.Scanner) string {
	scanner.Scan()
	return scanner.Text()
}

func readInts(scanner *bufio.Scanner) ([]int, error) {
	scanner.Scan()
	inputStr := scanner.Text()
	inputStrs := strings.Split(inputStr, " ")

	inputs := make([]int, len(inputStrs))
	for i, inputStr := range inputStrs {
		input, err := strconv.Atoi(inputStr)
		if err != nil {
			return nil, err
		}
		inputs[i] = input
	}

	return inputs, nil
}
