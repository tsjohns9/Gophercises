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
	prob := flag.String("problem", "camelCase", "camelCase or caesarCipher")
	word := flag.String("word", "", "Word for the challenge")
	flag.Parse()

	if *word == "" {
		fmt.Println("A word is required for the challenge")
		return
	}

	if *prob == "camelCase" {
		fmt.Println(camelcase(*word))
	}
	if *prob == "caesarCipher" {
		fmt.Println(caesarCipher(*word, int32(4)))
	}
}

func caesarCipher(s string, k int32) string {
	var str string
	for _, ch := range s {
		fmt.Println(ch)
	}
	return str
}
