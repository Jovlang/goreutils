package shared

import (
	"os"
	"fmt"
)

func Errhandler(err os.Error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v: %v\n", os.Args[0], err)
		os.Exit(1)
	}
}
