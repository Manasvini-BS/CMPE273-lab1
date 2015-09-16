package shape

/*
* @Author : Manasvini Banavara Suryanarayana
* @SJSU ID : 010102040
*/
import "math"

type Shape interface {
	area() float64
	perimeter() float64
}

type Rectangle struct {
     x1, y1, x2, y2 float64
}

type Circle struct {
     x, y, r float64
}

func (c *Circle) area() float64 {
     return math.Pi * c.r*c.r
}
func (c *Circle) perimeter() float64{
	var a float64 = (2 * math.Pi * c.r)
	return a
}
func distance(x1, y1, x2, y2 float64) float64 {
var a float64 = (x2 - x1)
var b float64 = (y2 - y1)
var c float64 = math.Sqrt(a*a + b*b)
return c

}

func (r *Rectangle) area() float64 {
     l := distance(r.x1, r.y1, r.x1, r.y2)
     w := distance(r.x1, r.y1, r.x2, r.y1)
     return l * w
}

func (r *Rectangle) perimeter() float64 {
	l := distance(r.x1, r.y1, r.x1, r.y2)
     w := distance(r.x1, r.y1, r.x2, r.y1)
     var a float64 = (2 * l) + (2 * w)
     return a
}
func measureArea(s Shape) float64 {
var a float64 = s.area()
return a

}
func measurePerimeter(s Shape) float64 {

var a float64 = s.perimeter()
return a

}