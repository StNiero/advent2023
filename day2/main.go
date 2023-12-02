package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	fileName := "input.txt"
	part1(fileName)
	part2(fileName)
}

func part1(fileName string) {
	sumOfPossible := 0
	var maxColors = map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}

	f, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(f)
	for i := 0; scanner.Scan(); i++ {
		line := scanner.Text()
		if line == "" {
			continue
		}

		reg, err := regexp.Compile(`\d+ \w+`)
		if err != nil {
			panic(err)
		}
		colorCounts := reg.FindAllString(line, -1)

		impossible := false
		for _, count := range colorCounts {
			split := strings.Split(count, " ")
			num, _ := strconv.Atoi(split[0])
			color := split[1]

			if num > maxColors[color] {
				impossible = true
				break
			}
		}
		if !impossible {
			sumOfPossible += i + 1
		}
	}
	fmt.Println(sumOfPossible)
}

func part2(fileName string) {
	sumOfPowers := 0

	f, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(f)
	for i := 0; scanner.Scan(); i++ {
		line := scanner.Text()
		if line == "" {
			continue
		}

		reg, err := regexp.Compile(`\d+ \w+`)
		if err != nil {
			panic(err)
		}
		colorCounts := reg.FindAllString(line, -1)

		var maxColors = map[string]int{
			"red":   0,
			"green": 0,
			"blue":  0,
		}

		for _, count := range colorCounts {
			split := strings.Split(count, " ")
			num, _ := strconv.Atoi(split[0])
			color := split[1]

			if num > maxColors[color] {
				maxColors[color] = num
			}
		}
		sumOfPowers += maxColors["red"] * maxColors["green"] * maxColors["blue"]
	}
	fmt.Println(sumOfPowers)
}
