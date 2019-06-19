package main

// you can also use imports, for example:
import (
	"fmt"
	"sort"
)

// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

//This is from codility training questions

func main() {
	ranks := []int{1, 0}
	fmt.Println(Solution(ranks))
}

func Solution(ranks []int) int {
	// write your code in Go 1.4
	sort.Ints(ranks)
	fmt.Println(ranks)
	var totalreportees, rankcount int

	for i := range ranks {
		if i == 0 {
			rankcount = 1
			continue
		} else if ranks[i] == ranks[i-1] {
			rankcount++
		} else if ranks[i] == ranks[i-1]+1 {
			totalreportees = totalreportees + rankcount
			rankcount = 1
		} else {
			rankcount = 1
			continue
		}
	}
	return totalreportees

}
