package main 
   
import "fmt"
   
func main() { 
   
  
    var y = 458 
       
 
    var p = &y 
   
    fmt.Println("Value stored in y before changing = ", y) 
    fmt.Println("Address of y = ", &y) 
    fmt.Println("Value stored in pointer variable p = ", p) 
  
    fmt.Println("Value stored in y(*p) Before Changing = ", *p) 
  
  
    *p = 500 
  
     fmt.Println("Value stored in y(*p) after Changing = ",y) 
  
} 