func max(num1, num2 int) int {
	
	result int
 
	if (num1 > num2) {
	   result = num1
	} else {
	   result = num2
	}
	return result 
 }

 func main()  {
	
	fmt.Println("enter the value: ", max(4,5))

 }

 