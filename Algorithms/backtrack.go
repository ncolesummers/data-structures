// The problem is to identify the combination of elements in an array
// of 10 elements whose sum is equal to 18.  The findElementsWithSum 
// method recursively tries to find the combination.  Whenever the
// sum goes beyond the k target, it backtracks.

package main

import "fmt"

//findELementsWithSum of k from arr of size
func findElementsWithSum(arr[10]int, combinations[19]int, size, k, addValue, l, m int) int {
	var num int = 0
	if addValue > k { return -1 }
	if addValue == k {
		num = num + 1
		var p int = 0
		for p = 0; p < m; p ++ {
			fmt.Printf("%d,", arr[combinations[p]])
		}
		fmt.Println(" ")
	}
	var i int
	for i = l; i < size; i++ {
		combinations[m] = l
		findElementsWithSum(arr, combinations, size, k, addValue+arr[i], l, m + 1)
		l = l + 1
	}
	return num
}

func main() {
	var (
		arr = [10]int{1,4,7,8,3,9,2,4,1,8}
		addedSum int = 18
		combinations [19]int
	)
	findElementsWithSum(arr, combinations, 10, addedSum, 0, 0, 0)
}