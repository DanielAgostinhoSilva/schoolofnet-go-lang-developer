package main

import "fmt"

func main() {
	slice := make([]int, 2)
	slice = append(slice, 1, 2, 3, 5)

	fmt.Println(slice)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))
}
