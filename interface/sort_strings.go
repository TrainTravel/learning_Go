package main

import (
    "sort"
    "fmt"
)

// 3 ways of sorting strings
func main () {
    s := []string{"Zeno", "John", "Al", "Jenny"}
    //S := sort.StringSlice(s)
    //S.sort()

    // since the type StringSlice has a method called Sort defined
    sort.StringSlice(s).Sort()
    fmt.Println(s)
    
    // sort.Sort accepts some "data" that implements the sort.Interface interface
    // Since sort.StringSlice has a method set which is a superset of {"Len", "Less", "Swap"}, it's a sort.Interface interface!
    sort.Sort(sort.StringSlice(s))
    fmt.Println(s)

    sort.Strings(s)
    fmt.Println(s)

    // sort.Reverse takes sth that implements the sort.Interface interface
    // and Reverse returns an interface
    sort.Sort(sort.Reverse(sort.StringSlice(s)))
    fmt.Println(s)
    //fmt.Println(S)

}
