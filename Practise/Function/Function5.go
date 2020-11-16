package main 
  
import( 
    "fmt"
    "strings"
) 
  
 
func joinstr(element...string)string{ 
    return strings.Join(element, "-")
} 
  
func main() { 
    
  
   fmt.Println(joinstr()) 
     
   
   fmt.Println(joinstr("Divyam", "Gupta")) 
   fmt.Println(joinstr("Divyam", "Gupta", "Hindu")) 
   fmt.Println(joinstr("S", "E", "E", "M", "S")) 
     
} 