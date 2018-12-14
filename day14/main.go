package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Println(t1(360781))
	fmt.Println(t2([]byte{3, 6, 0, 7, 8, 1}))
}

func t1(req int) []byte {
	buf := make([]byte, 2, req)
	buf[0], buf[1] = 3, 7

	first := 0
	second := 1

	for len(buf) < req+10 {
		score := buf[first] + buf[second]
		if d := score / 10; d > 0 {
			buf = append(buf, d)
		}

		buf = append(buf, score%10)

		first = (first + int(buf[first]) + 1) % len(buf)
		second = (second + int(buf[second]) + 1) % len(buf)
	}

	return buf[req : req+10]
}

func t2(req []byte) int {
	buf := make([]byte, 2, 2)
	buf[0], buf[1] = 3, 7

	first := 0
	second := 1

	for {
		score := buf[first] + buf[second]
		if d := score / 10; d > 0 {
			buf = append(buf, d)
		}

		buf = append(buf, score%10)

		first = (first + int(buf[first]) + 1) % len(buf)
		second = (second + int(buf[second]) + 1) % len(buf)
		if len(buf) >= 2*len(req) && bytes.Index(buf[len(buf)-2*len(req):], req) != -1 {
			break
		}
	}

	return bytes.Index(buf, []byte(req))
}
