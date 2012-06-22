package main

import (
	. "./shared"
	"os"
	"fmt"
	"flag"
)

var (
	P = flag.Bool("p", false, "Make parent directories if necessary")
)

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) == 0 {
		fmt.Fprintf(os.Stderr, "usage: mkdir [directory]\n")
		flag.PrintDefaults()
		os.Exit(1)
	}
	dir := args[0]
	if *P {
		err := os.MkdirAll(dir, 0755)
		Errhandler(err)
	} else {
		err := os.Mkdir(dir, 0755)
		Errhandler(err)
	}
}
