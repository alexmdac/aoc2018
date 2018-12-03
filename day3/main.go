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
	overlaps int
	goodID   int
}

type claim struct {
	id, x, y, w, h int
}

func (c *claim) parse(s string) error {
	_, err := fmt.Sscanf(s, "#%d @ %d,%d: %dx%d",
		&c.id, &c.x, &c.y, &c.w, &c.h)
	return err
}

func run() (solution, error) {
	lines, err := util.ReadLinesFile()
	if err != nil {
		return solution{}, err
	}

	claims := make([]claim, len(lines))
	for i := range lines {
		if err := claims[i].parse(lines[i]); err != nil {
			return solution{}, err
		}
	}

	var c [1000][1000]int
	for _, claim := range claims {
		for i := 0; i < claim.h; i++ {
			for j := 0; j < claim.w; j++ {
				c[i+claim.y][j+claim.x]++
			}
		}
	}
	overlaps := 0
	for _, row := range c {
		for _, x := range row {
			if x > 1 {
				overlaps++
			}
		}
	}

	goodID := -1
	for _, claim := range claims {
		good := true
		for i := 0; i < claim.h; i++ {
			for j := 0; j < claim.w; j++ {
				if c[i+claim.y][j+claim.x] != 1 {
					good = false
				}
			}
		}
		if good {
			goodID = claim.id
		}
	}

	return solution{overlaps: overlaps, goodID: goodID}, nil
}
