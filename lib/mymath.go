//package lib1  // cannot have two package names in a single folder. all files must be the same package
package lib2
import "fmt"

func init(){
	fmt.Println("init in mymath.go")
}

func Bar(x int) int { // public function should begin with uppercase
	return x
}
