package main

import "fmt"

func main() {
	sum := 0
	for index := 0; index < 10; index++ {
		sum += index
	}
	fmt.Println("sum is equal to ", sum)

	sum2 := 1
	for sum2 < 100 {
		fmt.Println("sum is equal to ", sum2)
		sum2 += sum2
	}
	fmt.Println("sum is equal to ", sum2)

	rating := map[string]float32{"C": 5, "Go": 4.5, "Python": 4.5, "C++": 2}
	for k, v := range rating {
		fmt.Printf("map's equal:%s:%.2f\n", k, v)
	}

	d := [...]int{4, 5, 6}
}
