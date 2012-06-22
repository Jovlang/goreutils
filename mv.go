package main

import (
	. "./shared"
	"os"
	"path"
	"fmt"
	"bufio"
	"flag"
)

var (
	V = flag.Bool("v", false, "Verbose mode")
	I = flag.Bool("i", false, "Prompt to overwrite")
)

func exists(name string) bool {
	stat, _ := os.Stat(name)
	return stat != nil
}

func isdir(name string) bool {
	stat, _ := os.Stat(name)
	return stat.IsDirectory()
}

func main() {
	flag.Parse()
	args := flag.Args()

	if len(args) < 2 {
		fmt.Fprintf(os.Stderr, "usage: mv oldfile(s) [newfile|directory]\n")
		flag.PrintDefaults()
		os.Exit(1)
	}

	source := args[:len(args)-1]
	dest := args[len(args)-1]

	// If there are more than 1 source files, dest must be a directory.
	if len(source) > 1 {
		if !exists(dest) || !isdir(dest) {
			fmt.Fprintf(os.Stderr, "%v: not a directory\n", dest)
			os.Exit(1)
		}
	}

	r := bufio.NewReader(os.Stdin)

	for _, file := range source {
		if !exists(file) {
			fmt.Fprintf(os.Stderr, "%v: no such file or directory\n", file)
			os.Exit(1)
		}

		var fulldest string
		if exists(dest) && isdir(dest) {
			fulldest = dest + "/" + path.Base(file)
		} else {
			fulldest = dest
		}

		if *V {
			fmt.Printf("%v -> %v\n", file, fulldest)
		}

		if *I && exists(fulldest) {
			fmt.Printf("Overwrite %v -> %v? ", file, fulldest)
			chars, _, err := r.ReadLine()
			if err != nil || chars[0] != 'y' {
				continue
			}
		}

		err := os.Rename(file, fulldest)
		Errhandler(err)
	}
}
