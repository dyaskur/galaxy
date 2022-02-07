package main

import (
	"bufio"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	//Init a new System Struct
	var s System
	s.Init()
	for {
		text, _, err := reader.ReadLine()
		str := string(text)
		if str == "" || err != nil {
			break
		}
		s.Translate(str)
		//fmt.Println(str)
	}
}
