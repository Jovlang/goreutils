package main

import (
	. "./shared"
	"os"
	"fmt"
)

func main() {
	dir, err := os.Getwd()
	Errhandler(err)
	fmt.Println(dir)
}
