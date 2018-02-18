package main

import ("fmt")

func main(){
    l := LinkedList{}
    n3 := Node{Value:3}
    nn3 := Node{Value:3}
    l.Add(&n3)
    l.Add(&Node{Value:2})
    l.Add(&Node{Value:3})
    l.Add(&Node{Value:1})
    l.Add(&nn3)
    l.Show()
    l.Remove(3)
    fmt.Println()
    l.Show()

}
