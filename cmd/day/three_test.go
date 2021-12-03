package day

import "testing"

var day3TestInput = []string{
	"00100",
	"11110",
	"10110",
	"10111",
	"10101",
	"01111",
	"00111",
	"11100",
	"10000",
	"11001",
	"00010",
	"01010",
}

func TestDay3Part1(t *testing.T) {
	expectedGammaInt := int64(22)
	expectedEpsilonInt := int64(9)
	expectedGamma := "10110"
	gammaInt, epsilonInt, gammaStr := day3Part1(day3TestInput)
	if gammaStr != expectedGamma {
		t.Fatalf("expected %s got %s", expectedGamma, gammaStr)
	}
	if gammaInt != expectedGammaInt {
		t.Fatalf("expected %d got %d", expectedGammaInt, gammaInt)
	}
	if epsilonInt != expectedEpsilonInt {
		t.Fatalf("expected %d got %d", expectedEpsilonInt, epsilonInt)
	}
}

func TestDay3Part2(t *testing.T) {
	expectedO2Str := "10111"
	expectedO2Int := 23
	expectedCO2Str := "01010"
	expectedCO2Int := 10

	o2Int, o2Str, co2Int, co2Str := day3Part2(day3TestInput)
	if o2Int != expectedO2Int {
		t.Fatalf("expected %d got %d", expectedO2Int, o2Int)
	}
	if o2Str != expectedO2Str {
		t.Fatalf("expected %s got %s", expectedO2Str, o2Str)
	}

	if co2Int != expectedCO2Int {
		t.Fatalf("expected %d got %d", expectedCO2Int, co2Int)
	}
	if co2Str != expectedCO2Str {
		t.Fatalf("expected %s got %s", expectedCO2Str, co2Str)
	}
}
