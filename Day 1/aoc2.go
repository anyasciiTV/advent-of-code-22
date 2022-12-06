package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	// Read the input file -- "input.txt"
	input, _ := os.Open("/d/advent-of-code/Day2/part2/input.txt")
	defer input.Close()
	sc := bufio.NewScanner(input)

	// Search for the max sum of calories
	maxCalories := 0
	currentCalories := 0

	// Scan file
	for sc.Scan() {
		snack, err := strconv.Atoi(sc.Text())
		currentCalories += snack

		/* If the error is different from "nil"
		then include an empty line */
		if err != nil {
			if currentCalories > maxCalories {
				maxCalories = currentCalories
			}

			// Begin with a new elf
			currentCalories = 0
		}
	}
	fmt.Println(maxCalories)
}
