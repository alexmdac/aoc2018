package util

import (
	"io"
	"io/ioutil"
	"strconv"
	"strings"
)

// ReadLines returns an array of strings, one per line of the input from the
// reader.
func ReadLines(r io.Reader) ([]string, error) {
	bytes, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}

	s := strings.TrimSpace(string(bytes))
	strs := strings.Split(s, "\n")

	return strs, nil
}

// ReadInts returns an array of ints, one per line of the input from the reader.
func ReadInts(r io.Reader) ([]int, error) {
	strs, err := ReadLines(r)
	if err != nil {
		return nil, err
	}

	ints := make([]int, len(strs))
	for i := range strs {
		ints[i], err = strconv.Atoi(strs[i])
		if err != nil {
			return nil, err
		}
	}

	return ints, nil
}
