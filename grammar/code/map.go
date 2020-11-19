package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(countWord("how do you do"))
	res()
}

/**
统计一个字符串中每个单词出现的次数
*/
func countWord(str string) map[string]int {
	var strs []string = strings.Split(str, " ")
	m := make(map[string]int)
	for _, s := range strs {
		_, ok := m[s]
		if ok {
			m[s]++
		} else {
			m[s] = 1
		}
	}
	return m
}
func res() {
	type Map map[string][]int
	m := make(Map)
	s := []int{1, 2}
	s = append(s, 3)
	fmt.Printf("%+v\n", s)
	m["q1mi"] = s
	s = append(s[:1], s[2:]...)
	fmt.Printf("%+v\n", s)
	fmt.Printf("%+v\n", m["q1mi"])
}
