package main

import (
    // "fmt"
    "data_structures"
)

func main(){
    // l := data_structures.LinkedList{}
    // n3 := data_structures.Node{Value:3}
    // nn3 := data_structures.Node{Value:3}
    // l.Add(&n3)
    // l.Add(&data_structures.Node{Value:2})
    // l.Add(&data_structures.Node{Value:3})
    // l.Add(&data_structures.Node{Value:1})
    // l.Add(&nn3)
    // l.Show()
    // l.Update(3,5)
    // fmt.Println()
    // l.Show()


    s := data_structures.NewSet(2)
    _ = s
    s.Add(data_structures.NewNode(5))
    s.Add(data_structures.NewNode(7))
    s.Show()
    // s.Add(e)
    // fmt.Println(e.Value)
}
