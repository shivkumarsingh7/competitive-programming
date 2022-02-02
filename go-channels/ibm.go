// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"math/rand"
	"time"
)

// // map[int]emp{}
// m[0]= 1
// m[1]= 4
// m[3]= 4
// m[2]= 15

// type emp struct{
// }

// listing to 3 channel
// go routines with sleep method
//
func main() {
	//
	ch1 := make(chan int, 10)
	ch2 := make(chan int, 10)
	ch3 := make(chan int, 10)
	go func() {
		for {
			fmt.Println("first")
			for i := 0; i < 10; i++ {
				time.Sleep(time.Duration(rand.Intn(10)) * time.Second)
				ch1 <- i
			}
		}
	}()
	go func() {
		for {
			fmt.Println("second")
			for i := 10; i < 20; i++ {
				time.Sleep(time.Duration(rand.Intn(10)) * time.Second)
				ch2 <- i
			}
		}
	}()
	go func() {
		for {
			fmt.Println("third")
			for i := 20; i < 30; i++ {
				time.Sleep(time.Duration(rand.Intn(10)) * time.Second)
				ch3 <- i
			}
		}
	}()

	for {
		select {
		case msg := <-ch1:
			fmt.Println("Message from First Channel: ", msg)
		case msg := <-ch2:
			fmt.Println("Message from Second Channel: ", msg)
		case msg := <-ch3:
			fmt.Println("Message from Third Channel: ", msg)
		}
	}

}
