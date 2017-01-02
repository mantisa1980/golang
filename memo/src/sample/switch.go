package main

import (
    "fmt"
    "runtime"
)

func switch1() {
    fmt.Print("Go runs on ")
    switch os := runtime.GOOS; os {  // assign os var , then switch os var
    case "darwin":
        fmt.Println("OS X.")
    case "linux":
        fmt.Println("Linux.")
    default:
        // freebsd, openbsd,
        // plan9, windows...
        fmt.Printf("%s.", os)
    }
}

func switch2() {  // switch with no condition : means switch true; can be a shorter way to implement long if/else statement
    t := time.Now()
    switch {  
    case t.Hour() < 12:
        fmt.Println("Good morning!")
    case t.Hour() < 17:
        fmt.Println("Good afternoon.")
    default:
        fmt.Println("Good evening.")
    }
}

func main() {
    switch1()
}

