package main

import (
	"fmt"
	"math/rand"
	"sort"
)

func shuffle(slice []string) {
	for i := range slice {
		j := rand.Intn(i + 1)
		slice[i], slice[j] = slice[j], slice[i]
	}
}
func shuffle2(slice []string) {
	sort.Slice(slice, func(i, j int) bool {
		return rand.Intn(len(slice)) == 0
	})
}
func main() {
	fruits := []string{"apple", "banana", "cherry", "date", "elderberry", "fig",
		"grape", "honeydew", "pineapple", "plum", "quince", "raspberry", "strawberry", "tangerine", "watermelon"}
	fruitPointer := &fruits

	shuffle(*fruitPointer)
	fmt.Println("\nShuffled fruits:", fruits)

	shuffle2(*fruitPointer)
	fmt.Println("\nShuffled2 fruits:", fruits)

	sort.Strings(fruits[:])
	fmt.Println("\nSorted fruits: ", fruits)
}
