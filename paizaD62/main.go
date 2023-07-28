//問題文URL:https://paiza.jp/works/mondai/string_primer/advance_step10

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
	A := strings.Split(sc.Text(), "")
	sc.Scan()
	B := strings.Split(sc.Text(), "")

	//AとBを足し算する
	var sum int = 0
	var sumArr []string = make([]string, len(A))
	for i := 0; i < len(A); i++ {
		a, _ := strconv.Atoi(A[i])
		b, _ := strconv.Atoi(B[i])
		sum = a + b
		sumArr = append(sumArr, strconv.Itoa(sum))
	}
	fmt.Println(strings.Join(sumArr, ""))
}
