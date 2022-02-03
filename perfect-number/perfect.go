package main

import "fmt"

func main() {
	//Enter your code here. Read input from STDIN. Print output to STDOUT
	var num int
	sum := 0
	fmt.Println("Enter a number :")
	fmt.Scanln(&num)

	for i := 1; i < num; i++ {
		if num%i == 0 {
			sum = sum + i
		}
	}
	if sum==num {
		fmt.Println("1")
	} else {
		fmt.Println("0")
	}

}
// 1,2,4,8,16,32