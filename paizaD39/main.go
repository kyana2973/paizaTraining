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
	//次に入力される文字列の文字数
	N, _ := strconv.Atoi(sc.Text())

	//ランダムなアルファベットが羅列した文字列が入力される
	strArr := make([]string, N)
	sc.Scan()
	strArr = strings.Split(sc.Text(), "")

	//カウンターの判定で使うアルファベット26文字をスライスへ格納する
	alphabet := "abcdefghijklmnopqrstuvwxyz"
	slAlphabet := strings.Split(alphabet, "")

	//入力された文字列を1文字ずつ、どのアルファベットかを判定して、
	//そのアルファベットの順番と同じ位置のカウンタースライスをカウントアップする
	counter := make([]int, 26)
	for _, v := range strArr {
		for i := 0; i < len(slAlphabet); i++ {
			if v == slAlphabet[i] {
				counter[i] = counter[i] + 1
			}
		}
	}

	//カウントの結果を表示する
	var output string
	for i := 0; i < len(counter); i++ {
		s := strconv.Itoa(counter[i])
		if i == len(counter)-1 {
			output = output + s
		} else {
			output = output + s + " "
		}
	}

	fmt.Println(output)
}
