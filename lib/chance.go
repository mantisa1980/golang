package lib2  // package name has not relationship with its path name
import "fmt"

func init(){
	fmt.Println("init in chance.go")
}

func Foo(x int) int { // public function should begin with uppercase
	return x
}
