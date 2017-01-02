package main

/*
Go does not have classes. However, you can define methods on types.

A method is a function with a special receiver argument.

The receiver appears in its own argument list between the func keyword and the method name.

In this example, the Abs method has a receiver of type Vertex named v.
*/

import (
    "fmt"
    "math"
)

type Vertex struct {
    X, y float64   // y lowercase can be accessed; Capital->export rule is for packages, not struct data.
}

// a normal function that receive Vertex as parameter
func Abs(v Vertex) float64 {
    return math.Sqrt(v.X*v.X + v.y*v.y)
    //return x+y
}

// a method defined on Vertex: there is a special receiver (can have only one, of course) between func declaration and function name.
// GO philosophy: better composition than inheritance
func (v Vertex) Abs() float64 {
    return math.Sqrt(v.X*v.X + v.y*v.y)
}

func (v Vertex) ScaleCopy(times float64) {
    v.X*=times
    v.y*=times
}

func (v* Vertex) ScalePtr(times float64) {
    v.X*=times
    v.y*=times
}

func main() {
    v := Vertex{3, 4}
    fmt.Println(v.Abs())  // call method of v
    fmt.Println(Abs(v))   // call a normal function
    fmt.Println("--------")

    v.ScaleCopy(10) // does not change v
    fmt.Println(v)
    p:=&v
    p.ScaleCopy(10)  // interpreted as (*p).ScaleCopy() automatically
    fmt.Println(v)

    fmt.Println("--------")

    v.ScalePtr(10) // will change v, just like normal function receiving pointers. GO interpretted automatically as (&v).ScalePtr(10)
    fmt.Println(v)

    /*
    ======= since ScalePtr takes a pointer receiver, no matter you call v.ScalePtr() or (&v).ScalePtr(), 
    they are the same. v.ScalePtr will be interpretted as (&v).ScalePtr() automatically. Pointer of v will be taken automatically for convenience =========

    */

    p=&v
    p.ScalePtr(10)
    fmt.Println(v)

    (&v).ScalePtr(10)
    fmt.Println(v)


       
}
