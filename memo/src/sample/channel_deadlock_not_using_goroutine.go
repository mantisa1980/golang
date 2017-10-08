package main
import "fmt"

/*
By default channels are unbuffered, meaning that they will only accept sends (chan <-) if there is a corresponding receive (<- chan) ready to receive the sent value.  (from [Go Example](https://gobyexample.com/channel-buffering))
*/
// 下面流程是send會先執行 但沒有對應的人在等接收, 而且是同一個goroutine切不過去 

func main() {
    c := make(chan int)  
    //needRoutine := true
    needRoutine := false
    //並不一定會進入go routine來計算，也就是說channel會是為空的．
    if needRoutine == true {
        go func() {
            c <- 1
        }();
    } else {
        c <- 0 //note this kind will not work, because channels need input value in go routine.
    }

    fmt.Printf("Not goes here")

    // 這裏將會變成deadlock，而無窮的等待．
    totalResult := <-c * 3 + 5
    fmt.Printf("result is %d", totalResult)
}
