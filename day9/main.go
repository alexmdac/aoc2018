package main

import (
	"container/ring"
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

type solution struct {
	part1, part2 int
}

func run() (*solution, error) {
	return &solution{
		// 471 players; last marble is worth 72026 points
		part1: game(471, 72026),
		part2: game(471, 7202600),
	}, nil
}

func game(players, last int) int {
	scores := make([]int, players)

	r := ring.New(1)
	r.Value = 0

	for m := 1; m <= last; m++ {
		if m%23 == 0 {
			for i := 0; i < 7; i++ {
				r = r.Prev()
			}
			scores[m%players] += m + r.Value.(int)
			r = r.Prev()
			r.Unlink(1)
			r = r.Next()
		} else {
			s := ring.New(1)
			s.Value = m
			r.Next().Link(s)
			r = r.Next().Next()
		}
	}

	max := 0
	for _, s := range scores {
		if s > max {
			max = s
		}
	}
	return max
}
