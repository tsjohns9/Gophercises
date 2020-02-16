package main

import (
	"flag"
	"fmt"
	"regexp"
)

func camelcase(s string) int32 {
	re := regexp.MustCompile(`[A-Z]`)
	res := re.FindAll([]byte(s), -1)
	return int32(len(res) + 1)
}

func main() {
	word := flag.String("word", "", "camelCase word")
	flag.Parse()
	if *word == "" {
		fmt.Println("A camelCase word is required")
		return
	}
	fmt.Println(camelcase(*word))
}
