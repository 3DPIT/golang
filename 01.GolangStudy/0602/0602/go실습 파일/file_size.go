package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	fileinfo, err := os.Stat("my.text")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(fileinfo.Size())
}
