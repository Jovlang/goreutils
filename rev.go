package main

import (
	"os"
	"utf8"
	"fmt"
	"bufio"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	for {
		chars, _, err := r.ReadLine()
		line := string(chars)
		if err == os.EOF {
			return
		}
		o := make([]int, utf8.RuneCountInString(line))
		i := len(o)
		for _, c := range line {
			i--
			o[i] = c
		}
		fmt.Println(string(o))
	}
}
