package main

import (
	"fmt"
	"os"
	"unicode"

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
	numRemaining int
	shortest     int
}

func run() (*solution, error) {
	lines, err := util.ReadLinesFile()
	if err != nil {
		return nil, err
	}
	if len(lines) != 1 {
		return nil, fmt.Errorf("expected single line, but got %d", len(lines))
	}
	input := lines[0]
	return &solution{
		numRemaining: numRemaining(input),
		shortest:     shortest(input),
	}, nil
}

func numRemaining(s string) int {
	return numRemainingForRunes([]rune(s))
}

func numRemainingForRunes(rs []rune) int {
	for {
		end := 0
		for i := 1; i < len(rs); i++ {
			if end >= 0 {
				fst, snd := rs[end], rs[i]
				if unicode.IsLower(fst) != unicode.IsLower(snd) &&
					unicode.ToLower(fst) == unicode.ToLower(snd) {
					end--
					continue
				}
			}
			end++
			rs[end] = rs[i]
		}
		if end == len(rs)-1 {
			return len(rs)
		}
		rs = rs[:end+1]
	}
}

func shortest(s string) int {
	rs := []rune(s)
	min := len(rs)

	low := map[rune]bool{}
	for _, r := range rs {
		low[unicode.ToLower(r)] = true
	}

	for toRemove := range low {
		updated := make([]rune, 0, len(rs))
		for _, r := range rs {
			if unicode.ToLower(r) != toRemove {
				updated = append(updated, r)
			}
		}
		n := numRemainingForRunes(updated)
		if n < min {
			min = n
		}
	}
	return min
}
