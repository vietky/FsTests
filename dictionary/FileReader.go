package dictionary

import (
	"bufio"
	"bytes"
	"io"
	"os"
	"strconv"
)

// FileReader ...
type FileReader struct {
	FilePath string
}

// ReadFile ...
func (reader FileReader) ReadFile() []WordData {
	var result []WordData
	f, err := os.Open(reader.FilePath)
	check(err)
	defer f.Close()

	buffer := bufio.NewReader(f)
	wordByte := make([]byte, MaxWordSize)
	explanationSizeInByte := make([]byte, MaxExplanationSize)
	var explanationSize int
	currentAddress := 0
	for {
		_, err := buffer.Read(wordByte)
		if err == io.EOF {
			break
		}
		check(err)
		currentAddress += len(wordByte)

		_, err = buffer.Read(explanationSizeInByte)
		check(err)
		currentAddress += len(explanationSizeInByte)

		indexZero := bytes.IndexByte(explanationSizeInByte, 0)
		if indexZero >= 0 {
			explanationSize, err = strconv.Atoi(string(explanationSizeInByte[:indexZero]))
			check(err)
		}
		// fmt.Printf("explanationSizeInByte ne %v %v\n", (explanationSizeInByte), indexZero)
		// fmt.Printf("explanationSize ne %v\n", (explanationSize))
		_, err = buffer.Discard(explanationSize)
		if err != nil {
			if err == io.EOF {
				result = append(result, WordData{Word: string(wordByte), ExplanationSize: explanationSize})
				break
			}
			panic(err)
		}
		result = append(result, WordData{Word: string(wordByte), Address: currentAddress, ExplanationSize: explanationSize})
	}
	return result
}
