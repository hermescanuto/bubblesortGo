package main

import (
	"fmt"
	"math/rand"
	"sort"
)

func BubbleSort(numbers []int) {
	for i := len(numbers); i > 0; i-- {
		for j := 1; j < i; j++ {
			if numbers[j-1] > numbers[j] {
				numbers[j], numbers[j-1] = numbers[j-1], numbers[j]
			}
		}
	}
}
func main() {
	//a := []int{2, 212, 3001, 14, 501, 7800, 9932, 33, 45, 45, 45, 91, 99, 37, 102, 102, 104, 106, 109, 106}
	a := rand.Perm(100)
	sort.Ints(a)
	//a = rand.Perm(100000)
	//BubbleSort(a)
	fmt.Println(a)
}
