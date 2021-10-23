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
func dfs(int idx, int cnt) {

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
