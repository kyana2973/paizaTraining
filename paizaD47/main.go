package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	intArr := strings.Split(sc.Text(), " ")
	A, _ := strconv.Atoi(intArr[0])
	B, _ := strconv.Atoi(intArr[1])
	C, _ := strconv.Atoi(intArr[2])
	D, _ := strconv.Atoi(intArr[3])

	//2乗する
	X := math.Pow(float64((A+B)*C), 2)

	//「X mod D」の計算をして出力する
	fmt.Println(math.Mod(X, float64(D)))
}
