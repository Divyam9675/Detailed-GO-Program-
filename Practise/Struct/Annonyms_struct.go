package main 
  
import "fmt"
  
// Main function 
func main() { 
  
    Element := struct { 
        name      string 
        branch    string 
        language  string 
        Particles int
    }{ 
        name:      "Pikachu", 
        branch:    "ECE", 
        language:  "C++", 
        Particles: 498, 
    } 
   
    fmt.Println(Element) 
} 