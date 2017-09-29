package main

import (
    "fmt"
    )


func main() {
    var a map[int] int = make(map[int]int)
    a[1] = 2
    for i,j:= range a {
        fmt.Println(i,j);
    }
    x, y := a[1]
    fmt.Println(x,y);
    if y {
        fmt.Println("true")
    } else {
        fmt.Println("false")
    }
    x, y = a[2]
    fmt.Println(x,y);
    if y {
       fmt.Println("true")
    } else {
        fmt.Println("false");
    }
}

