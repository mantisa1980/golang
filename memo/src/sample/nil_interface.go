/*
注意接收pointer的function要檢查nil pointer, 不然呼叫會tuntime error , segmentation violation
*/

package main

import "fmt"

type I interface {
    M()
}

type T struct {
    S string
}

func (t *T) M() {
    //fmt.Println(t.S) will trigger error for nil pointer 
    if t == nil {
        fmt.Println("<nil>")
        return
    }
    fmt.Println(t.S)
}

func main() {

    var i I

    var t *T  // default initalized as nil
    if t == nil {
        fmt.Println("ptr init as nil")
    }
    i = t
    describe(i)
    i.M()

    i = &T{"hello"}
    describe(i)
    i.M()


}

func describe(i I) {
    fmt.Printf("(%v, %T)\n", i, i)
}
