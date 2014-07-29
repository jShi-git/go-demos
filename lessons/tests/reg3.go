package main

import (
	"fmt"
	"strings"
)

func main() {

	src := string("D:/APMServ5.2.6/www/htdocs/down/tpls")

	src2 := string(`D:\APMServ5.2.6\www\htdocs\down\tpls`)

	src = strings.Replace(src, "/", "\\\\", -1)
	src2 = strings.Replace(src2, "\\", "\\\\", -1)

	fmt.Println(src)
	fmt.Println(src2)
}
