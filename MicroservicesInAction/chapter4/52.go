package main

import "fmt"

type IPrint interface {
    MyPrint()
}

type IS1 struct {
    A,B int
    S string
}

type IS2 struct {
    S string
}

func main() {
    var is1 IPrint
    s1 := IS1{A:1,B:1,S:"hello"}
    //is1.MyPrint()
    is1 = s1
    is1.MyPrint()
    //fmt.Println(is1.S)
    is1 = IS2{S:"hello world"}
    is1.MyPrint()
}

func (i IS1) MyPrint() {
    fmt.Println(i.S)
}

func (i IS2) MyPrint() {
    fmt.Println(i.S)
}

