package main

import "fmt"

type LinkedList struct {
    length int
	first *Node
}

type Node struct {
	value int
	next *Node
}

func newNode(value int) Node {
	n := Node{value : value}
	return n
}

func (l *LinkedList) add(n *Node) {
    n.next = l.first
    l.first = n
    l.length++
}

func (l *LinkedList) show(){
    for n := l.first; n != nil; n = n.next {
        fmt.Print(n.value)
        if n.next != nil {
            fmt.Print(", ")
        }
    }
}


func (l *LinkedList) search(value int) *Node {
    for n := l.first; n != nil; n = n.next {
        if n.value == value {
            return n
        }
    }
    return nil
}
