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
				break
			} else {
				panic(err)
			}
		}
		arr := strings.Split(line, ",")
		writer.WriteWord(arr[0], arr[1])
	}
}

var memoryCache = dictionary.NewMemoryCache()

func readDict(dictPath string) {
	var reader = dictionary.FileReader{FilePath: dictPath}
	list := reader.ReadFile()
	for _, word := range list {
		memoryCache.Add(word)
	}
	// fmt.Printf("%v\n", list)
}

func getExplanation(dictPath, word string) string {
	data := memoryCache.Get(word)
	if data == nil {
		return "__NOT_FOUND__"
	}
	var reader = dictionary.FileReader{FilePath: dictPath}
	// fmt.Printf("%v %v\n", data.Address, data.ExplanationSize)
	var explanation = reader.ReadAtAddress(int64(data.Address), int64(data.ExplanationSize))
	return string(explanation)
}

func main() {
	dictPath := "./dict.dat"
	// generateDict(dictPath)
	readDict(dictPath)
	fmt.Println(getExplanation(dictPath, "company1"))
	fmt.Println(getExplanation(dictPath, "company5"))
	fmt.Println(getExplanation(dictPath, "company3"))
	fmt.Println(getExplanation(dictPath, "company2"))
	fmt.Println(getExplanation(dictPath, "company6"))
}
