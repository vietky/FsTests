package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/vietky/FsTests/dictionary"
)

func generateDict(dictPath string) {
	var writer = dictionary.FileWriter{FilePath: dictPath}
	f, err := os.Open("./raw_dict.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	buf := bufio.NewReader(f)
	for {
		line, err := buf.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				arr := strings.Split(line, ",")
				writer.WriteWord(arr[0], arr[1])
				break
			} else {
				panic(err)
			}
		}
		arr := strings.Split(line, ",")
		writer.WriteWord(arr[0], arr[1])
	}
}

var memoryCache = dictionary.MemoryCache{}

func readDict(dictPath string) {
	var reader = dictionary.FileReader{FilePath: dictPath}
	list := reader.ReadFile()
	fmt.Printf("%v\n", list)
}

func main() {
	dictPath := "./dict.dat"
	// generateDict(dictPath)
	readDict(dictPath)
}
