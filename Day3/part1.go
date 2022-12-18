package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func main() {
	
	// Read input file

	input, _ := os.Open("input.txt")
	defer input.Close()
	sc := bufio.NewScanner(input)
	
	// Declare variables

	var sumOfPriorities int

	for sc.Scan() {
		items := make(map[rune]bool)

		// Create a set with all the elements of the first component

		for _, itemLeftPart := range sc.Text()[:len(sc.Text())/2] {
			items[itemLeftPart] = true
		}

		// Range all the items of the second compartment

		for _, itemRightPart := range sc.Text()[len(sc.Text())/2:] {

			// If an item is in the first set, it is located in both compartments

			if items[itemRightPart] {
				sumOfPriorities += int(unicode.ToLower(itemRightPart) - 96) // 70+26
				if unicode.IsUpper(itemRightPart) {
					sumOfPriorities += 26
				}
				break
			}
		}
	}
	fmt.Println(sumOfPriorities) // Print answer
}
