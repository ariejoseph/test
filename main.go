package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("test")
	f, err := os.Create("b.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("success", f.Name())
}
