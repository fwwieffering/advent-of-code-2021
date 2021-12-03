package day

import (
	"fmt"
	"sort"
	"strconv"
)

func day3() error {
	input, err := inputs.ReadFile("inputs/three")
	if err != nil {
		return err
	}
	rows := parseLinesStr(input)
	printPart(1)
	gamma, epsilon, _ := day3Part1(rows)
	part1Answer := gamma * epsilon
	fmt.Printf("gamma: %d * epsilon: %d = %d\n\n", gamma, epsilon, part1Answer)
	fmt.Printf("Answer: %d\n", part1Answer)
	printPart(2)
	o2, _, co2, _ := day3Part2(rows)
	part2Answer := o2 * co2
	fmt.Printf("o2: %d * co2: %d = %d\n\n", o2, co2, part2Answer)
	fmt.Printf("Answer: %d\n", part2Answer)
	return nil
}

func day3Part1(input []string) (int64, int64, string) {
	numRows := len(input)
	rowLen := len(input[0])

	runningTally := make([]int, rowLen)

	for _, r := range input {
		for idx, c := range r {
			intC, _ := strconv.Atoi(string(c))
			runningTally[idx] += intC
		}
	}

	gamma := ""
	for _, t := range runningTally {
		if t > (numRows / 2) {
			gamma += "1"
		} else {
			gamma += "0"
		}
	}

	gammaInt, _ := strconv.ParseUint(gamma, 2, 64)
	// no bitwise not operator, so use OR with 1111...
	return int64(gammaInt), int64(1<<rowLen-1) ^ int64(gammaInt), gamma
}

func day3Part2(input []string) (int, string, int, string) {
	o2Reading := func(rows []string, colIdx int) []string {
		numRows := len(rows)
		middle := numRows / 2
		// sort row by colidx value, then find most common value finding the first 1
		sort.Slice(rows, func(i, j int) bool {
			return rows[i][colIdx] < rows[j][colIdx]
		})
		for rowIdx := range rows {
			if rows[rowIdx][colIdx] == '1' {
				var newRows []string
				if rowIdx > middle {
					newRows = rows[:rowIdx]
				} else {
					newRows = rows[rowIdx:]
				}
				return newRows
			}
		}
		return nil
	}
	co2Reading := func(rows []string, colIdx int) []string {
		numRows := len(rows)
		middle := numRows / 2
		// sort row by colidx value, then find most common value finding the first 1
		sort.Slice(rows, func(i, j int) bool {
			return rows[i][colIdx] < rows[j][colIdx]
		})
		for rowIdx := range rows {
			if rows[rowIdx][colIdx] == '1' {
				var newRows []string
				if rowIdx <= middle {
					newRows = rows[:rowIdx]
				} else {
					newRows = rows[rowIdx:]
				}
				return newRows
			}
		}
		return nil
	}

	rowLen := len(input[0])
	// o2 reading
	rows := input
	colIdx := -1
	for len(rows) > 1 && colIdx < rowLen {
		colIdx++
		rows = o2Reading(rows, colIdx)
	}
	o2 := rows[0]
	o2Int, _ := strconv.ParseInt(o2, 2, 64)
	// co2 reading
	co2rows := input
	colIdx = -1
	for len(co2rows) > 1 && colIdx < rowLen {
		colIdx++
		co2rows = co2Reading(co2rows, colIdx)
	}
	co2 := co2rows[0]
	co2Int, _ := strconv.ParseInt(co2, 2, 64)

	return int(o2Int), o2, int(co2Int), co2
}
