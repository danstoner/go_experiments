package main

import (
       "fmt"
       "time"
)

func main() {
     var time1 = time.Now()  // this assigns the actual timestamp to time1
     fmt.Println("The time is", time1)
     fmt.Println("Welcome to the playground!")
     fmt.Println(time1)
     fmt.Println(time1)
     fmt.Println("sleeping for 2....")
     time.Sleep(2 * time.Second)
     fmt.Println(time1)

     fmt.Print("Looping...")
     for i := 0; i < 10000; i++ {
     	 if i % 100 == 0 {
	      	 fmt.Print(".")
		 }
	 time.Sleep(1 * time.Millisecond)
	 }

	 fmt.Println()


	 fmt.Println("The program took", time.Since(time1))
}
