package main

// you can also use imports, for example:
//import "fmt"
// import "os"
import "sort"


func main() {
  // This is to test the Solution function
  fmt.Println([]int{2,3,4,1,5,6})
  
}

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(A []int) int {
    // write your code in Go 1.4
  //sort internally is nlog(n) time complexity
    
    sort.Ints(A)
  //  fmt.Println(A)
    
    for i:=1; i<=len(A);i++{
        if A[i-1]==i {
         continue   
        } else {
            return 0
        }
        
        
    }
    return 1
    
    
    

}
