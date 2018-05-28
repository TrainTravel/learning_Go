package main

import (
    "fmt"
)

func main(){
    c1 := incrementor("foo")
    c2 := incrementor("bar")
    c3 :=puller(c1)
    c4 :=puller(c2)

    // for blocking main to exit before things are done
    fmt.Println("Final counter", <-c3+<-c4)
}


func incrementor(s string) chan int {
    out := make(chan int)
    go func() {
        for i:=0; i<20; i++{
            out<-1
            fmt.Println(s, i)
        }
        close(out)
    }()
    return out
}

func puller(c chan int) chan int {
    out := make(chan int)
    go func() {
        var sum int
        for n := range c {
            sum += n
        }
        out<-sum
        close(out)
    }()
    return out
}


