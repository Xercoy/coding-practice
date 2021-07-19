package main

import "fmt"

func sort(list []int) {
	for current := 0; current <= len(list)-2; current++ {
		for next := current + 1; next < len(list); next++ {
			if list[next] < list[current] {
				tmp := list[current]       // save current node value
				list[current] = list[next] // overwrite with next
				list[next] = tmp
			}
		}
	}
}

func main() {
	cases := [][]int{
		[]int{5, 4, 6, 23, 5, 2, 5, 4, 6, 9}, // random
		[]int{0},                             // one element
		[]int{},                              // empty element
		[]int{9, 5},                          // two
		[]int{54, 22, 1},                     // three
	}

	for _, c := range cases {
		fmt.Printf("\n\n\n\nSelection Sort; Before: %v", c)

		sort(c)

		fmt.Printf("\nSelection Sort; After: %v", c)
	}
}
