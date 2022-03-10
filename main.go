package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"strings"
)

//go:embed hello.txt
var hello string

func main() {
	s := bufio.NewScanner(strings.NewReader(hello))
	for s.Scan() {
		pair := strings.Split(s.Text(), " ")
		fmt.Println(pair)
	}
}
