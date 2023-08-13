//問題文URL:https://paiza.jp/works/mondai/array_primer/array_primer__elm_change

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	num := strings.Split(sc.Text(), " ")
	var output []string
	for i := len(num) - 1; -1 < i; i-- {
		output = append(output, num[i])
	}
	fmt.Println(strings.Join(output, " "))
}
