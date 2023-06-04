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

	a := readInt()
	b := readInt()
	c := readInt()
	x := readInt()

	var res int
	for i := 0; i <= a; i++ {
		for j := 0; j <= b; j++ {
			for k := 0; k <= c; k++ {
				if x == i*500+j*100+k*50 {
					res++
				}
			}
		}
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
