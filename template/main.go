package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func process(inputs []int) (outputs []int, err error) {
	// ここに実際の処理を記述してください。
	fmt.Println(inputs)
	return outputs, nil
}

func main() {
	var reader io.Reader

	if len(os.Args) > 1 {
		file, err := os.Open(os.Args[1])
		if err != nil {
			fmt.Println("ファイルを開けません:", err)
			os.Exit(1)
		}
		defer file.Close()
		reader = file
	} else {
		reader = os.Stdin
	}

	scanner := bufio.NewScanner(reader)

	scanner.Scan()
	n, err := strconv.Atoi(scanner.Text())
	if err != nil {
		fmt.Println("入力が整数ではありません:", scanner.Text())
		os.Exit(1)
	}

	scanner.Scan()
	inputStr := scanner.Text()
	inputStrs := strings.Split(inputStr, " ")

	if len(inputStrs) != n {
		fmt.Println("入力された数と実際の要素数が一致しません。")
		os.Exit(1)
	}

	inputs := make([]int, len(inputStrs))
	for i, inputStr := range inputStrs {
		input, err := strconv.Atoi(inputStr)
		if err != nil {
			fmt.Println("入力が整数ではありません:", inputStr)
			os.Exit(1)
		}
		inputs[i] = input
	}

	outputs, err := process(inputs)
	if err != nil {
		fmt.Println("エラーが発生しました:", err)
		os.Exit(1)
	}

	for _, output := range outputs {
		fmt.Println(output)
	}
}
