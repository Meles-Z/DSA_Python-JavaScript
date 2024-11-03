package main

import "fmt"

func greaterNumber(list []int) int {
	greater := list[0]
	fmt.Println("Before update:", greater)
	if len(list) == 0 {
		fmt.Println("List is empty")
		return 0
	}

	for i := 0; i < len(list); i++ {
		if list[i] > greater {
			greater = list[i]
		}
	}
	return greater
}

func main() {
	list := []int{8, 23, 4, 56, 90}
	fmt.Println("Greater value is", greaterNumber(list))

}
