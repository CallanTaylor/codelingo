package main

import (
	"os"
)

func main() {

	file, _ := os.Create("First.txt")
	file.Close()

}
