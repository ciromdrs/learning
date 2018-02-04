package main

import (
    "fmt"
    "math/rand"
    "math"
    "math/cmplx"
)

func add(x int, y int) int {
    return x + y
}

func swap(x,y string) (string,string) {
    return y,x
}

func split(sum int) (x,y int) {
    x = sum * 4 / 9
    y = sum - x
    return
}

var c, python, java bool = true,true,true

func main() {
    
    fmt.Println("My favorite number is", rand.Intn(10))
    
    fmt.Printf("Now you have %g problems.\n", math.Sqrt(7))
    
    fmt.Println("pi =", math.Pi)
    
    fmt.Println("1 + 1 =",add(1,1))
    
    a,b := swap("hello", "world")
    fmt.Println(a,b)
    
    fmt.Println(split(19))
    
    var i int
    fmt.Println(i, c, python, java)
    
    var c, python, java = false, false, "no!"
    fmt.Println(c,python,java)

    c2, python2, java2 := false, false, "no!"
    fmt.Println(c2,python2,java2)
    
    var (
        ToBe    bool        = false
        MaxInt  uint64      = 1<<64-1
        z       complex128  = cmplx.Sqrt(-5 + 12i)
    )
    fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
    fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
    fmt.Printf("Type: %T Value: %v\n", z, z)
    
    var ii int
    var ff float64
    var bb bool
    var ss string
    fmt.Printf("%v %v %v %q\n", ii, ff, bb, ss)
    
    var x, y = 3, 4
    var f float64 = math.Sqrt(float64(x*x + y*y))
    var zz uint = uint(f)
    fmt.Println(x, y, zz)
}
