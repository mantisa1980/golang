package main

import (
    "fmt"
    "reflect")


func main() {
    //test_reference()
    //test_ptr()
    //test_struct()
    //test_array()
    //test_types()
    //test_slice_capacity()
    //test_copy_on_capacity_overload()
    test_looping()
}

func test_types() {
    var a = [1]string{"Hello"}  // this defines a type!
    var b = [1]string{"World"}  
    var c = [2]string{}  // this defines another different type!

    fmt.Println(reflect.TypeOf(a), reflect.TypeOf(b), reflect.TypeOf(c))
    fmt.Println(reflect.TypeOf(a)==reflect.TypeOf(b), reflect.TypeOf(b)==reflect.TypeOf(c))
}

func test_reference() {
    type test struct {
        a int
    }

    b := &test{2}
    c := &test{2}
    fmt.Println(c == b)
    b = c
    fmt.Println(c == b, *c ==*b)  // c==b is equal to python is(reference comparison), whereas *c==*b is equal to python == (referred object's comparison)


}

func test_ptr() { 
    i, j := 1, 2
    p := &i         // point to i
    fmt.Println(*p) // read i through the pointer
    *p = 21         // set i through the pointer
    p = &j
    fmt.Println(*p)
    
    {
        k:=3
        p=&k
        fmt.Println(p ,*p)
    }
    fmt.Println(p,*p)
}

func test_struct() { 
    type Vertex struct {
        X int
        Y int
    }

    a,b := Vertex{X:1, Y:2}, Vertex{X:3, Y:4}
    p1 := &Vertex{X:5}

    // w := Vertex(3,4)   // this is error. () for function call, {} for initialization 
    a.Y = 0
    p2 := &a
    fmt.Println(a,b,*p1)

    x := 0
    //c=&x  // c is already declared (initialized as *Vertex), cannot cast to int
    fmt.Println(x,*p2)
}

func modify_passed_array(a [2]string){  // array is copied; [2]string is a specific type; [3]string is another different type; cannot be casted !
    a[0] = "Modified in array"
    fmt.Println("a in modify_passed_array:", a)
}

func modify_passed_slice(a []string){  // pass by slice reference; array is not copied; []string: slice type
    a[0] = "Modified in slice"
    fmt.Println("a in modify_passed_alice:", a)
}

func modify_passed_pointer(a *[2]string){
    
    (*a)[0] = "Modified in pointer"
    fmt.Println("a in modify_passed_pointer:", a, *a)
}

/// slice creates a view of a subset of array.
func test_array() {
    // var b [2]string{"Hello", "World"}  // error


    var x = [2]string{ "Hello","Hello"}
    fmt.Println(x[0]==x[1])  // string compare

    var a = [2]string{"Hello", "World"}
    fmt.Println("Before a=", a)

    modify_passed_array(a)  // arrays are like structs, with key-values as 0:value, 1:value ,2:value rather than struct key1:value, key2:value2
    // so passing array will make a COPY of it , just like struct ! To pass oroginal array, pass a pointer to the array
    fmt.Println("After passing to modify array func ,a=",a)

    modify_passed_slice(a[0:]) // pass a slice that wraps underlying array! so this will change original array!
    fmt.Println("After passing to modify slice func ,a=",a)

    modify_passed_slice(a[1:]) // pass a slice that wraps underlying array! so this will change original array!
    fmt.Println("After passing to modify slice func ,a=",a)

    //p1 := &a
    //fmt.Println((*p1)[0], (*p1)[1])
    //fmt.Println(reflect.TypeOf(p1))

    modify_passed_pointer(&a)
    fmt.Println("After passing to modify pointer func ,a=",a)

    //modify_passed_pointer(a[])
    //fmt.Println("After passing to modify pointer func ,a=",a)
}

func test_slice_capacity() {
    s := []int{2, 3, 5, 7, 11, 13}
    fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)

    // reslice the slice to give it zero length.
    s = s[:0]
    fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)

    // Extend its length.
    s = s[:4]
    fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)

    // Drop its first two values.
    s = s[:7]
    fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

// slice implementation differs from python slicing (copy on write)
func test_copy_on_capacity_overload() {

    /*
    a :=[5]int{0,1,2,3,4}

    [0  1  2  3  4]
    a-------------
    b      ----
    
    b = append(b,5)

    [0  1  2  3  5]
    a-------------
    b      -------           

    b = append(b,6)
    [0  1  2  3  5]
    a-------------
                     [2  3  5  6]
                     b----------

    */

    a :=[5]int{0,1,2,3,4}
    b:=a[2:4]
    fmt.Println(cap(a), cap(a[0:1]), cap(a[0:2]),cap(a[2:4]), cap(a[3:4]), cap(a[4:4])   )  // capacity: head to end of actual array 

    b = append(b,5)  // this affacts a also, since it still has enough capacity
    fmt.Println(a, ",", b)
    
    b = append(b,6) // since capacity is exceeded, b is moved to another copy of oroginal underlying array, so now a, b use separate memory
    fmt.Println(a, ",", b)

}


func test_looping() {
    sum:=0
    for i:=0;i<10;i++ {
        sum+=i
    }
    fmt.Println(sum)
    // for without any ; is while loop in GO
    sum=0
    for sum<10 {
        sum+=1
    }

    fmt.Println(sum)

    // for {}   this is infinite loop
    
    sum = 0
    for {
        sum+=1
        if sum > 10 {
            break
        } else {  // must be placed at the same line after } of if
            continue
        }
    }
    
    fmt.Println(sum)

}