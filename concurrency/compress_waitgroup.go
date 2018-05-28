package main

import (
        "compress/gzip"
        "fmt"
        "io"
        "os"
        "sync"
       )

// Create a WaitGroup
var wg sync.WaitGroup

func main() {
    for _, file := range os.Args[1:] {
        // Add workers to waitgroup
        wg.Add(1)
            go func(filename string) {
                fmt.Println(compress(filename))
                    // Mark task as done
                    wg.Done()
            }(file)
    }
    // On main thread, wait for others to complete
    wg.Wait()
}

func compress(file string) error {
    in, err := os.Open(file)
        if err != nil {
            return err
        }
    defer in.Close()
        // Create new file
        out, err := os.Create(file + ".tar.gz")
        if err != nil {
            return err
        }
    defer out.Close()
        // Create compressed stream
        gzout := gzip.NewWriter(out)
        // Copy from the input to this compressed file
        _, err = io.Copy(gzout, in)
        gzout.Close()
        return err
}
