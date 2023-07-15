package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func getKeys(strCount map[string]int) []string {
	var keys []string
	for key := range strCount {
		keys = append(keys, key)
	}
	return keys
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	//この後に入力される回数
	N, _ := strconv.Atoi(sc.Text())

	//文字列が入力されるので、mapに格納する
	//また入力された文字列のカウントを行う。全く同じ文字列が入力されたら、mapのバリューをカウントアップする。
	strCount := make(map[string]int)
	for i := 0; i < N; i++ {
		sc.Scan()
		s := sc.Text()
		if _, ok := strCount[s]; ok {
			strCount[s] += 1
		} else {
			strCount[s] = 1
		}
	}

	//文字列を名前順にする
	keys := getKeys(strCount)
	sort.Strings(keys)

	//名前順で、キー（文字列）、バリュー（カウント）を出力する
	for i := 0; i < len(keys); i++ {
		for k, v := range strCount {
			if k == keys[i] {
				fmt.Printf("%s %d\n", k, v)
			}
		}
	}
}
