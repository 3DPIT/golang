package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	N := 0
	fmt.Fscanf(reader, "%d", &N)
	for i := N; i >= 1; i-- {
		fmt.Fprintf(writer, "%d\n", i)
	}
	writer.Flush()
}
