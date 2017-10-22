package dictionary

import (
	"bufio"
	"io"
	"strconv"
)

// FileWriter writer
type FileWriter struct {
	f io.Writer
}

// Write ...
func (writer FileWriter) Write(word, explanation string) {
	var wordByte [MaxWordSize]byte
	var explanationSizeInByte [MaxExplanationSize]byte

	copy(wordByte[:], word)
	explanationByte := []byte(explanation)
	explanationSize := len(explanationByte)
	copy(explanationSizeInByte[:], strconv.Itoa(explanationSize))

	w := bufio.NewWriter(writer.f)
	w.Write(wordByte[:])
	w.Write(explanationSizeInByte[:])
	w.Write(explanationByte)
	w.Flush()
}
