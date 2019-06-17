/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

import "container/list"


func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    
  

    l3 :=list.New()
    var carryover,number int
    
    for l:=l1;l!=nil; l=l.Next {
        number =0
        if l.Next == nil  && l2.Next!=nil {
            number =  carryover+0 + l2.Val
            
        } else if l.Next !=nil && l2.Next == nil {
            number = carryover+l1.Val + 0 
         //else if l.Next == nil && l2.Next == nil {
        //    return l3
       // } 
        } else {
            number = carryover+ l.Val+l2.Val
            if number >=10 {
                carryover= number/10
                number = number%10
                fmt.Println("Carryover")
                fmt.Println(carryover)
             fmt.Println(number)
                
               
            } else {
                carryover =0
            }
            
        }
        
        //newNode.Val = number
        l3.PushBack(number)
        
        l2=l2.Next
  //      l3 = l3.Next
      
    //    fmt.Println(l3)
   //     fmt.Println("next")
    }
    
    
    fmt.Println(l3)
    var Head *ListNode
    Head.Val = int(l3.Front())
    Head.Next = l3.Next()
    return Head
}

