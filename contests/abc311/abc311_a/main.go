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

	scanner.Scan()
	_, err := strconv.Atoi(scanner.Text())
	if err != nil {
		panic(err)
	}
	in := readString()
	var a bool
	var b bool
	var c bool

	for i, v := range in {
		if string(v) == "A" {
			a = true
		}
		if string(v) == "B" {
			b = true
		}
		if string(v) == "C" {
			c = true
		}
		if a && b && c {
			fmt.Println(i + 1)
			return
		}
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

func sortDesc(ints []int) []int {
	sort.Sort(sort.Reverse(sort.IntSlice(ints)))
	return ints
}
