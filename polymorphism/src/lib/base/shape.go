package baselib
import "fmt"

type Render interface {
    Draw2D()
    Draw3D()
}

type Shape struct {
    Height int
    Width int
}

type Circle struct {
    Shape  // anonymous field. You now have everything in Shape, so it's like inheritance
}

func (a *Shape) Draw3D() {
    fmt.Println("Shape (height:%d, width:%d) calls Draw3D:",a.Height,a.Width)
}

func (a *Shape) Draw2D() {
    fmt.Println("Shape (height:%d, width:%d) calls Draw2D:",a.Height,a.Width)
}

