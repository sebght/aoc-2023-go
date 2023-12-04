package utils

import (
	"fmt"
	"os"
)

func Check(e error) {
	if e != nil {
		fmt.Fprintf(os.Stderr, "error: %v", e)
		os.Exit(1)
	}
}
