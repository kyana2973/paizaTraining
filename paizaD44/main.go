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
	var numbers []string = make([]string, num)
	//スペース区切りで文字列が入力されるので、Splitで格納する
	numbers = strings.Split(sc.Text(), " ")

	var outputNum []string
	var checkFlg int = 0
	//文字列の中で重複している文字は削除して、出力用のsliceに格納する
	for i, v := range numbers {
		if i == 0 {
			outputNum = append(outputNum, v)
		} else {
			for _, v2 := range outputNum {
				if v2 != v {
					checkFlg = 0
				} else {
					checkFlg = 1
				}
			}
			if checkFlg == 0 {
				outputNum = append(outputNum, v)
			}
		}
	}

	//スペース区切りで出力する
	for i, v := range outputNum {
		if i == len(outputNum)-1 {
			fmt.Println(v)
		} else {
			fmt.Print(v + " ")
		}
	}
}
