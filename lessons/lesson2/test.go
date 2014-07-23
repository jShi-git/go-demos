package main

import "fmt"

func main() {
	var c complex64 = 5 + 5i
	//output: (5+5i)
	fmt.Printf("Value is: %v\n", c)

	//str := "test string"

	m := `hello
		    world`

	fmt.Printf("String is: %s\n", m)

	d := [...]int{4, 5, 6}
	e := append(d[:], 7)
	fmt.Printf("%d\n", e)

	slice := []byte{'a', 'b', 'c', 'd'}

	fmt.Printf("%s\n", slice[1:3])
	newslice := append(slice, 'e')
	fmt.Printf("%s\n", newslice)

	numbers := make(map[string]int)
	//var numbers = map[string]int
	numbers["one"] = 1  //赋值
	numbers["ten"] = 10 //赋值
	numbers["three"] = 3
	fmt.Printf("%d\n", numbers["ten"])

	// 初始化一个字典
	rating := map[string]float32{"C": 5, "Go": 4.5, "Python": 4.5, "C++": 2}
	// map有两个返回值，第二个返回值，如果不存在key，那么ok为false，如果存在ok为true
	csharpRating, ok := rating["C#"]
	if ok {
		fmt.Println("C# is in the map and its rating is ", csharpRating)
	} else {
		fmt.Println("We have no rating associated with C# in the map")
	}
}
