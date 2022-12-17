package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func main() {

	// Read the input file
	input, _ := os.Open("input.txt")
	defer input.Close()
	sc := bufio.NewScanner(input)

	// Declare variables

	var sumOfPriorities int

	for sc.Scan() {

		// Create three sets with elements of the elves

		itemsFirst := createSetOfItems(sc.Text())
		sc.Scan()
		itemsSecond := createSetOfItems(sc.Text())
		sc.Scan()
		itemsThird := createSetOfItems(sc.Text())

		// For each item of the respective elf,
		// check if the other elves also have the item

		for itemFirstElf := range itemsFirst {
			if itemsSecond[itemFirstElf] && itemsThird[itemFirstElf] {
				sumOfPriorities += int(unicode.ToLower(itemFirstElf) - 96) // Note : 70 + sumOfPriorities
				if unicode.IsUpper(itemFirstElf) {
					sumOfPriorities += 26
				}
				break
			}
		}
	}
	fmt.Println(sumOfPriorities)
}

// Create function and map rune to the case

func createSetOfItems(items string) (set map[rune]bool) {
	set = make(map[rune]bool)
	for _, item := range items {
		set[item] = true
	}
	return
}
