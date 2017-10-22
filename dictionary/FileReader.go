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
	f, err := os.Open(reader.FilePath)
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
	wordByte := make([]byte, MaxWordSize)
	explanationSizeInByte := make([]byte, MaxExplanationSize)
	var explanationSize int
	currentAddress := 0
	for {
		position := currentAddress

		_, err := buffer.Read(wordByte)
		if err == io.EOF {
			break
		}
		check(err)
		currentAddress += len(wordByte)

		_, err = buffer.Read(explanationSizeInByte)
		if err == io.EOF {
			break
		}
		check(err)
		currentAddress += len(explanationSizeInByte)

		explanationSize, err = strconv.Atoi(getCleanStringFromByteArray(explanationSizeInByte))

		_, err = buffer.Discard(explanationSize)
		currentAddress += explanationSize

		if err != nil {
			if err == io.EOF {
				result = append(result, WordData{Word: string(wordByte), ExplanationSize: explanationSize})
				break
			}
			panic(err)
		}
		result = append(result, WordData{Word: getCleanStringFromByteArray(wordByte), Address: position, ExplanationSize: explanationSize})
	}
	return result
}
