package dictionary

// MaxWordSize default size for every word
const MaxWordSize = 1024
const MaxExplanationSize = 20

// WordData ...
type WordData struct {
	Word            string
	Address         int
	ExplanationSize int
	Explanation     *string
}
