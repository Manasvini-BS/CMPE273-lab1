package sleep
/*
* @Author : Manasvini Banavara Suryanarayana
* @SJSU ID : 010102040
*/

import "testing"
import "time"


func TestSleep(t *testing.T){
	a := time.Now()
	Sleep(5)
	b := time.Since(a)
	if b.Seconds() < 5{
		t.Error("Expected sleep time was 5, but system slepts only for : ",b.Seconds())
	}
}