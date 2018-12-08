package main

import (
	"fmt"
	"os"
)

func main() {
	soln, err := run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s: %s\n", os.Args[0], err)
		os.Exit(1)
	}

	fmt.Printf("%+v\n", soln)
}

type solution struct{}

func run() (*solution, error) {
	return &solution{}, nil
}
