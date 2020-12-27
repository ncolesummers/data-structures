package main

import "fmt"

// findElement method given array and k element
func findElement(arr[10] int, k int) bool {
	var i int
	for i = 0; i < 10; i++ {
		if arr[i] == k { return true }
	}
	return false
}

func main() {
	var (
		arr = [10]int{1, 4, 7, 8, 3, 9, 2, 4, 1, 8}
		check bool = findElement(arr, 10)
		check2 bool = findElement(arr, 9)
	)

	fmt.Println(check)
	fmt.Println(check2)
}