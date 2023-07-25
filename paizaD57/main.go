//問題文URL:

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
	strArr := strings.Split(sc.Text(), "")

	var outputs []string
	var checkFlg bool
	//重複した文字は削除して出力する
	for i, v := range strArr {
		if i == 0 {
			outputs = append(outputs, v)
		} else {
			for _, w := range outputs {
				if w != v {
					checkFlg = true
				} else {
					checkFlg = false
					break
				}
			}
			if checkFlg {
				outputs = append(outputs, v)
			}
		}
	}
	fmt.Println(strings.Join(outputs, ""))
}
