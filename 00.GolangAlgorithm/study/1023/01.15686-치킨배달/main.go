package main

import (
	"bufio"
	"fmt"
	"os"
)

var N, M, answer int

var board [51][51]int

type Data struct {
	y int
	x int
}

var chicken []Data // 치킨 데이터 저장하는 슬라이스
var D []int        //치킨집 뽑기
func init1() {
	N = 0
	M = 0
	answer = 0x7fffffff
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscanln(reader, &N, &M)
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			fmt.Fscanln(reader, &board[i][j])
			if board[i][j] == 2 {
				chicken = append(chicken, Data{i, j})
			}
		}
	}
}

// func dfs(idx, cnt int) {
// 	if idx >= len(chicken) {
// 		return
// 	}
// 	if cnt == M-1 {
// 	}
// 	D = append(D, idx)
// 	dfs(idx+1, cnt+1)
// 	D = D[:len(D)-1]

// }

func main() {
	T := 0
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	fmt.Fscanln(reader, &T)
	for testCase := 1; testCase <= T; testCase++ {
		init1()
		fmt.Fprintln(writer, chicken)
		//dfs(0, 0)
		fmt.Fprintln(writer, testCase, answer)
	}
	writer.Flush()
}
