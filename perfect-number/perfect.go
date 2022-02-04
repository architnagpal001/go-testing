package main

import "fmt"

func main() {
	//Enter your code here. Read input from STDIN. Print output to STDOUT
	var n int
	sum := 0
	fmt.Println("Enter a number :")
	fmt.Scanln(&n)

	for i := 1; i < n; i++ {
		if n%i == 0 {
			sum = sum + i
		}
	}
	if sum==n {
		fmt.Println("1")
	} else {
		fmt.Println("0")
	} 

}