package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

var scanner *bufio.Scanner

func main() {
	in := readInt()
	var ints []int
	for i := 0; i < in; i++ {
		ints = append(ints, readInt())
	}
	ints = sortDesc(ints)
	var m []int
	for _, v := range ints {
		if len(m) > 0 && m[len(m)-1] == v {
			continue
		}
		m = append(m, v)
	}
	fmt.Println(len(m))
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

func sortDesc(ints []int) []int {
	sort.Sort(sort.Reverse(sort.IntSlice(ints)))
	return ints
}
