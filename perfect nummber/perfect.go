package main
import "fmt"

func main() {
 //Enter your code here. Read input from STDIN. Print output to STDOUT
    var num int
    sum := 0
    fmt.Println("Enter a number :")
    fmt.Scanln(&num)
    
    for i:=1 ; i < num ; {
        if num%i==0 {
            sum = sum + 1
            i++
        }
        if sum==num{
            fmt.Println(1)
        }
        fmt.Println(0)
    }
}