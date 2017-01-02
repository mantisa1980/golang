package main

import "fmt"
/*
A closure is a function value that references variables from outside its body.
The function may access and assign to the referenced variables; in this sense the function is "bound" to the variables.
*/

// grouped const declaration
const (
    // Create a huge number by shifting a 1 bit left 100 places.
    // In other words, the binary number that is 1 followed by 100 zeroes.
    Big = 1 << 100
    // Shift it right again 99 places, so we end up with 1<<1, or 2.
    Small = Big >> 99
)

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
    var fn_1 int = 1
    var fn_2 int = 0
    return func() int {  // return a closure. fn_1 and fn_2 is referenced by this closure so states are reserved, until the closure no logner referenced
        var fn_2_backup int = fn_2
        fn_2 = fn_1
        fn_1 = fn_1 + fn_2_backup
        return fn_1 
    }
}

func main() {
    f := fibonacci()
    for i := 0; i < 5; i++ {
        fmt.Printf("%d ", f())
    }
    fmt.Println()
    
    // state is cleared once get a new closure
    f = fibonacci()
    for i := 0; i < 5; i++ {
        fmt.Printf("%d ", f())
    }
}
