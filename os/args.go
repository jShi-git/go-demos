package main

import (
	"fmt"
	"github.com/lealife/leacrawler"
	"os"
)

func main() {
	path := "d:\\goWorks" //保存根目录
	// url and the target path

	fmt.Println(os.Args[1])
	fmt.Println(path)

	//自定义目录参数
	if len(os.Args) == 3 {
		fmt.Println(os.Args[2])
		path = path + "\\" + os.Args[2]
	}

	leacrawler.Fetch(os.Args[1], path)
}
