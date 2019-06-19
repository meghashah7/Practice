package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, playground")
	//Below array represents 0 - means car passing to east 1 means car passing to west. Count total pairs.
	nums := []int{0,0,0,0,0,0,1}
	m := make(map[int]int)
	count:=0
	
	for i := range nums {
	    m[nums[i]]++
	
	      if nums[i] == 1 {
	       if m[0] >0 {
	         count = m[0] + count
	        } 
	
	}
	}
	if count > 1000000000 {
	fmt.Println(count)
	} else {
	fmt.Println(count)
	}
	
	
	
}
