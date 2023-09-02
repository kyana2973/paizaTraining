// 問題文URL:https://paiza.jp/works/mondai/conditions_branch/conditions_branch__mod_boss

package main

import (
	"fmt"
	"strconv"
)

func main() {
	var strN string
	fmt.Scanf("%s", &strN)
	intN, _ := strconv.Atoi(strN)
	//3と5で割り切れる場合は"FizzBuzz"、3でしか割り切れない場合は"Fizz"、5でしか割り切れない場合は"Buzz"となる
	//どれも当てはまらない場合は、intNをそのまま出力する
	if intN%3 == 0 && intN%5 == 0 {
		fmt.Println("FizzBuzz")
	} else if intN%5 == 0 {
		fmt.Println("Buzz")
	} else if intN%3 == 0 {
		fmt.Println("Fizz")
	} else {
		fmt.Println(intN)
	}
}
