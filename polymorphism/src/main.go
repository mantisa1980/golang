package main
import ("lib/base"  // import path is $GOPATH + import path. import only to the folder that contains the package.
        "fmt"
        //"gopkg.in/mgo.v2"
        ) // fmt is located at GOROOT path so you can find it this way.

func main() { // executable programs must be in package main
    var i baselib.AnimalBehavior
    //i = baselib.Animal{Height:100,Name:"Horse"}  // baselib.Animal does not implement baselib.AnimalBehavior; only *baselib.Animal does.
    i = &baselib.Animal{Height:100,Name:"Horse"}

    i.MakeSound()
    i.Move()
    i.ChangeName("Pig")
    fmt.Println(i.GetName())

    var j baselib.Render
    j = &baselib.Shape{3,4}
    j.Draw2D()
    j.Draw3D()

    j = &baselib.Circle{baselib.Shape{9,16}}
    j.Draw2D()
    j.Draw3D()

}
