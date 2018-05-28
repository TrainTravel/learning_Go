package main

import (
    "net/http"
    "io/ioutil"
    "fmt"
    //"os"
)

var url string = "http://apilayer.net/api/validate?access_key=c371e93454a5ec53d531ffdabd209058&number=0933905654&country_code=TW&format=1"

func main() {
    //res, _ := http.Get("http://golang.org/")
    res, _ := http.Get(url)
    page, _ := ioutil.ReadAll(res.Body)
    fmt.Println("%s", string(page))
    //res, err := http.Get("http://golang.org/")
    //if err != nil {
    //    fmt.Printf("%s", err)
    //    os.Exit(1)
    //} else {
    //    page, err := ioutil.ReadAll(res.Body)
    //    if err != nil{ 
    //        fmt.Printf("%s", err)
    //        os.Exit(1)
    //    }else {
    //        res.Body.Close()
    //    }
    //    text = string(res.page)
    //    fmt.Println("%s", text )
    //}
}
