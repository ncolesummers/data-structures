// This example comes from the Learn Data Structures and Algorithms with Golang text
package main

import (
	"fmt"
	"container/list"
)

func main() {
	var intList list.List
	intList.PushBack(11)
	intList.PushBack(23)
	intList.PushBack(34)

	for element := intList.Front(); element != nil; element=element.Next() {
		fmt.Println(element.Value.(int))
	}
}

