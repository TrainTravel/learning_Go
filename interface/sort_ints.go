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
    N := sort.IntSlice{7,4,8,2,9,19,12,32,3}
    n := []int{7,4,8,2,9,19,12,32,3}
    N.Sort()
    sort.Ints(n)
    fmt.Println("sort.Ints(n)", n)
    fmt.Println("IntSlice.Sort()", N)

}
