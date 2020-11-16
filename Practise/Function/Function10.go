package main 
  
import "fmt"
  
 // Returning anonymous function  
 func GFG() func(i, j string) string{ 
     myf := func(i, j string)string{ 
          return i + j + "Divyam Gupta"
     } 
    return myf 
 } 
    
func main() { 
    value := GFG() 
    fmt.Println(value("Welcome ", "Back")) 
} 