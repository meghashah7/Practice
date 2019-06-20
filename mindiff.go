package solution

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(A []int) int {
    // write your code in Go 1.4
    var sum,leftsum, rightsum int
    var mindiff int
    for _,element :=range A {
        sum = sum+element
    }
   // fmt.Println(sum)
    
    leftsum = A[0]
    rightsum = sum-A[0]
    mindiff = rightsum-leftsum
    if mindiff <0 {
        mindiff= mindiff*-1
    }
    
    for i:=1 ;i <len(A) ;i++ {
     rightsum = rightsum-A[i]   
     leftsum = leftsum + A[i]
     diff := rightsum - leftsum  
     if diff <0 {
         diff = diff *-1
     } 
     if diff < mindiff {
        mindiff = diff
     }
    }
    return mindiff
}
