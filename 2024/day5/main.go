package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	pageSets := getPageSets()
	pagePrecedence := getPagePrecedence()

	var problem1Sum, problem2Sum atomic.Int64
	var wg sync.WaitGroup
	for _, ps := range pageSets {
		wg.Add(1)
		go func() {
			if ps.isOrdered(pagePrecedence) {
				problem1Sum.Add(int64(ps.MiddleValue()))
			} else {
				ps.Order(pagePrecedence)
				problem2Sum.Add(int64(ps.MiddleValue()))
			}
			wg.Done()
		}()
	}
	wg.Wait()

	fmt.Println("Problem 1 answer: ", problem1Sum.Load())
	fmt.Println("Problem 2 answer: ", problem2Sum.Load())

}
