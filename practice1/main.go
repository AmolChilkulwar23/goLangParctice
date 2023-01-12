package main

import "fmt"

func add(x int, y int) int {
	var out = x + y
	return out
}

func calc(x, y int) (int, int, int) {
	var out1 = x - y
	out2 := x * y
	out3 := x / y
	return out1, out2, out3
}

func main(){
	num1 := 123
	num2 := 890
	result := add(num1,num2)
	result1, result2, result3 := calc(num1, num2)
	fmt.Println(result,"\n",result1,"\n",result2,"\n",result3)
}