package day

import (
	"bufio"
	"bytes"
	"embed"
	"fmt"
	"strconv"
)

//go:embed inputs/*
var inputs embed.FS

var DayFunc = map[int]func() error{
	1: dayOne,
}

func parseLinesInt(lines []byte) []int {
	var res []int

	lineReader := bufio.NewScanner(bytes.NewReader(lines))

	for lineReader.Scan() {
		t := lineReader.Text()
		if len(t) > 0 {
			intT, err := strconv.Atoi(t)
			if err == nil {
				res = append(res, intT)
			}
		}
	}
	return res
}

func printPart(part int) {
	fmt.Printf("------------------------------ PART %d ------------------------------\n\n", part)
}
