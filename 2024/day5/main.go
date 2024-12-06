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
type Precedence map[int][]int

func main() {
	precedence := getPrecedence()
	pageSets := getPageSets()
	fmt.Println(len(pageSets))

	problemOne(precedence, pageSets)

	var sum int
OUTER:
	for _, ps := range pageSets {
		for i := len(ps) - 1; i >= 0; i-- {
			for j := i - 1; j >= 0; j-- {
				if !isCorrectOrder(ps[j], precedence[ps[i]]) {
					ps.Order(precedence)
					sum += ps[len(ps)/2]
					continue OUTER
				}
			}
		}
	}

	fmt.Println(pageSets[len(pageSets)-1])
	for _, p := range pageSets[len(pageSets)-1] {
		fmt.Println(p, precedence[p])
	}

	fmt.Println("problem 2 answer: ", sum)
}

func (ps PageSet) Order(precedence Precedence) {
	slices.SortFunc(ps, func(a, b int) int {
		if slices.Contains(precedence[a], b) {
			return 1
		}
		return -1
	})
}

func (ps PageSet) isOrdered

func problemOne(precedence Precedence, pageOrder []PageSet) {
	var sum, count int
OUTER:
	for _, p := range pageOrder {
		for i := len(p) - 1; i >= 0; i-- {
			arr := precedence[p[i]]
			for j := i - 1; j >= 0; j-- {
				if !isCorrectOrder(p[j], arr) {
					continue OUTER
				}
			}
		}

		count++
		sum += p[len(p)/2]
	}

	fmt.Println("1 count: ", count)
	fmt.Println("problem 1 answer: ", sum)
}

func isCorrectOrder(n int, arr []int) bool {
	return !slices.Contains(arr, n)
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

func getPrecedence() Precedence {
	f, err := os.Open("precedence.txt")
	if err != nil {
		panic("precedence: ouchie")
	}
	defer f.Close()

	precedence := make(Precedence)
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
