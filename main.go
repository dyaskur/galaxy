package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

type Variable struct {
	galaxyUnit   map[string]string
	galaxyCredit map[string]float32
}

func translator(input string) string {
	var output string

	r, _ := regexp.Compile(`([.a-z]+) is ([.A-Z]+)$`)

	// Using FindStringSubmatch you are able to access the
	// individual capturing groups
	for index, match := range r.FindStringSubmatch(input) {
		fmt.Printf("[%d] %s\n", index, match)
	}

	return output
}
func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		text, _, err := reader.ReadLine()
		str := string(text)
		if str == "" || err != nil {
			break
		}
		fmt.Println(translator(str))
		//fmt.Println(str)
	}
}
