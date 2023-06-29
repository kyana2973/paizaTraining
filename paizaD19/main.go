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
	//標準入力から「5 8」といった値を受け取り、左辺は「N」に、右辺は「M」へ格納する
	strArr := strings.Split(sc.Text(), " ")
	N, _ := strconv.Atoi(strArr[0])
	M, _ := strconv.Atoi(strArr[1])

	//1行目に、1〜Nまでの値をスペース区切りで出力して行末で改行する
	for i := 0; i < N; i++ {
		s := strconv.Itoa(i + 1)
		if (i + 1) == N {
			fmt.Println(s)
			break
		}
		fmt.Print(s + " ")
	}

	//2行目に、1〜Mまでの値をスペース区切りで出力して行末で改行する
	for i := 0; i < M; i++ {
		s := strconv.Itoa(i + 1)
		if (i + 1) == M {
			fmt.Println(s)
			break
		}
		fmt.Print(s + " ")
	}
}
