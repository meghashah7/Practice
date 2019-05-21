package main

// you can also use imports, for example:
import (
	"fmt"
	"sort"
)

// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func main() {
	ranks := []int{1}
	fmt.Println(Solution(ranks))
}

func Solution(A []int) int {
	// write your code in Go 1.4
	var b []int
	var count int
	length := len(A)

	for i := range A {
		if i+1 <= length {
			b = append(b, A[:i]...)
			b = append(b, A[i+1:]...)

			fmt.Println(b)
			if sort.IntsAreSorted(b) {
				count++

			}
			b = nil
		}

	}
	return count
}
