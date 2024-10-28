package main

import "fmt"

type Stack struct{
	num []int
}

func (s *Stack) push(val int){
	s.num=append(s.num, val)
}

func (s *Stack) pop() int{
	if len(s.num)==0{
		return -1
	}

	lastIndex:=len(s.num)-1
	popped:=s.num[lastIndex]
	s.num=s.num[:lastIndex]
	return popped
}
func main(){

	stack:=Stack{}
	stack.push(1)
	stack.push(9)
	stack.push(7)

	stack.pop()
	stack.pop()
	fmt.Println(stack)

}