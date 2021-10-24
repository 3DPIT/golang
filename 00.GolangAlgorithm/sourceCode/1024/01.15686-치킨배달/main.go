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
	n, m    int
	answer                = 0
	board                 = [51][51]int{}
	reader  *bufio.Reader = bufio.NewReader(os.Stdin)
	writer  *bufio.Writer = bufio.NewWriter(os.Stdout)
	chicken []Data
	home    []Data
	D       [15]int
)

func printf(f string, a ...interface{}) { fmt.Fprintf(writer, f, a...) }
func scanf(f string, a ...interface{})  { fmt.Fscanf(reader, f, a...) }

func main() {
	defer writer.Flush()
	inits()
	dfs(0, 0)
	printf("%d\n", answer)
}
func inits() {

	n = 0
	m = 0
	answer = 0x7fffffff
	scanf("%d %d\n", &n, &m)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fmt.Fscan(reader, &board[i][j])
			if 1 == board[i][j] {
				home = append(home, Data{i, j})
			}
			if 2 == board[i][j] {
				chicken = append(chicken, Data{i, j})
			}

		}
	}
	// fmt.Println(chicken)
	// fmt.Println(home)
}

func abs(num int) int {
	if num < 0 {
		num = num * -1
	}
	return num
}

func dfs(idx, cnt int) {

	if idx > len(chicken) {
		return
	}
	if cnt == m {
		// fmt.Println(D)
		sum := 0
		// fmt.Println("homeLen", len(home))
		// fmt.Println("chichken", len(chicken))
		homeMin := 0
		for i := 0; i < len(home); i++ {
			homeMin = 0x7fffffff
			for j := 0; j < len(chicken); j++ {
				// fmt.Println(i, j)
				if D[j] == 1 {
					absNum := abs(home[i].y-chicken[j].y) + abs(home[i].x-chicken[j].x)
					// fmt.Println("y", home[i].y, chicken[j].y)
					// fmt.Println("x", home[i].x, chicken[j].x)
					// fmt.Println("absNum", absNum)
					if homeMin > absNum {
						homeMin = absNum
						// fmt.Println("homeMin", homeMin)
					}
				}
			}
			sum += homeMin

			// fmt.Println("homMin fix", homeMin)
		}
		// fmt.Println("sum", sum)
		// fmt.Println("answer", answer)

		if answer > sum {
			answer = sum
		}
		return
	}

	D[idx] = 1
	dfs(idx+1, cnt+1)
	D[idx] = 0
	dfs(idx+1, cnt)
}
