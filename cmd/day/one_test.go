package day

import "testing"

var testInput = []int{199, 200, 208, 210, 200, 207, 240, 269, 260, 263}

func TestPartTwo(t *testing.T) {
	res, _ := dayOnePart2(testInput)
	if res != 5 {
		t.Fatalf("expected: %d, got %d", 5, res)
	}
}
