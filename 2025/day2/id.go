package main

import (
	"fmt"
	"iter"
	"strconv"
	"strings"
)

type id int

func parseIdFromString(s string) (id, error) {
	n, err := strconv.Atoi(s)
	if err != nil {
		return 0, err
	}

	return parseIdFromInt(n)
}

func parseIdFromInt(n int) (id, error) {
	if !isValidId(n) {
		return 0, fmt.Errorf("invalid id")
	}

	return id(n), nil
}

func isValidId(n int) bool {
	s := strconv.Itoa(n)
	if len(s)%2 == 0 {
		return s[:len(s)/2] != s[len(s)/2:]
	}

	return true
}

func isValidStrictId(n int) bool {
	return !isInvalidStrictId(n)
}

func isInvalidStrictId(n int) bool {
	s := strconv.Itoa(n)
	l := len(s)
OUTER:
	for size := 1; size <= l/2; size++ {
		if l%size != 0 {
			continue
		}

		sub := s[:size]
		for i := size; i < l; i += size {
			if s[i:i+size] != sub {
				continue OUTER
			}
		}

		return true
	}

	return false
}

func parseIdRange(s string) (start, end int, err error) {
	parts := strings.Split(s, "-")
	if len(parts) != 2 {
		err = fmt.Errorf("invalid id range")
		return
	}

	// if these were parseIdFromString() the whole range would fail
	start, err = strconv.Atoi(parts[0])
	if err != nil {
		return
	}

	end, err = strconv.Atoi(parts[1])
	if err != nil {
		return
	}

	return
}

func idRangeSeq(start, end int) iter.Seq[int] {
	return iter.Seq[int](func(yield func(int) bool) {
		for i := start; i <= end; i++ {
			if !yield(i) {
				return
			}
		}
	})
}
