package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var scanner *bufio.Scanner

func main() {

	in1 := readInt()
	in2 := readInts()
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
	scanner = bufio.NewScanner(os.Stdin)
}

func readInt() int {
	scanner.Scan()
	i, err := strconv.Atoi(scanner.Text())
	if err != nil {
		panic(err)
	}
	return i
}

func readString() string {
	scanner.Scan()
	return scanner.Text()
}

func readInts() []int {
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
