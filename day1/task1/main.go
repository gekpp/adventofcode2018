package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func main() {
	var res int64 = 0

	stdin := os.Stdin
	r := bufio.NewReader(stdin)
	for {
		bytes, _, err := r.ReadLine()
		if err == io.EOF || len(bytes) == 0 {
			break
		}
		str := string(bytes)

		t, _ := strconv.ParseInt(str[1:], 10, 32)
		if str[0] == '+' {
			res += t
		} else {
			res -= t
		}
		fmt.Printf("%s = %v\n", str, res)
	}

	fmt.Println(res)
}
