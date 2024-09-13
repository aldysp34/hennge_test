package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	numberOfTestCase, _ := strconv.Atoi(scanner.Text())

	input := make([]string, 0)
	for i := 0; i < numberOfTestCase*2; i++ {
		scanner.Scan()
		input = append(input, scanner.Text())
	}

	var processCase func(int)
	processCase = func(i int) {
		if i >= numberOfTestCase {
			return
		}

		result := handleCase(input, i*2)
		fmt.Println(result)
		processCase(i + 1)
	}

	processCase(0)
}

func sumSquares(numbers []int) int {
	if len(numbers) == 0 {
		return 0
	}

	current := 0
	if numbers[0] >= 0 {
		current = numbers[0] * numbers[0]
	}

	return current + sumSquares(numbers[1:])
}

func parseNumbers(line string) []int {
	parts := strings.Fields(line)
	numbers := make([]int, 0, len(parts))
	for _, part := range parts {
		num, _ := strconv.Atoi(part)
		numbers = append(numbers, num)
	}
	return numbers
}

func handleCase(input []string, index int) int {
	numbers := parseNumbers(input[index+1])
	return sumSquares(numbers)
}
