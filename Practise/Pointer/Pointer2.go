package main 
  
import "fmt"
  
func main() { 

    y := 458 
   
    p := &y 
  
    fmt.Println("Value stored in y = ", y) 
    fmt.Println("Address of y = ", &y) 
    fmt.Println("Value stored in pointer variable p = ", p) 
} 