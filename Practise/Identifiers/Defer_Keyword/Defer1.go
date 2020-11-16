
package main 
  
import "fmt"
  

func mul(a1, a2 int) int { 
  
    res := a1 * a2 
    fmt.Println("Result: ", res) 
    return 0 
} 
  
func show() { 
    fmt.Println("Hello!, Divyam") 
} 
  

func main() { 
   
    mul(10, 10) 
  
    defer mul(20, 50) 
 
    show() 
} 