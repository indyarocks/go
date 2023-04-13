package main

import (
	"bytes"
	"fmt"
)

func intToString(a []int) string {
	var buff bytes.Buffer
	buff.WriteByte('[')
	for i, v := range a {
		if i > 0 {
			buff.WriteByte(',')
		}
		fmt.Fprintf(&buff, "%d", v)
	}
	buff.WriteByte(']')
	return buff.String()
}

func commaRecursive(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return commaRecursive(s[:n-3]) + "," + commaRecursive(s[n-3:])
}

func byteBuffer(s string) string {
	var buff bytes.Buffer
	for i := 0; i < len(s); i++ {
		if i > 0 && len(s[i:])%3 == 0 {
			buff.WriteByte(',')
		}

		fmt.Fprintf(&buff, "%v", string(s[i]))
	}
	return buff.String()
}

func main() {
	fmt.Println(intToString([]int{1, 2, 3, 4}))
	fmt.Println(commaRecursive("1234566"))
	fmt.Println(byteBuffer("1234566787878988989891"))
}
