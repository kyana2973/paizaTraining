//問題文URL:https://paiza.jp/works/mondai/string_primer/advance_step7

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
	//小数点が複数ある場合は、最初の１つ目が小数点となり他は不要となる
	strArr := strings.Split(sc.Text(), ".")
	if len(strArr) >= 2 {
		strArr[0] = strArr[0] + "."
	}
	//分割された文字列を連結する
	str := strings.Join(strArr, "")
	//左側にある「0」はを削除する
	str = strings.TrimLeft(str, "0")
	//小数点が存在する場合で、右側に不要な「0」が存在する場合は削除する
	if strings.Contains(str, ".") {
		str = strings.TrimRight(str, "0")
	}
	//文字列の最初が「.」の場合は、先頭に「0」を付与する
	for i, r := range str {
		if i == 0 && string(r) == "." {
			str = "0" + str
			break
		}
	}
	fmt.Println(str)
}
