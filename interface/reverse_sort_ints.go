package main

import (
    "sort"
    "fmt"
)

//func (s []string) Len() int {
//    return len(s)
//}
//
//func (s []string) Less(i, j int) bool {
//    return s[i] < s[j]
//}
//
//func (s []string) Swap(i, j int) {
//    s[i], s[j] = s[j], s[i]
//}

func main () {
    n := []int{7,4,8,2,9,19,12,32,3}
    sort.Sort(sort.Reverse(sort.IntSlice(n)))
    //sort.Ints(n)
    //fmt.Println("sort.Ints(n)", n)
    fmt.Println(n)

}
