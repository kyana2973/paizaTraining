//問題文URL:https://paiza.jp/works/mondai/string_primer/advance_step9

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
	//文字列で計算式を受け取る
	strArr := strings.Split(sc.Text(), "")

	//数字と演算子と分ける
	var editArr []string
	var saveStr string
	for i, v := range strArr {
		_, e := strconv.Atoi(v)
		if i == len(strArr)-1 {
			saveStr = saveStr + v
			editArr = append(editArr, saveStr)
		}

		if e == nil {
			saveStr = saveStr + v
		} else {
			editArr = append(editArr, saveStr)
			saveStr = ""
			editArr = append(editArr, v)
		}
	}

	//受け取った計算式を数値型で実行して出力する
	var sum int
	var key string
	for i, v := range editArr {
		num, _ := strconv.Atoi(v)

		if i == 0 {
			sum = num
			continue
		}

		if (i+1)%2 == 0 && v == "+" {
			key = "+"
			continue
		} else if (i+1)%2 == 0 && v == "-" {
			key = "-"
			continue
		}

		if key == "+" {
			sum += num
		} else {
			sum -= num
		}
	}
	fmt.Println(sum)
}
