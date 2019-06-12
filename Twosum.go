func twoSum(nums []int, target int) []int {
    
  //  var retarray []int
    
  //  for i:=0;i< len(nums)-1 ; i++ {
        
  //      for j:= 1 ; j< len(nums) ; j++ {
            
 //           if nums[i]+nums[j] == target && i!=j{
 //               retarray= append(retarray,i)
 //               retarray= append(retarray,j)
 //               return retarray
 //           } else {
 //           continue
  //          }
                       
  //      }
        
//}
    m := make(map[int]int)
    for i:= range nums {
        m[nums[i]] =i
    }
    
    for j:= range nums {
        complement := target - nums[j]
        if _,ok := m[complement] ; ok && m[complement] !=j {
            return []int{j,m[complement]}
        }
    }
    
    return []int{0,0}
}
