package main

import (
	. "./shared"
	"os"
	"fmt"
	"flag"
	"io/ioutil"
)

var (
	L = flag.Bool("l", false, "Show additional file info")
	S = flag.Bool("s", false, "With -l, show size in kilobytes")
)

func prefix(f *os.FileInfo) (s string) {
	if *L {
		s = fmt.Sprintf("%-4v %-4v %-4v ", f.Uid, f.Gid, size(f))
	}
	return
}

func size(f *os.FileInfo) (s string) {
	if *S {
		s = fmt.Sprintf("%dK", f.Size/1024)
	} else {
		s = fmt.Sprintf("%d", f.Size)
	}
	return
}

func symbol(f *os.FileInfo) (s string) {
	switch {
	case f.IsDirectory():
		s = "/"
	case f.IsFifo():
		s = "|"
	case f.IsSocket():
		s = "="
	case f.IsSymlink():
		target, err := os.Readlink(f.Name)
		Errhandler(err)
		s = " -> " + target
	}
	return
}

func main() {
	flag.Parse()
	args := flag.Args()

	dir := "."
	if len(args) > 0 {
		dir = args[0]
	}

	err := os.Chdir(dir)
	Errhandler(err)

	files, err := ioutil.ReadDir(".")
	Errhandler(err)

	for _, f := range files {
		fmt.Printf("%s%s%s\n", prefix(f), f.Name, symbol(f))
	}
}
