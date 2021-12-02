package day

import "fmt"

func dayOne() error {
	data, _ := inputs.ReadFile("inputs/one")
	depths := parseLinesInt(data)

	// part 1
	// iterate over list, checking for times depth increased
	countDepthIncreased := 0
	for idx := range depths {
		if idx > 0 {
			depthIncreased := depths[idx] > depths[idx-1]
			if depthIncreased {
				countDepthIncreased++
			}
		}
	}
	printPart(1)
	fmt.Printf("Depth increased %d times\n\n", countDepthIncreased)
	fmt.Printf("Answer: %d\n", countDepthIncreased)

	// part 2
	// iterate over lists, checking if sum of sliding window increased
	countWindowsIncreased, err := dayOnePart2(depths)
	if err != nil {
		return err
	}

	printPart(2)
	fmt.Printf("Window depth increased %d times\n\n", countWindowsIncreased)
	fmt.Printf("Answer: %d\n", countWindowsIncreased)

	return nil

}

func dayOnePart2(input []int) (int, error) {
	// windows is a map of start index window + the sum
	countWindowsIncreased := 0
	// these will hold our rolling window values
	one, two, three := 0, 0, 0
	for _, cur := range input {
		// have we set all positions in rolling window (hopefully none of the input values are 0)
		if one != 0 && two != 0 && three != 0 {
			// check if depth increased
			if two+three+cur > one+two+three {
				countWindowsIncreased++
			}
		}
		// move our window
		one, two, three = two, three, cur

	}
	return countWindowsIncreased, nil
}
