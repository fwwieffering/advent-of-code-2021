package day

import (
	"fmt"
	"strconv"
	"strings"
)

func dayTwo() error {
	data, _ := inputs.ReadFile("inputs/two")
	commands := parseLinesStr(data)

	part1 := day2Part1(commands)
	printPart(1)
	fmt.Printf("Product of x * z = %d\n\n", part1)
	fmt.Printf("Answer: %d\n", part1)

	part2 := day2Part2(commands)
	printPart(2)
	fmt.Printf("Product of x * z = %d\n\n", part2)
	fmt.Printf("Answer: %d\n", part2)

	return nil
}

func parseCommand(c string) (string, int) {
	splitC := strings.Split(c, " ")
	units, _ := strconv.Atoi(splitC[1])
	return splitC[0], units
}

func movePart1(xCur int, zCur int, command string, units int) (int, int) {
	switch command {
	case "forward":
		return xCur + units, zCur
	case "up":
		return xCur, zCur - units
	case "down":
		return xCur, zCur + units
	default:
		return xCur, zCur
	}
}

func movePart2(xCur int, zCur int, aimCur int, command string, units int) (int, int, int) {
	switch command {
	case "forward":
		return xCur + units, zCur + (aimCur * units), aimCur
	case "up":
		return xCur, zCur, aimCur - units
	case "down":
		return xCur, zCur, aimCur + units
	default:
		return xCur, zCur, aimCur
	}
}

func day2Part1(commands []string) int {
	x, z := 0, 0

	for _, cmd := range commands {
		direction, unit := parseCommand(cmd)
		x, z = movePart1(x, z, direction, unit)
	}
	return x * z
}

func day2Part2(commands []string) int {
	x, z, aim := 0, 0, 0

	for _, cmd := range commands {
		direction, unit := parseCommand(cmd)
		x, z, aim = movePart2(x, z, aim, direction, unit)
	}
	return x * z
}
