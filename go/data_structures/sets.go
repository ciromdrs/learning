package data_structures

import "fmt"

type Set struct {
    Elements []LinkedList
    Size int
}

func NewSet(capacity int) *Set {
    s := &Set{}
    s.Elements = make([]LinkedList,capacity,capacity)
    return s
}

func (s *Set) Show() {
    for i, l := range s.Elements {
        fmt.Printf("|%d|-",i)
        for n := l.First; n != nil; n = n.Next {
            fmt.Printf(">[%d]-", n.Value)
        }
        fmt.Println("|")
    }
}

func (s *Set) Add(e *Node) {
    if s.Size < cap(s.Elements) {
        index := e.Value % cap(s.Elements)
        s.Elements[index].Add(e)
        s.Size++
    } else {
        // TODO: re-hash
    }
}

/*func (l *LinkedList) Show(){
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
}*/