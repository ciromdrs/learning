package data_structures

import "fmt"

type LinkedList struct {
    Length int
	First *Node
}

type Node struct {
	Value int
	Next *Node
}

func NewNode(value int) *Node {
	n := &Node{Value : value}
	return n
}

func (l *LinkedList) Add(n *Node) {
    n.Next = l.First
    l.First = n
    l.Length++
}

func (l *LinkedList) Show(){
    for n := l.First; n != nil; n = n.Next {
        fmt.Print(n.Value)
        if n.Next != nil {
            fmt.Print(", ")
        }
    }
}


func (l *LinkedList) Search(value int) *Node {
    for n := l.First; n != nil; n = n.Next {
        if n.Value == value {
            return n
        }
    }
    return nil
}

func (l *LinkedList) Remove(value int) {
    var p *Node // previous node
    for n := l.First; n != nil; n = n.Next {
        if n.Value == value {
            switch n {
            case l.First:
                l.First = n.Next
            default:
                p.Next = n.Next
            }
            // Comment next line to remove all elements with this value
            //break 
        }
        p = n
    }
}

func (l *LinkedList) Update(old, new int){
    for n := l.First; n != nil; n = n.Next {
        if n.Value == old {
            n.Value = new
        }
    }
}
