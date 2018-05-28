package main

import (
        "sync"
        "time"
        "math/rand"
        "fmt"
       )

var wg sync.WaitGroup
var counter int
var mutex sync.Mutex

func main() {
    wg.Add(2)
    go incrementor("foo")
    go incrementor("bar")
    wg.Wait()
    fmt.Println("Final Counter:", counter)
}

func incrementor(s string) {
    for i:=0; i<20; i++ {
        time.Sleep( time.Duration(rand.Intn(20)) * time.Millisecond )
        mutex.Lock()
        counter++
        fmt.Println(s, i, "Counter:", counter)
        mutex.Unlock()
        
    }
    wg.Done()
}
