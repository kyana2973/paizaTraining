// 問題文URL:https://paiza.jp/works/mondai/array_utilization_primer/array_utilization_primer__bowling

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	pins := [][]string{}
	pinNumber := 0
	targetPinNumber := 0
	pinCount := 0
	//ピン情報を取得する
	for i := 0; i < 4; i++ {
		sc.Scan()
		pinRow := strings.Split(sc.Text(), " ")
		pins = append(pins, pinRow)
	}
	//ピンの残数と狙いを定めるピンを選定する
	for i := 3; i > -1; i-- {
		for j := len(pins[i]) - 1; j > -1; j-- {
			//ピン番号
			pinNumber++
			//残っているピンの残数をカウント
			if pins[i][j] == "1" {
				pinCount++
			}
			//狙うピンを選定
			if pins[i][j] == "1" && targetPinNumber == 0 {
				targetPinNumber = pinNumber
			}
		}
	}
	//狙うピン番号とピンの残数を改行区切りで出力する
	fmt.Println(targetPinNumber)
	fmt.Println(pinCount)
}
