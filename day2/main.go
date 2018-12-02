package main

import (
	"fmt"
	"os"
	"strings"

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
	checksum int
	common   string
}

func run() (solution, error) {
	strs, err := util.ReadLinesFile()
	if err != nil {
		return solution{}, nil
	}

	return solution{
		checksum: checksum(strs),
		common:   common(strs),
	}, nil
}

func checksum(strs []string) int {
	twos, threes := 0, 0
	for _, s := range strs {
		cs := map[rune]int{}
		for _, c := range s {
			cs[c]++
		}
		hasTwo, hasThree := false, false
		for _, c := range cs {
			if c == 2 {
				hasTwo = true
			}
			if c == 3 {
				hasThree = true
			}
		}
		if hasTwo {
			twos++
		}
		if hasThree {
			threes++
		}
	}
	return twos * threes
}

func common(strs []string) string {
	seen := map[string]struct{}{}
	for _, str := range strs {
		for i := 0; i < len(str)-1; i++ {
			s := str[:i] + "_" + str[i+1:]
			if _, ok := seen[s]; ok {
				return strings.Replace(s, "_", "", -1)
			}
			seen[s] = struct{}{}
		}
	}
	return ""
}
