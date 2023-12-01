package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var consts = []string{
	"1",
	"2",
	"3",
	"4",
	"5",
	"6",
	"7",
	"8",
	"9",
	"one",
	"two",
	"three",
	"four",
	"five",
	"six",
	"seven",
	"eight",
	"nine",
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(f)

	sum := 0
	for i := 0; scanner.Scan(); i++ {
		line := scanner.Text()
		if line == "" {
			continue
		}

		checkPrefix := func(line string) string {
			for _, digit := range consts {
				if strings.HasPrefix(line, digit) {
					return digit
				}
			}
			return ""
		}

		checkSuffix := func(line string) string {
			for _, digit := range consts {
				if strings.HasSuffix(line, digit) {
					return digit
				}
			}
			return ""
		}

		first := ""
		for first == "" {
			firstVal := checkPrefix(line)
			if firstVal == "" {
				line = line[1:]
				continue
			}
			first = parse(firstVal)
		}

		last := ""
		for last == "" {
			lastVal := checkSuffix(line)
			if lastVal == "" {
				line = line[:len(line)-1]
				continue
			}
			last = parse(lastVal)
		}

		val, err := strconv.Atoi(first + last)
		if err != nil {
			fmt.Println(err)
		}
		sum += val
	}
	fmt.Println(sum)
}

func parse(s string) string {
	switch s {
	case "one":
		return "1"
	case "two":
		return "2"
	case "three":
		return "3"
	case "four":
		return "4"
	case "five":
		return "5"
	case "six":
		return "6"
	case "seven":
		return "7"
	case "eight":
		return "8"
	case "nine":
		return "9"
	default:
		return s
	}
}
