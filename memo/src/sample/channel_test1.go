package main
import "fmt"
func main() {
    messages := make(chan string)
    go func() { 
        messages <- "ping"
        fmt.Println("put ping")
    }()
    fmt.Println("main pulling msg")
    msg := <-messages
    fmt.Println(msg)
    fmt.Println("done")
}

