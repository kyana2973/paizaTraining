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
	num := strings.Split(sc.Text(), " ")
	//N = 店の商品名の数、M = 買い物リストの商品名の数
	N, _ := strconv.Atoi(num[0])
	M, _ := strconv.Atoi(num[1])

	//店の商品名と値段が入力されるので、mapに格納する
	shopItems := make(map[string]int)
	for i := 0; i < N; i++ {
		sc.Scan()
		item := strings.Split(sc.Text(), " ")
		price, _ := strconv.Atoi(item[1])
		shopItems[item[0]] = price
	}

	//買い物リストの名称が入力されるので、sliceに格納する
	shoppingList := make([]string, M)
	for i := 0; i < M; i++ {
		sc.Scan()
		shoppingList[i] = sc.Text()
	}

	//店に買い物リストの商品が存在した場合は、値段を改行して表示する
	//存在しない場合は、「-1」を表示する
	checkFlg := 0
	for _, s := range shoppingList {
		for k, v := range shopItems {
			if s == k {
				checkFlg = 0
				fmt.Println(v)
				break
			} else {
				checkFlg = 1
			}
		}
		if checkFlg == 1 {
			fmt.Println(-1)
		}
	}
}
