package main
import ("golang/lib"  // import path is $GOPATH + import path
		"fmt") // fmt is located at GOROOT path so you can find it this way.

func main() { // executable programs must be in package main
	fmt.Println("hello world")
	fmt.Println(lib2.Foo(123))
	fmt.Println(lib2.Foo(456))

}
