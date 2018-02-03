package main

import (
	"fmt"
)

func main() {
	arr := []int{6, 5, 1, 3, 8, 4, 7, 9, 2}
	//arr := []int{5, 2, 3, 1, 6, 4}

	quickSort(arr)

	fmt.Println(arr)
}
