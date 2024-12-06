package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type PageSet []int
type PagePrecedence map[int][]int

func (ps PageSet) MiddleValue() int {
	return ps[len(ps)/2]
}

func (ps PageSet) Order(p PagePrecedence) {
	slices.SortFunc(ps, func(a, b int) int {
		if slices.Contains(p[a], b) {
			return 1
		}
		return -1
	})
}

func (ps PageSet) isOrdered(p PagePrecedence) bool {
	for i := len(ps) - 1; i >= 0; i-- {
		for j := i - 1; j >= 0; j-- {
			if slices.Contains(p[ps[i]], ps[j]) {
				return false
			}
		}
	}

	return true
}

func getPageSets() []PageSet {
	f, err := os.Open("pageSets.txt")
	if err != nil {
		panic("pageSets: ouchies")
	}
	defer f.Close()

	var res []PageSet
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		words := strings.Split(scanner.Text(), ",")
		nums := make([]int, len(words))
		for i, w := range words {
			n, err := strconv.Atoi(w)
			if err != nil {
				panic(fmt.Sprintf("failed to parse: %s to int", w))
			}

			nums[i] = n
		}

		res = append(res, nums)
	}

	return res
}

func getPagePrecedence() PagePrecedence {
	f, err := os.Open("precedence.txt")
	if err != nil {
		panic("precedence: ouchie")
	}
	defer f.Close()

	precedence := make(PagePrecedence)
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		pages := strings.Split(scanner.Text(), "|")
		if len(pages) != 2 {
			panic("too many pages on line")
		}

		nums := make([]int, len(pages))
		for i, p := range pages {
			n, err := strconv.Atoi(p)
			if err != nil {
				panic(fmt.Sprintf("failed to parse: %s to int", p))
			}

			nums[i] = n
		}

		arr, ok := precedence[nums[0]]
		if !ok {
			precedence[nums[0]] = []int{nums[1]}
		} else {
			precedence[nums[0]] = append(arr, nums[1])
		}
	}

	return precedence
}
