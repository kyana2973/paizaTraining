//問題文URL:https://paiza.jp/works/mondai/string_primer/advance_step6

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
	//出力する値の文字数
	length, _ := strconv.Atoi(sc.Text())
	sc.Scan()
	//置き換える文字と、置き換えるindex番号を取得する
	count, _ := strconv.Atoi(sc.Text())
	var changeIntArr []int
	var changeStrArr []string
	for i := 0; i < count; i++ {
		sc.Scan()
		arr := strings.Split(sc.Text(), " ")
		index, _ := strconv.Atoi(arr[0])
		changeIntArr = append(changeIntArr, index)
		changeStrArr = append(changeStrArr, arr[1])
	}

	sc.Scan()
	//置き換えが必要ない時のデフォルト文字
	defaultStr := sc.Text()

	//文字列を連結していく、置き換えが必要な箇所は文字を差し替える
	var output string
	var checkFlg bool
	for i := 0; i < length; i++ {
		checkFlg = true
		for j := 0; j < count; j++ {
			if i == changeIntArr[j]-1 {
				output = output + changeStrArr[j]
				checkFlg = false
				break
			}
		}
		if checkFlg {
			output = output + defaultStr
		}
	}

	fmt.Println(output)
}
