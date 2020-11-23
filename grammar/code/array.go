package main

import "fmt"

func main() {
	arr := [5]int{1, 3, 5, 7, 8}
	fmt.Println(sumArray(arr))
	fmt.Println(findIndex(arr, 8))
}

func sumArray(arr [5]int) int {
	var sum = 0
	for _, value := range arr {
		sum += value
	}
	return sum
}
func findIndex(arr [5]int, sum int) [][2]int {
	var res [][2]int
	i := 0
	for i1, v1 := range arr {
		for i2, v2 := range arr[i1+1:] {
			if v1+v2 == sum {
				res = append(res, [2]int{i1, i2 + i1 + 1})
				i++
			}
		}
	}
	return res
}
