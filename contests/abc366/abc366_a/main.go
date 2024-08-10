package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

var scanner = bufio.NewScanner(os.Stdin)
var wtr = bufio.NewWriter(os.Stdout)

func main() {
	defer flush()

	in := readInts(3)
	res := "No"
	if nta(in) {
		res = "Yes"
	}
	print(res)
}

func nta(nta []int) bool {
	nokori := nta[0] - (nta[1] + nta[2])
	if nta[1] > nta[2] {
		if nta[1] > nta[2]+nokori {
			return true
		}
	} else {
		if nta[2] > nta[1]+nokori {
			return true
		}
	}
	return false
}

// --- init

const baseRune = 'a'

func init() {
	scanner.Buffer([]byte{}, math.MaxInt64)
	scanner.Split(bufio.ScanWords)
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

func readInts(len int) []int {

	inputs := make([]int, len)
	for i := 0; i < len; i++ {
		scanner.Scan()
		inputStr := scanner.Text()
		input, err := strconv.Atoi(inputStr)
		if err != nil {
			panic(err)
		}
		inputs[i] = input
	}

	return inputs
}

func int2() (int, int) {
	ints := readInts(2)
	return ints[0], ints[1]
}

// 文字列をアルファベット順のインデックスに変換して数値配列で読み取る
func readStrAsIndex() []int {
	scanner.Scan()
	s := scanner.Text()
	return sToIndex(s, baseRune)
}

// 文字列を基準文字に対する相対的な整数配列に変換
func sToIndex(s string, baseRune rune) []int {
	r := make([]int, len(s))
	for i, v := range s {
		r[i] = int(v - baseRune)
	}
	return r
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

// ユークリッド距離
func euclidean(x1, x2, y1, y2 int) float64 {
	return math.Sqrt(math.Pow(float64(x1-x2), 2) + math.Pow(float64(y1-y2), 2))
}

func flush() {
	e := wtr.Flush()
	if e != nil {
		panic(e)
	}
	if rec := recover(); rec != nil {
		fmt.Printf("panic: %+v\n", rec)
	}
}

func print(v ...interface{}) {
	_, e := fmt.Fprintln(wtr, v...)
	if e != nil {
		panic(e)
	}
}

func printInts(sl []int) {
	r := make([]string, len(sl))
	for i, v := range sl {
		r[i] = iToS(v)
	}
	print(strings.Join(r, " "))
}
