package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Hello, Megha")
	A := []int{-1, 0, 1, 100000}
	i := solution(A)
	fmt.Println(i)
}

func solution(A []int) int {
	sort.Ints(A)
	fmt.Println(A)
	fmt.Println(len(A))
	var i int
	for i = 1; i <= 1000000; i++ {
		//	n = sort.SearchInts(A, i)
		j := sort.Search(len(A), func(j int) bool { return A[j] >= i })
		fmt.Printf("i: %d, j:%d,\n", i, j)
		if j < len(A) && A[j] == i {
			fmt.Println("if", j, len(A), i)
			// i is present at the A[j]
		} else {
			return i
		}

	}
	return i
}
