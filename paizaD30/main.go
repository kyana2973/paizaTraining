package main

import (
	"fmt"
	"strconv"
)

func main() {
	var numArr []string
	//スライスに九九の値を格納していく
	for i := 1; i < 10; i++ {
		for j := 1; j < 10; j++ {
			v := i * j
			s := strconv.Itoa(v)
			numArr = append(numArr, s)
		}
	}

	//改行時の位置を決める
	over := len(numArr) % 9

	//1桁の数字は2桁になるようスペース埋めして出力する
	for i := 0; i < len(numArr); i++ {
		if (i+1)%9 == over {
			//改行する
			fmt.Printf("%2s\n", numArr[i])
		} else {
			//数字との間は、「 | 」となるようにする
			fmt.Printf("%5s", numArr[i]+" | ")
		}

		//最初と最後には、区切り線を入れない
		if (i+1)%9 == over && (i+1) != len(numArr) {
			fmt.Println("==========================================")
		}
	}
}
