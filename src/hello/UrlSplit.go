package main

import (
	"fmt"
	"strings"
)

var level int = 3

func main() {
	str := "spi.lakala.com/markting-activity-redpacket-1111/js/flexible.js"
	//var str2 string
	fmt.Println(str)

	s := strings.Split(str, "/")
	if len(s) > level {
		str = strings.Join(s[0:level], "/")
	}
	fmt.Println(str)
}
