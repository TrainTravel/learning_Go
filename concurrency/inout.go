package main

import (
        "io"
        "os"
        "time"
       )

func main() {
    go echo(os.Stdin, os.Stdout)
        // Blocking code. Makes main routine to wait for 5 seconds
        // time.Sleep(5 * time.Second)
        <-time.After(5 * time.Second)
}

func echo(out io.Writer, in io.Reader) {
    io.Copy(out, in)
}
