package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/michaelgov-ctrl/aoc/2025/utils"
)

type ingredient struct {
	id int
}

func (in *ingredient) isFresh(ranges []idRange) bool {
	for _, r := range ranges {
		if r.start <= in.id && in.id <= r.end {
			return true
		}
	}

	return false
}

func parseIngredient(s string) (ingredient, error) {
	id, err := strconv.Atoi(s)
	if err != nil {
		return ingredient{}, err
	}

	return ingredient{id: id}, nil
}

func parseIngredients(path string) []ingredient {
	var ingrs []ingredient
	utils.ForEachLineInFile(path, func(s string) {
		ingr, err := parseIngredient(s)
		if err != nil {
			panic(fmt.Sprintf("%v: %s", err, s))
		}

		ingrs = append(ingrs, ingr)
	})

	return ingrs
}

type idRange struct {
	start int
	end   int
}

func parseIdRange(s string) (idRange, error) {
	parts := strings.Split(s, "-")
	if len(parts) != 2 {
		return idRange{}, fmt.Errorf("invalid id range")
	}

	// if these were parseIdFromString() the whole range would fail
	start, err := strconv.Atoi(parts[0])
	if err != nil {
		return idRange{}, err
	}

	end, err := strconv.Atoi(parts[1])
	if err != nil {
		return idRange{}, err
	}

	return idRange{
		start: start,
		end:   end,
	}, nil
}

func parseIdRanges(path string) []idRange {
	var ranges []idRange

	utils.ForEachLineInFile(path, func(s string) {
		ir, err := parseIdRange(s)
		if err != nil {
			panic(fmt.Sprintf("%v: %s", err, s))
		}

		ranges = append(ranges, ir)
	})

	return ranges
}

func mergeRanges(r1, r2 idRange) (idRange, error) {
	if r1.start > r2.end || r2.start > r1.end {
		return idRange{}, fmt.Errorf("ranges do not overlap")
	}

	return idRange{
		start: min(r1.start, r2.start),
		end:   max(r1.end, r2.end),
	}, nil
}

func reduceRanges(ranges []idRange) []idRange {
	ranges = append(make([]idRange, 0, len(ranges)), ranges...) // copy

	found := true
	for found {
		found = false
	OUTER:
		for i := 0; i < len(ranges); i++ {
			for j := i + 1; j < len(ranges); j++ {
				merged, err := mergeRanges(ranges[i], ranges[j])
				if err != nil {
					continue
				}

				ranges = append(ranges[:i], ranges[i+1:]...) // delete indices i & j and append them as merged
				ranges = append(ranges[:j-1], ranges[j:]...) // subtract 1 since slice has shrunk by 1
				ranges = append(ranges, merged)

				found = true
				break OUTER
			}
		}
	}

	return ranges
}
