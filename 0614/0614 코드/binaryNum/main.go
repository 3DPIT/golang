package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.O_APPEND, os.O_CREATE, os.O_RDONLY, os.O_WRONLY, os.O_RDWR)
	fmt.Printf("%016b\n", os.O_RDONLY)
	fmt.Printf("%016b\n", os.O_WRONLY)
	fmt.Printf("%016b\n", os.O_RDWR)
	fmt.Printf("%016b\n", os.O_CREATE)
	fmt.Printf("%016b\n", os.O_APPEND)

	fmt.Println(os.FileMode(0700))
	fmt.Println(os.FileMode(0070))
	fmt.Println(os.FileMode(0007))

	fmt.Println(os.FileMode(17))
	fmt.Println(os.FileMode(249))
	fmt.Println(os.FileMode(1000))

	for i := 0; i <= 19; i++ {
		fmt.Printf("%3d: %04o\n", i, i)
	}
	fmt.Printf("%09b %s\n", 0000, os.FileMode(0000))
	fmt.Printf("%09b %s\n", 0111, os.FileMode(0111))
	fmt.Printf("%09b %s\n", 0222, os.FileMode(0222))
	fmt.Printf("%09b %s\n", 0333, os.FileMode(0333))
	fmt.Printf("%09b %s\n", 0444, os.FileMode(0444))
	fmt.Printf("%09b %s\n", 0555, os.FileMode(0555))
	fmt.Printf("%09b %s\n", 0666, os.FileMode(0666))
	fmt.Printf("%09b %s\n", 0777, os.FileMode(0777))

}
