package main

import (
	"fmt"
)

func noTwoSlash(url string) string {
	output := make([]rune, len(url), len(url))
	i := 0
	foundSlash := false
	for _, c := range url {
		if c == '/' {
			if !foundSlash {
				foundSlash = true
				output[i] = c
			} else {
				continue
			}
		} else {
			if foundSlash {
				foundSlash = false
			}
			output[i] = c
		}
		i++
	}
	return string(output[:i])
}

func main() {
	fmt.Println("hello world!")
}
