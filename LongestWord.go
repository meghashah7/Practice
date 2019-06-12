package main
import "fmt"
import "strings"
import "regexp"
import "log"

func LongestWord(sen string) string { 

  // code goes here   
  // Note: feel free to modify the return type of this function 
  
  reg, err := regexp.Compile("[^a-zA-Z ]+")
  if err !=nil {
      log.Fatal(err)
  }
  new := reg.ReplaceAllString(sen,"")
 words := strings.Split(new," ")
  previousLength := len(words[0])
  i:=0
  var maxIndex int
  for i=1; i< len(words) ;i++ {
      if len(words[i]) >  previousLength {
          previousLength = len(words[i])
          maxIndex =i
      }
  }
  
  return words[maxIndex]
            
}

func main() {

    fmt.Println(LongestWord(readline()))

}
