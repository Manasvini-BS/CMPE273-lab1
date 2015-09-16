package sleep
/*
* @Author : Manasvini Banavara Suryanarayana
* @SJSU ID : 010102040
*/
import "time"

func Sleep(sec int) {
    <-time.After(time.Second * time.Duration(sec))
}