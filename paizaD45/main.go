//問題文URL:https://paiza.jp/works/mondai/data_structure/data_structure__set_step3

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	//2回目に入力される文字列の数
	num, _ := strconv.Atoi(sc.Text())
	sc.Scan()
	//スペース区切りで文字列が入力されるので、Splitで格納する
	numbers := strings.Split(sc.Text(), " ")

	checkFlg := 0
	var answers []string
	//numbersの2文字目から、slice内に重複した値があるかを検索する
	//その際に、判定する値をslice[i]として、slice[0] <= slice[i-1]の範囲で重複値が無いかを検索する
	for i := 1; i < num; i++ {
		for j := 0; j < i; j++ {
			if numbers[i] == numbers[j] {
				checkFlg = 0
				break
			} else {
				checkFlg = 1
			}
		}
		if checkFlg == 0 {
			answers = append(answers, "Yes")
		} else {
			answers = append(answers, "No")
		}
	}

	//2文字目からの判定を出力する
	for i := 0; i < len(answers); i++ {
		fmt.Println(answers[i])
	}
}
