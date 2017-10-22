package dictionary

import (
	"bufio"
	"bytes"
	"fmt"
)

func check(err error) {
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}

func getCleanStringFromByteArray(byteArr []byte) string {
	indexZero := bytes.IndexByte(byteArr, 0)
	if indexZero >= 0 {
		return string(byteArr[:indexZero])
	}
	return string(byteArr[:])
}

func readEnoughBytes(buffer *bufio.Reader, byteArr []byte) ([]byte, error) {
	n, err := buffer.Read(byteArr)
	if err != nil {
		return byteArr, err
	}
	check(err)
	if n < len(byteArr) {
		n, err = buffer.Discard(len(byteArr) - n)
		check(err)
	}
	return byteArr, nil
}
