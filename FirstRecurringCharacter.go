package main

import "fmt"

func main() {
//Below array is for testing purpose
	//givenArray := []int{2, 5, 1, 2, 3, 5, 1, 2, 4}
	//givenArray := []int{2, 1, 1, 2, 3, 5, 1, 2, 4}
	givenArray := []int{1, 2, 3, 4}
	fmt.Println(givenArray)

	fmt.Println(Solution(givenArray))

}

func Solution(givenArray []int) int {

	m := make(map[int]int)
	for _, key := range givenArray {
		m[key]++
		if m[key] == 2 {
			return key
		}
	}
	return 0
}
