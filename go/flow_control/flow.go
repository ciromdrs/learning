package main

import ("fmt"
	"math")

func main() {
	fmt.Println("the answer is", Sqrt(37))
}

func Sqrt(x float64) float64 {
	z := x / 2

	old := z
	z -= (z*z - x) / (2*z)
	change := z - old
	fmt.Println(z, "-", old, "=", change)

	for math.Abs(change) > 0.01 {
		old := z
		z -= (z*z - x) / (2*z)
		change = z - old
		fmt.Println(z, "-", old, "=", change)
	}
	

	return z
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
