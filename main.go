package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("test")
	os.Create("b.txt")

}
