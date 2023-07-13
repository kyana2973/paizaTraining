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

	//0〜9がランダムで並んだ文字列が入力される
	numbers := make([]string, N)
	sc.Scan()
	numbers = strings.Split(sc.Text(), " ")

	//numbersから1文字ずつ数値を判定して、その数値のカウンターをカウントアップする
	counter := make([]int, 10)
	for _, v := range numbers {
		i, _ := strconv.Atoi(v)
		switch i {
		case 0:
			counter[0] = counter[0] + 1
		case 1:
			counter[1] = counter[1] + 1
		case 2:
			counter[2] = counter[2] + 1
		case 3:
			counter[3] = counter[3] + 1
		case 4:
			counter[4] = counter[4] + 1
		case 5:
			counter[5] = counter[5] + 1
		case 6:
			counter[6] = counter[6] + 1
		case 7:
			counter[7] = counter[7] + 1
		case 8:
			counter[8] = counter[8] + 1
		case 9:
			counter[9] = counter[9] + 1
		}
	}

	//カウントの結果をスペース区切りで出力
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
