package dictionary

import (
	"bufio"
	"io"
	"strconv"
)

// FileWriter writer
type FileWriter struct {
	Writer io.Writer
}

// WriteWord ...
func (writer FileWriter) WriteWord(word, explanation string) {
	var wordByte [MaxWordSize]byte
	var explanationSizeInByte [MaxExplanationSize]byte

	copy(wordByte[:], word)
	explanationByte := []byte(explanation)
	explanationSize := len(explanationByte)
	copy(explanationSizeInByte[:], strconv.Itoa(explanationSize))

	// log.Printf("wordByte %v %v %v\n", word, (wordByte), len(wordByte))
	// log.Printf("explanationSizeInByte %v %v\n", (explanationSizeInByte), len(explanationSizeInByte))
	// log.Printf("explanationSize %v\n", explanationSize)

	w := bufio.NewWriter(writer.Writer)
	w.Write(wordByte[:])
	w.Write(explanationSizeInByte[:])
	w.Write(explanationByte)
	w.Flush()

}
