package main 
  
import "fmt"
  
// Main function 
func main() { 
      
    // using maps 
    mmap := map[int]string{ 
        22:"Hello", 
        33:"Divyam", 
        44:"Gupta", 
    } 
    for key, value:= range mmap { 
       fmt.Println(key, value)  
    } 
    
} 