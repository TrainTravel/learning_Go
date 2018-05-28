package main

import (
        "fmt"
        "time"
       )

// WORKERCOUNT is number of workers
var WORKERCOUNT int

func worker(jobs chan int, done chan bool, workerID int) {
    for elem := range jobs {
        fmt.Println("Worker", workerID, "Consuming:", elem)
            time.Sleep(3 * time.Second)
    }
    done <- true
}

func main() {
        WORKERCOUNT = 5
        jobs := make(chan int, 5)
        done := make(chan bool)

        //Fill Queue
        go func() {
            for i := 0; i < 10; i++ {
                fmt.Println("Producing: ", i)
                    jobs <- i
                    time.Sleep(1 * time.Second)
            }
            defer close(jobs)
        }()

    // Spawn Workers
    for j := 0; j < WORKERCOUNT; j++ {
        go worker(jobs, done, j)
    }
    <-done
}
