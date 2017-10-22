package dictionary

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// FileWriter writer
type FileWriter struct {
	FilePath string
}

// WriteWord ...
func (writer FileWriter) WriteWord(word, explanation string) {
	var wordByte [MaxWordSize]byte
	var explanationSizeInByte [MaxExplanationSize]byte

	copy(wordByte[:], word)
	explanationByte := []byte(explanation)
	explanationSize := len(explanationByte)
	copy(explanationSizeInByte[:], strconv.Itoa(explanationSize))

	f, err := os.OpenFile(writer.FilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	check(err)

	fmt.Printf("wordByte %v\n", wordByte)
	fmt.Printf("explanationSizeInByte %v\n", explanationSizeInByte)
	fmt.Printf("explanationSize %v\n", explanationSize)

	w := bufio.NewWriter(f)
	w.Write(wordByte[:])
	w.Write(explanationSizeInByte[:])
	w.Write(explanationByte)
	w.Flush()
}
