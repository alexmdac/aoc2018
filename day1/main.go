package main

import (
	"fmt"
	"os"

	"github.com/alexmdac/aoc2018/util"
	"golang.org/x/tools/container/intsets"
)

func main() {
	soln, err := run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s: %s\n", os.Args[0], err)
		os.Exit(1)
	}

	fmt.Printf("%+v\n", soln)
}

type solution struct {
	sum              int
	firstRepeatedSum int
}

func run() (solution, error) {
	ints, err := util.ReadIntsFile()
	if err != nil {
		return solution{}, err
	}

	return solution{sum(ints), firstRepeatedSum(ints)}, nil
}

func sum(ints []int) int {
	s := 0
	for _, n := range ints {
		s += n
	}
	return s
}

func firstRepeatedSum(ints []int) int {
	s, seen := 0, intsets.Sparse{}
	for {
		for _, n := range ints {
			if seen.Has(s) {
				return s
			}
			seen.Insert(s)
			s += n
		}
	}
}
