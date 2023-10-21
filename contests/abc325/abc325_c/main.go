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

var xi = []int{0, 0, -1, -1, -1, 1, 1, 1}
var xj = []int{-1, 1, -1, 0, 1, -1, 0, 1}

func main() {

	in := readInts()

	m := [][]string{}
	for i := 0; i < in[0]; i++ {
		s := readString()
		m = append(m, strings.Split(s, ""))
	}

	m2 := make(map[int][]int)
	var aa int
	for i, s := range m {
		for j, v := range s {
			if v == "." {
				continue
			}
			m2[aa] = append(m2[aa], -1, i, j)
			if isIncrement(m, i, j) {
				aa++
			}
			if i == 0 {
				fmt.Println("aa, j")
				fmt.Println(aa, j)
			}
			if i == 1 {
				fmt.Println(aa, j)
			}
		}
	}

	fmt.Printf("%+v\n", m2)
	fmt.Println(len(m2))
}

func isIncrement(m [][]string, i, j int) bool {

	// https://atcoder.jp/contests/abc325/editorial/7477
	for k := 0; k < 8; k++ {
		if isSharp(m, i+xi[k], j+xj[k]) {
			if i == 1 {
				fmt.Println("return false")
			}
			return false
		}
	}
	return true
}

func isSharp(m [][]string, i, j int) bool {
	if len(m) <= i || i < 0 {
		return false
	}
	if len(m[i]) <= j || j < 0 {
		return false
	}
	return m[i][j] == "#"
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

func reverseString(s string) string {
	rs := []rune(s)
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		rs[i], rs[j] = rs[j], rs[i]
	}
	return string(rs)
}
