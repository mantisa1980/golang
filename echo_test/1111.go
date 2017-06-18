package main

import (
    "fmt"
    "math/rand"
    "time"
    )


func main() {
    
    fmt.Println(rand.Perm(10))
    sec := time.Now().UTC()
    fmt.Println(sec)
    //fmt.Println(sec.Format("20060102150405"))

    timestamp := time.Now().Unix()
    fmt.Println(timestamp)

    //rand.Seed(timestamp)
    fmt.Println(rand.Perm(10)[:3] )
    fmt.Println(rand.Perm(10)[:3] )

    //return Time{sec + unixToInternal, nsec, Local}
}


//func Perm(n int) []timestamp

