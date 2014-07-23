package main

import (
	"fmt"
	"github.com/Kenshin/curl"
	"io"
	"os"
	"testing"
)

func TestCurl(t *testing.T) {

	// curl.Get
	code, res, _ := Get("http://dist.u.qiniudn.com/latest/SHASUMS.txt")
	if code != 0 {
		return
	}

	// close
	defer res.Body.Close()

	// parse callback
	processFunc := func(content string, line int) bool {
		fmt.Printf("line is %v, content is %v", line, content)
		return false
	}

	// relad line
	if err := ReadLine(res.Body, processFunc); err != nil && err != io.EOF {
		fmt.Println(err)
	}

	// download
	code = New("http://dist.u.qiniudn.com/latest/node.exe", "node.exe", os.TempDir()+"/"+"node.exe")
	fmt.Printf("curl.New return code is %v\n", code)

}
