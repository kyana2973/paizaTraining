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
	strArr := strings.Split(sc.Text(), "")
	sc.Scan()
	mountStr := sc.Text()
	sc.Scan()
	mountIndex, _ := strconv.Atoi(sc.Text())

	//mountIndexの時に、mountStrを末尾に連結する
	var output string
	for i, v := range strArr {
		if i == mountIndex-1 {
			output = output + v + mountStr
		} else {
			output = output + v
		}
	}
	fmt.Println(output)
}
