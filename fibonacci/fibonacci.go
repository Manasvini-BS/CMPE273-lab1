package fibonacci
/*
* @Author : Manasvini Banavara Suryanarayana
* @SJSU ID : 010102040
*/
func fib(x int)int {
	if x==0 {
		return 0
	}
	if x==1 {
		return 1
	}

    var z int = fib(x-1)+fib(x-2)
    return z


}



