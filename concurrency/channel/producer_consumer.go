package main

import (
        "fmt"
        "time"
       )

func main() {
                   isProducerDone := make(chan bool)
                   buffer := make(chan int)

                    // Start a producer
                    go func() {
                        for i := 0; i < 5; i++ {
                            fmt.Println("Producing..", i)
                                buffer <- i
                                time.Sleep(5 * time.Second)
                        }
                        isProducerDone <- true
                    }()

                // Start a consumer
                go func() {
                    for product := range buffer {
                        fmt.Println("Consuming..", product)
                    }
                }()

                <-isProducerDone
}
