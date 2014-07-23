package main

import (
	"fmt"
	"os"
)

func main() {
	os.Mkdir("test1", 0777)
	os.MkdirAll("test1/test2", 0777)
	err := os.Remove("test1")
	if err != nil {
		fmt.Println(err)
	}
	os.RemoveAll("test1")
}
