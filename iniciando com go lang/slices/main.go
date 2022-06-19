package main

import "fmt"

func main() {
	slice := make([]int, 2, 5)
	//slice = append(slice, 1, 2, 3, 5)

	fmt.Println(slice)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))

	//for i := 0; i < 20; i++ {
	//	slice = append(slice, i)
	//	fmt.Println("Len: ", len(slice), " Cap: ", cap(slice))
	//}

	sliceTest := slice
	slice = append(slice, 1, 2, 3, 4)
	slice[0] = 10
	fmt.Println(slice)
	fmt.Println(sliceTest)
}
