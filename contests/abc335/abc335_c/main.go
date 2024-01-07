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

	n, q := int2()
	l := make([][2]int, n)
	j := 0
	for i := n; i > 0; i-- {
		// 初期位置の配列をリバースで作成
		l[j] = [2]int{i, 0}
		j++
	}

	for k := 0; k < q; k++ {
		query := readss() // クエリ取得
		if query[0] == "1" {
			x := 0
			y := 0
			if query[1] == "R" {
				x = 1
			}
			if query[1] == "L" {
				x = -1
			}
			if query[1] == "U" {
				y = 1
			}
			if query[1] == "D" {
				y = -1
			}

			// 頭の位置情報を作成
			head := [2]int{l[len(l)-1][0] + x, l[len(l)-1][1] + y}
			l = l[1:]           // スライスの先頭を削除
			l = append(l, head) // 頭をスライスの末尾に追加
		} else {
			p := sToI(query[1])                  // string to int
			m := l[len(l)-p]                     // 指定の位置を取得
			print(iToS(m[0]) + " " + iToS(m[1])) // 指定フォーマットで出力
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

func readss() []string {
	scanner.Scan()
	inputStr := scanner.Text()
	inputStrs := strings.Split(inputStr, " ")

	inputs := make([]string, len(inputStrs))
	for i, inputStr := range inputStrs {
		inputs[i] = inputStr
	}

	return inputs
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

func sortDesc(ints []int) {
	sort.Sort(sort.Reverse(sort.IntSlice(ints)))
}

func sortAsc(ints []int) {
	sort.Ints(ints)
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

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func print(a ...any) {
	for _, v := range a {
		fmt.Printf("%+v\n", v)
	}
}

func f(format string, a ...any) string {
	return fmt.Sprintf(format, a...)
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
