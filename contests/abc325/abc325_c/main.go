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

	used := make(map[[2]int]*struct{})
	var aa int
	for i, s := range m {
		for j, v := range s {
			if v == "." || used[[2]int{i, j}] != nil {
				continue
			}
			q := NewQueue(0)
			q.Push([2]int{i, j})
			for !q.IsEmpty() {
				p := q.Front()
				q.Pop()
				ii := p[0]
				jj := p[1]
				for k := 0; k < 8; k++ {
					ix := ii + xi[k]
					jx := jj + xj[k]
					if 0 <= ix && ix < in[0] && 0 <= jx && jx < in[1] && m[ix][jx] == "#" && used[[2]int{ix, jx}] == nil {
						used[[2]int{ix, jx}] = &struct{}{}
						q.Push([2]int{ix, jx})
					}
				}
			}
			aa++
		}
	}
	fmt.Println(aa)
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
