package main

import ("fmt"
		 "time"
)		 

func display1(chan1 chan string ){

	time.Sleep(3 * time.Second)
	chan1 <- "Welcome to channel 1"

}

func display2(chan2 chan string){


	time.Sleep(1*time.Second)
	chan2 <- "Welcome to channel 2"
		
}


func main(){

	R1 := make(chan string)
	R2 := make(chan string)

	go display1(R1)
	go display2(R2)

	select {
	case op1 := <- R1:
		fmt.Println(op1)
	
	case op2 := <- R2:
	     fmt.Println(op2)
	
	}

}