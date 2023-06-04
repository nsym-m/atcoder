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

	var ok []int
	in := readInts()
	for i := 1; i <= in[0]; i++ {
		ints := strings.Split(strconv.Itoa(i), "")
		var sum int
		for _, v := range ints {
			i, _ := strconv.Atoi(v)
			sum += i
		}
		if in[1] <= sum && sum <= in[2] {
			ok = append(ok, i)
		}
	}

	var res int
	for _, v := range ok {
		res += v
	}

	fmt.Println(res)
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
