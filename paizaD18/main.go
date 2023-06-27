package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// 1〜n/2の値をint型のスライスに格納
func FirstRow(n int) []int {
	nHalf := n / 2
	var intArr []int
	for i := 0; i < nHalf; i++ {
		intArr = append(intArr, i+1)
	}
	return intArr
}

// n/2〜nの値をint型のスライスに格納
func SecondRow(n int) []int {
	nHalf := n / 2
	var intArr []int
	for i := nHalf; i < n; i++ {
		intArr = append(intArr, i+1)
	}
	return intArr
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())

	var FirstRow []int = FirstRow(n)
	var SecondRow []int = SecondRow(n)

	//1行目に、1〜n/2の値を出力
	for i, v := range FirstRow {
		s := strconv.Itoa(v)
		if (i + 1) == len(FirstRow) {
			fmt.Println(s)
			break
		}
		fmt.Print(s + " ")
	}

	//2行目に、n/2〜１nの値を出力
	for i, v := range SecondRow {
		s := strconv.Itoa(v)
		if (i + 1) == len(SecondRow) {
			fmt.Println(s)
			break
		}
		fmt.Print(s + " ")
	}

}
