package dictionary

import (
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
