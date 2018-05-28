package main

import (
        "sync"
        "time"
        "math/rand"
        "fmt"
        "sync/atomic"
       )

var wg sync.WaitGroup
// var counter int
var counter int64

func main() {
    wg.Add(2)
    go incrementor("foo")
    go incrementor("bar")
    wg.Wait()
    fmt.Println("Final Counter:", counter)
}

func incrementor(s string) {
    for i:=0; i<20; i++ {
        time.Sleep( time.Duration(rand.Intn(3)) * time.Millisecond )
        // no need to pack with mutex lock thing
        // mutex.Lock()
        // counter++

        //fmt.Println(&counter)
        atomic.AddInt64(&counter, 1)
        fmt.Println(s, i, "Counter:", counter)
        // mutex.Unlock()
        
    }
    wg.Done()
}
