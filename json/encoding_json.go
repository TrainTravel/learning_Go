package main

import (
    "encoding/json"
    "os"
)

type Person struct {
    First       string
    Last        string
    Age         int
    notExported int
}


func main () {
    p1 := Person {"James", "Bond", 25, 007}

    // .NewEncoder takes an object of type Writer, which implements the .write method( os.Stdout is the Writer )
    //         and returns a pointer pointing to an object of type Encoder, which has .encode method
    // Here, writing to Stdout(open file descriptors which we always can wrtie to)
    // Stdout = NewFile(******)  //where the type NewFile returns a pointer of File type (*File)
    // so with os.Stdout we now have a File on which ".write()" can be called
    // type Writer is an interface which should have a method
    // Since the File type is implicitly implementing the Writer interface, IT IS A "Writer" INTERFACE!
    json.NewEncoder(os.Stdout).Encode(p1)
}
