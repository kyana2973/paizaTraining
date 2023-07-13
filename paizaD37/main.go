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
	strArr := strings.Split(sc.Text(), " ")
	//H = Rows、W = Columns、r = rowNumber、c = columnNumber
	H, _ := strconv.Atoi(strArr[0])
	W, _ := strconv.Atoi(strArr[1])
	r, _ := strconv.Atoi(strArr[2])
	c, _ := strconv.Atoi(strArr[3])

	//２次元スライスを、H行とW列で構成する
	mazeRow := make([][]string, H)
	for i := 0; i < H; i++ {
		mazeRow[i] = make([]string, W)
	}

	//「..#.」といった文字列が、H行だけ入力されるので２次元スライスへ格納する
	for i := 0; i < H; i++ {
		sc.Scan()
		mazeRow[i] = strings.Split(sc.Text(), "")
	}

	//mazeRow[r-1][c-1]の値が、「#」の時は「Yes」で、「.」の時は「No」と出力する
	if mazeRow[r-1][c-1] == "#" {
		fmt.Println("Yes")
	} else if mazeRow[r-1][c-1] == "." {
		fmt.Println("No")
	}
}
