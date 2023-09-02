// 問題文URL:https://paiza.jp/works/mondai/conditions_branch/conditions_branch__mod_step4

package main

import (
	"fmt"
	"strconv"
)

func main() {
	var strN string
	fmt.Scanf("%s", &strN)
	//intNには、1~31日までのどれかの日が入力されます
	intN, _ := strconv.Atoi(strN)
	weeks := []string{"Sat", "Sun", "Mon", "Tue", "Wed", "Thu", "Fri"}
	//1日が日曜日の場合で、intN日が何曜日が出力する
	fmt.Println(weeks[intN%7])
}
