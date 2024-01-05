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

	a, b := int2()
	res := 0
	for i := 1; i < b; {
		i -= 1
		i += a
		res++
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

func int2() (int, int) {
	ints := readInts()
	return ints[0], ints[1]
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

func iToS(i int) string {
	return strconv.Itoa(i)
}

func sToI(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxu(a, b uint64) uint64 {
	if a > b {
		return a
	}
	return b
}

type Queue struct {
	data [][2]int
	size int
}

func NewQueue(len int) *Queue {
	return &Queue{
		data: make([][2]int, len, 0),
		size: 0,
	}
}

// キューにデータを追加する
func (q *Queue) Push(i [2]int) {
	q.data = append(q.data, i)
	q.size++
}

// キューの最新データを削除
func (q *Queue) Pop() bool {
	if q.IsEmpty() {
		return false
	}
	q.size--
	q.data = q.data[1:]
	return true
}

// キューの最新データを取得
func (q *Queue) Front() [2]int {
	return q.data[0]
}

func (q *Queue) IsEmpty() bool {
	return q.size == 0
}

func (q *Queue) String() string {
	return fmt.Sprint(q.data)
}
