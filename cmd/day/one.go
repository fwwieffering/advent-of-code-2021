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
	var windows = map[int]int{}
	countWindowsIncreased := 0
	windowsCounted := 0
	for idx := range input {
		// check for sum of windows, and tabulate them
		for _, windowIdx := range []int{idx - 2, idx - 1} {
			if windowIdx >= 0 {
				curSum, ok := windows[windowIdx]
				if !ok {
					return 0, fmt.Errorf("something went wrong, we should always have a running sum for windowIdx %d at idx %d", windowIdx, idx)
				}
				windows[windowIdx] = curSum + input[idx]
			}
		}
		// start new window sum for current index
		windows[idx] = input[idx]
		// try to check if the most recently completed window (i-2) had a depth increase
		if idx-3 >= 0 {
			windowsCounted++
			// fmt.Printf("IDX %d: checking windows[%d] > windows[%d]: %d > %d\n", idx, idx-2, idx-3, windows[idx-2], windows[idx-3])
			windowDepthIncreased := windows[idx-2] > windows[idx-3]
			if windowDepthIncreased {
				countWindowsIncreased++
			}
		}
	}
	return countWindowsIncreased, nil
}
