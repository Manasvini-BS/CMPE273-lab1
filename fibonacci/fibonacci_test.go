package fibonacci
/*
* @Author : Manasvini Banavara Suryanarayana
* @SJSU ID : 010102040
*/

import "testing"

import "fmt"

func TestFibonacci(t *testing.T){
	var x int
	x = fib(5)
	fmt.Println(x)
	if x != 5 {
		t.Error("Expected output is 5, but got ", x)
	}
}