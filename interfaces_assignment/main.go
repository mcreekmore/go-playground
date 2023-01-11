package main

import (
	"fmt"
	"os"
)

func main() {
	fb, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println("ERROR:", err)
	}
	fmt.Println(string(fb))
}
