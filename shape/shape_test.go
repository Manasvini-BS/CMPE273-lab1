package shape
/*
* @Author : Manasvini Banavara Suryanarayana
* @SJSU ID : 010102040
*/
import "testing"

func TestShape(t *testing.T){

c := Circle{0, 0, 5}
r := Rectangle{0,0,5,4}
 var y float64 = measureArea(&c)
 var x float64 = measureArea(&r)
 var y1 float64 = measurePerimeter(&c)
 var x1 float64 = measurePerimeter(&r)

	if y != 78.53981633974483 {
		t.Error("area of circle is mismatching with actual value, value we got is ", y)
	}
		if x != 20 {
		t.Error("area of rectangle is mismatching with actual value, value we got is", x)
	}
		if y1 != 31.41592653589793 {
		t.Error("perimeter of circle is mismatching with actual value, value we got is", y1)
	}
		if x1 != 18 {
		t.Error("perimeter of rectangle is mismatching with actual value, value we got is", x1)
	}
}