//package lib1  // cannot have two package names in a single folder. all files must be the same package
package baselib
import "fmt"

type AnimalBehavior interface {
    Move()
    MakeSound()
    ChangeName(s string)
    GetName() string
}

type Animal struct {
    Height int
    Name string
}

func (a *Animal) Move() {
    fmt.Println("Animal %s (height:%d) calls implements Move:",a.Name,a.Height)
}

func (a * Animal) MakeSound() {
    fmt.Println("Animal %s (height:%d) calls implements MakeSound:",a.Name,a.Height)
}
func (a * Animal) ChangeName(s string) {
    a.Name = s
    //fmt.Println("Animal %s (height:%d) calls implements ChangeName:",a.Name,a.Height)
}
func (a * Animal) GetName() string {
    return a.Name
}
