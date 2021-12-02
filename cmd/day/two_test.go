package day

import "testing"

var day2TestInput = []string{
	"forward 5",
	"down 5",
	"forward 8",
	"up 3",
	"down 8",
	"forward 2",
}

func TestDay2Part1(t *testing.T) {
	res := day2Part1(day2TestInput)
	expected := 150
	if res != expected {
		t.Fatalf("expected %d got %d", expected, res)
	}
}

func TestDay2Part2(t *testing.T) {
	res := day2Part2(day2TestInput)
	expected := 900
	if res != expected {
		t.Fatalf("expected %d got %d", expected, res)
	}
}
