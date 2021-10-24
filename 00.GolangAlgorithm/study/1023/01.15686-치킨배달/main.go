package main

import (
	"bufio"
	"fmt"
	"os"
)

type Data struct {
	y, x int
}

var (
	n, m, answer int
	board        [51][51]int
	reader       *bufio.Reader = bufio.NewReader(os.Stdin)
	writer       *bufio.Writer = bufio.NewWriter(os.Stdout)
	chicken      []Data
	D            []int
)

func printf(f string, a ...interface{}) { fmt.Fprintf(writer, f, a...) }
func scanf(f string, a ...interface{})  { fmt.Fscanf(reader, f, a...) }

func inits() {
	n = 0
	m = 0
	answer = 0
	scanf("%d %d\n", &n, &m)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			scanf("%d", &board[i][j])
		}
	}
}
func dfs(idx, cnt int) {
	if idx >= len(chicken) {
		return
	}
	if cnt == m {
		fmt.Println(D)
		return
	}
	D = append(D, idx)
	dfs(idx+1, cnt+1)
	D = D[:len(D)-1]
}
func main() {
	defer writer.Flush()
	T := 0
	scanf("%d", &T)
	for testCase := 1; testCase <= T; testCase++ {
		inits()
		dfs(0, 0)
		printf("#%d %d\n", testCase, answer)
	}
}
