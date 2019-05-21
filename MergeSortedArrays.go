package main

import (
	"fmt"
	"sort"
)

func main() {
	var array1, array2 []int
	array1 = []int{0, 3, 4, 28}
	array2 = []int{4, 6, 31}

	array1 = append(array1, array2...)
	sort.Ints(array1)
	fmt.Println(array1)
}
