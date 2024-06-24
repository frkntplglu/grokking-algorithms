package main

import (
	"fmt"

	binarysearch "github.com/frkntplglu/grokking-algorithms/binary_search"
)

func main() {
	my_list := []int{1, 3, 5, 7, 9, 11, 13, 15}

	fmt.Println(binarysearch.BinarySearch(my_list, 9))
	fmt.Println(binarysearch.BinarySearch(my_list, 6))
}
