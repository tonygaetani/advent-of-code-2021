package adventofcode

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"strconv"
)

func Day1() {
	home, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Failed to get user home directory")
		return
	}
	file, err := os.Open(path.Join(home, ".config", "advent-of-code-2021", "days", "1", "input"))
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

	// fmt.Printf("%d\n", len(input))

	// for i := 0; i < len(input); i++ {
	// 	if i == 0 {
	// 		// no change
	// 		continue
	// 	}

	// 	if input[i] > input[i-1] {
	// 		fmt.Printf("%d: %d is > %d\n", i, input[i], input[i-1])
	// 		increases++
	// 	} else {
	// 		fmt.Printf("%d: %d is not > %d\n", i, input[i], input[i-1])
	// 	}
	// }

	fmt.Printf("%d increases\n", increases)
}
