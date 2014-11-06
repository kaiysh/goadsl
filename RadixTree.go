/*
    Radix Tree
*/

package main

import "fmt"

type Node struct {
    Key []byte
    Value int
    Link *Node
    Next *Node
}

func NewNode() *Node {
    return &Node{}
}

func (n *Node) Prefix(b []byte) int {
    for i := 0; i < len(b); i++ {
        if i == len(n.Key) || n.Key[i] != b[i] {
            return i
        }
    }
    return len(b)
}

func (n *Node) Split(k int) {
    newnode := NewNode()
    newnode.Key = n.Key[k:]
    newnode.Link = n.Link
    newnode.Value = n.Value
    n.Value = 0
    n.Link = newnode
    n.Key = n.Key[:k]
}

func (n *Node) Merge() {
    link := n.Link
    n.Key = append(n.Key,link.Key...)
    n.Value = link.Value
    n.Link = link.Link
}

func (n *Node) search(b []byte) *Node {
    k := n.Prefix(b)
    if k == 0 {
        if n.Next == nil { return nil }
        return n.Next.search(b)
    } else if k == len(b) {
        return n
    } else if k == len(n.Key) {
        if n.Link == nil { return nil }
        return n.Link.search(b[k:])
    }
    return nil
}

func (n *Node) insert(b []byte, i int) *Node {
    k := n.Prefix(b)
    if k == 0 {
        if n.Next == nil {
            n.Next = NewNode()
            n.Next.Key = b
            n.Next.Value = i
        } else {
            n.Next = n.Next.insert(b,i)
        }
    } else if k < len(b) {
        if k < len(n.Key) {
            n.Split(k)
        }
        if n.Link == nil {
            n.Link = NewNode()
            n.Link.Key = b[k:]
            n.Link.Value = i
        } else {
            n.Link = n.Link.insert(b[k:],i)
        }
    } else {
        n.Value = i
    }
    return n
}

func (n *Node) delete(b []byte) *Node {
    k := n.Prefix(b)
    if k == 0 {
        if n.Next != nil {
            n.Next = n.Next.delete(b)
        }
    } else if k == len(b) {
        return n.Next
    } else if k == len(n.Key) {
        if n.Link != nil {
            n.Link = n.Link.delete(b[k:])
            if (n.Link != nil) && (n.Link.Next == nil) {
                n.Merge()
            }
        }
    }
    return n
}

type RadixTree struct {
    Root *Node
}

func NewRadixTree() *RadixTree {
    return &RadixTree{NewNode()}
}

func (r *RadixTree) Search(s string) (int, bool) {
    n := r.Root.search(append([]byte(s),0))
    if n == nil { return 0, false }
    return n.Value, true
}

func (r *RadixTree) Insert(s string, i int) {
    r.Root.insert(append([]byte(s),0),i)
}

func (r *RadixTree) Delete(s string) {
    r.Root.delete(append([]byte(s),0))
}

func main() {
    r := NewRadixTree()
    r.Insert("hello",10)
    r.Insert("hel",5)
    r.Insert("helloworld",100)
    fmt.Print("Searching for 'hel': ")
    fmt.Println(r.Search("hel"))
    fmt.Print("Searching for 'hello': ")
    fmt.Println(r.Search("hello"))
    fmt.Print("Searching for 'heloworld': ")
    fmt.Println(r.Search("helloworld"))
    r.Delete("hello")
    fmt.Print("Searching for 'hel': ")
    fmt.Println(r.Search("hel"))
    fmt.Print("Searching for 'hello': ")
    fmt.Println(r.Search("hello"))
    fmt.Print("Searching for 'heloworld': ")
    fmt.Println(r.Search("helloworld"))
}
