package dictionary

import (
	"bufio"
	"io"
	"os"
	"strconv"
)

// FileReader ...
type FileReader struct {
	FilePath string
}

// ReadAtAddress ...
func (reader FileReader) ReadAtAddress(start, length int64) []byte {
	f, err := os.OpenFile(reader.FilePath, os.O_RDONLY, 0644)
	// f, err := os.Open(reader.FilePath)
	check(err)
	defer f.Close()

	_, err = f.Seek(start+MaxWordSize+MaxExplanationSize, io.SeekStart)
	check(err)
	buf := make([]byte, length)
	_, err = f.Read(buf)
	check(err)
	return buf
}

// ReadFile ...
func (reader FileReader) ReadFile() []WordData {
	var result []WordData
	f, err := os.Open(reader.FilePath)
	check(err)
	defer f.Close()

	buffer := bufio.NewReader(f)
	currentAddress := 0
	for {
		wordByte := make([]byte, MaxWordSize)
		explanationSizeInByte := make([]byte, MaxExplanationSize)
		var explanationSize int

		position := currentAddress

		_, err = readEnoughBytes(buffer, wordByte)
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}
		currentAddress += len(wordByte)

		_, err = readEnoughBytes(buffer, explanationSizeInByte)
		check(err)
		currentAddress += len(explanationSizeInByte)

		explanationSize, err = strconv.Atoi(getCleanStringFromByteArray(explanationSizeInByte))

		// log.Printf("wordByte %v %v %v\n", (wordByte), len(wordByte), wordCount)
		// log.Printf("explanationSizeInByte %v %v %v\n", (explanationSizeInByte), len(explanationSizeInByte), explanationSizeInByteCount)
		// log.Printf("explanationSize %v\n", explanationSize)

		_, err = buffer.Discard(explanationSize)
		check(err)

		currentAddress += explanationSize
		result = append(result, WordData{Word: getCleanStringFromByteArray(wordByte), Address: position, ExplanationSize: explanationSize})
	}
	return result
}
