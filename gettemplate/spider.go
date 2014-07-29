package main

import (
	"fmt"
	"github.com/lealife/leacrawler"
	"os"
)

func main() {
	path := "D:\\APMServ5.2.6\\www\\htdocs\\down\\tpls" //保存根目录

	//自定义目录参数
	if len(os.Args) == 3 {
		fmt.Println(os.Args[2])
		path = path + "\\" + os.Args[2]
	}

	if len(os.Args) > 1 {
		leacrawler.Fetch(os.Args[1], path)
	} else {
		fmt.Println("Error: --> usage: spider site-url* [save path]")
	}
}
