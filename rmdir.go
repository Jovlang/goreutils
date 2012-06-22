package main

import (
	. "./shared"
	"os"
	"fmt"
	"io/ioutil"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Fprintf(os.Stderr, "usage: rmdir [directory]\n")
		os.Exit(1)
	}
	name := os.Args[1]
	stat, _ := os.Stat(name)
	if stat == nil {
		fmt.Fprintf(os.Stderr, "%v: no such directory", name)
		os.Exit(1)
	}
	if !stat.IsDirectory() {
		fmt.Fprintf(os.Stderr, "%v: not a directory", name)
		os.Exit(1)
	}
	files, err := ioutil.ReadDir(name)
	Errhandler(err)
	if len(files) != 0 {
		fmt.Fprintf(os.Stderr, "%v: directory not empty", name)
		os.Exit(1)
	}
	err = os.Remove(name)
	Errhandler(err)
}
