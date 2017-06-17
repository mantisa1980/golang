/*
he empty interface
The interface type that specifies zero methods is known as the empty interface:

interface{}
An empty interface may hold values of any type. (Every type implements at least zero methods.)

Empty interfaces are used by code that handles values of unknown type. For example, fmt.Print takes any number of arguments of type interface{}.

*/

package main

import ("time"; "fmt")

func main() {
    c1 := make(chan int)
    c2 := make(chan int)

    go func() {
        fmt.Println("G1 begins")
        for i := range c1{
            println("G1 got", i)
            c2 <- i
            println("G1 sent", i)
        }
        fmt.Println("G1 ends")
    }()

    go func() {
        fmt.Println("G2 begins")
        for i := range c2 {
            println("G2 got", i)
            c1 <- i
            println("G2 sent", i)
        }
        fmt.Println("G2 ends")
    }()
    fmt.Println("main sending 1")
    c1 <- 1
    fmt.Println("main sent 1")
    c1 <- 2
    fmt.Println("main sent 2")
    time.Sleep(1000000000 * 50)
    //time.Sleep(5000000000)
    //var aaa = <- c1
    //fmt.Println(aaa)
    fmt.Println("main done")
}

