package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
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
	sum   int
	value int
}

type input struct {
	ns []int
	i  int
}

func newInput(ns []int) *input {
	return &input{ns: ns, i: 0}
}

func (in *input) nextInt() (int, error) {
	if in.i >= len(in.ns) {
		return 0, errors.New("no more ints")
	}
	n := in.ns[in.i]
	in.i++
	return n, nil
}

func (in *input) walkTree(f func(childvals []int, meta []int) int) (int, error) {
	nchild, err := in.nextInt()
	if err != nil {
		return 0, err
	}

	nmeta, err := in.nextInt()
	if err != nil {
		return 0, err
	}

	childvals := make([]int, nchild)
	for i := 0; i < nchild; i++ {
		c, err := in.walkTree(f)
		if err != nil {
			return 0, err
		}
		childvals[i] = c
	}

	meta := make([]int, nmeta)
	for i := 0; i < nmeta; i++ {
		m, err := in.nextInt()
		if err != nil {
			return 0, err
		}
		meta[i] = m
	}

	return f(childvals, meta), nil
}

func run() (*solution, error) {
	bytes, err := ioutil.ReadFile("input.txt")
	if err != nil {
		return nil, err
	}

	strs := strings.Split(string(bytes), " ")
	ns := make([]int, len(strs))

	for i := range strs {
		n, err := strconv.Atoi(strs[i])
		if err != nil {
			return nil, err
		}
		ns[i] = n
	}

	sum, err := sum(ns)
	if err != nil {
		return nil, err
	}

	value, err := value(ns)
	if err != nil {
		return nil, err
	}

	return &solution{sum: sum, value: value}, nil
}

func sum(ns []int) (int, error) {
	in := newInput(ns)
	sum, err := in.walkTree(func(childvals []int, meta []int) int {
		sum := 0
		for _, c := range childvals {
			sum += c
		}
		for _, m := range meta {
			sum += m
		}
		return sum
	})
	if err != nil {
		return 0, err
	}
	return sum, nil
}

func value(ns []int) (int, error) {
	in := newInput(ns)
	value, err := in.walkTree(func(childvals []int, meta []int) int {
		if len(childvals) == 0 {
			sum := 0
			for _, m := range meta {
				sum += m
			}
			return sum
		}

		value := 0
		for _, m := range meta {
			idx := m - 1
			if idx >= 0 && idx < len(childvals) {
				value += childvals[idx]
			}
		}
		return value
	})
	if err != nil {
		return 0, err
	}
	return value, nil
}
