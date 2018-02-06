package main

import "fmt"

func main() {
	for i:=1;i<=10;i++{
		fmt.Println(fib(i))
	}
}

func fib(n int) int {
    f1  := 1
    f2  := 1
    
    var r int
    
    for i := 0; i < n; i++ {
        r = f1
        f1 = f2
        f2 = r + f1
    }
    
    return r
}
