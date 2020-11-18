package main

import "fmt"

func main() {
	text := "hello Go 语言"
	count := 0
	for _, s := range text {
		if s > 256 {
			fmt.Printf("%c\n", s)
			count++
		}
	}

	fmt.Println(count)
}
