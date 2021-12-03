package adventofcode

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"strconv"
	"strings"
)

func getInput(day string) (*os.File, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Failed to get user home directory")
		return nil, err
	}
	path := path.Join(home, ".config", "advent-of-code-2021", "days", day, "input")
	file, err := os.Open(path)
	if err != nil {
		fmt.Printf("Failed to open file %s\n", path)
		return nil, err
	}
	return file, nil
}

func Day1() {
	file, err := getInput("1")
	if err != nil {
		fmt.Println("Failed to get input for day 1")
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan() // skip the first line, since there's nothing to compare it to
	text := scanner.Text()
	previous, err := strconv.Atoi(text)
	if err != nil {
		fmt.Printf("%s cannot be converted to an integer\n", text)
		return
	}

	increases := 0
	for scanner.Scan() {
		text := scanner.Text()
		current, err := strconv.Atoi(text)
		if err != nil {
			fmt.Printf("%s cannot be converted to an integer\n", text)
			return
		}

		if current > previous {
			increases++
		}

		previous = current
	}

	fmt.Printf("%d increases\n", increases)
}

func Day2() {
	day2Part1()
	day2Part2()
}

func day2Part1() {
	horizontal, depth := 0, 0
	processInput("2", func(command string, value int) {
		switch command {
		case "forward":
			horizontal += value
			break
		case "down":
			depth += value
			break
		case "up":
			depth -= value
			break
		default:
			fmt.Printf("unknown command %s", command)
			return
		}
	})

	fmt.Printf("%d position\n", horizontal*depth)
}

func day2Part2() {
	horizontal, depth, aim := 0, 0, 0
	processInput("2", func(command string, value int) {
		if value == 0 {
			fmt.Println("NOPE NOPE NOPE")
			return
		}
		switch command {
		case "forward":
			horizontal += value
			depth += aim * value
			break
		case "down":
			aim += value
			break
		case "up":
			aim -= value
			break
		default:
			fmt.Printf("unknown command %s", command)
			return
		}
	})

	fmt.Printf("%d position\n", horizontal*depth)
}

func processInput(day string, fn func(string, int)) {
	file, err := getInput(day)
	if err != nil {
		fmt.Println("Failed to get input for day 2")
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		split := strings.Split(text, " ")
		command := split[0]
		value, err := strconv.Atoi(split[1])
		if err != nil {
			fmt.Println("Failed to parse input")
			return
		}
		fn(command, value)
	}
}
