/*
Interfaces
An interface type is defined as a set of method signatures.

A value of interface type can hold any value that implements those methods.

Note: There is an error in the example code. Vertex (the value type) doesn't implement Abser 

because the Abs method is defined only on *Vertex (the pointer type).
*/

package main

import (
    "fmt"
    "math"
    "reflect"
)

type Abser interface {
    Abs() float64
}

type Abser2 interface { 
    Abser
    Abs2() float64
}

func (f MyFloat) Abs2() float64 {
    if f < 0 {
        return float64(-f)
    }
    return float64(f)
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
    if f < 0 {
        return float64(-f)
    }
    return float64(f)
}

type Vertex struct {
    X, Y float64
}

// do not need to specify Abser (the interface name). Will decouple def. of interface from implementation. 
// You can switch implementation between packages easily (think again)
func (v *Vertex) Abs() float64 { 
    return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// cannot redeclare ,so you can only implement either by receiving pointer (will noy copy) or by receiving copy
//func (v Vertex) Abs() float64 {
//    return math.Sqrt(v.X*v.X + v.Y*v.Y)
//}

func main() {

    var a Abser
    var b Abser2

    f := MyFloat(-math.Sqrt2)
    v := Vertex{3, 4}
    pv := &v
    fmt.Println(reflect.TypeOf(v), reflect.TypeOf(pv), pv.Abs())

    a = f  // a MyFloat implements Abser
    fmt.Printf("%v %T \n",a,a) // interface value: is actually a tuple of (value, type) : value of a specific underlying type
    a = &v // a *Vertex implements Abser

    b = f
    fmt.Println("abs2", b.Abs(), b.Abs2())

    // In the following line, v is a Vertex (not *Vertex)
    // and does NOT implement Abser.
    //a = v  // this will cause error

    fmt.Println(a.Abs())
    fmt.Printf("%v %T \n",a,a) // interface value: is actually a tuple of (value, type) : value of a specific underlying type
}

