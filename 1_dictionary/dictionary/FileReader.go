package dictionary

import (
	"bufio"
	"io"
	"os"
	"strconv"
)

// FileReader ...
type FileReader struct {
}

// ReadFile ...
func (reader FileReader) ReadFile(filePath string) []WordData {
	var result []WordData
	f, err := os.Create(filePath)
	check(err)
	defer f.Close()

	buffer := bufio.NewReader(f)
	wordByte := make([]byte, MaxWordSize)
	explanationSizeInByte := make([]byte, MaxExplanationSize)
	var explanationSize int
	currentAddress := 0
	for {
		_, err := buffer.Read(wordByte)
		check(err)
		currentAddress += len(wordByte)

		_, err = buffer.Read(explanationSizeInByte)
		check(err)
		currentAddress += len(explanationSizeInByte)
		explanationSize, err = strconv.Atoi(string(explanationSizeInByte))
		check(err)

		_, err = buffer.Discard(explanationSize)
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}
		result = append(result, WordData{Word: string(wordByte), ExplanationSize: explanationSize})
	}
	return result
}
