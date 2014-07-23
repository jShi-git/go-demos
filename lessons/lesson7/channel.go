package main

import (
	"fmt"
	"time"
)

func sum(a []int, c chan int) {
	total := 0
	for _, v := range a {
		total += v
	}
	c <- total // send total to c
}

func fibonacci(n int, c chan int) {
	x, y := 1, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func fibonacci2(c, quit chan int) {
	x, y := 1, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func main() {
	a := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	go sum(a[:len(a)/2], c)
	go sum(a[len(a)/2:], c)
	x, y := <-c, <-c // receive from c

	fmt.Println(x, y, x+y)

	cc := make(chan int, 10)
	go fibonacci(cap(cc), cc)
	for i := range cc {
		fmt.Println(i)
	}

	ccc := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-ccc)
		}
		quit <- 0
	}()
	fibonacci2(ccc, quit)

	c4 := make(chan int)
	o := make(chan bool)
	go func() {
		for {
			select {
			case v := <-c4:
				println(v)
			case <-time.After(5 * time.Second):
				println("timeout")
				o <- true
				break
			}
		}
	}()
	<-o
}
