package main

import ("fmt")

func main(){
    l := LinkedList{}
	l.show()
    n3 := Node{value:3}
    l.add(&n3)
    l.add(&Node{value:2})
    l.add(&Node{value:1})
    //~ l.show()
    found := l.search(3)
    fmt.Printf("%d %d \n",found.value, n3.value)
    fmt.Printf("%p %p \n",found, &n3)
}
