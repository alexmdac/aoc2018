package main

import (
	"fmt"
	"os"

	"github.com/alexmdac/aoc2018/util"
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
	f, err := os.Open("input.txt")
	if err != nil {
		return solution{}, err
	}

	ints, err := util.ReadInts(f)
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
	s, m := 0, map[int]struct{}{}
	for {
		for _, n := range ints {
			if _, ok := m[s]; ok {
				return s
			}
			m[s] = struct{}{}
			s += n
		}
	}
}
