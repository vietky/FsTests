# Dictionary

## PROBLEM

We were asked to implement a single-file dictionary to lookup a word and show its definition in less than O(N) (N is number of words) with some constraints:
- A word is small string (1024 characters long) which is about 1KB at maximum
- A word's definition is a long string (1GB long)

## SOLUTION

Since the word is small and its explanation is too big to fit in memory, at the first time we read the file, we will put all the words in a dictionary structure with key is **word** and the value is **the address of word's explanation in the file**. So, whenever we check if a word is already in dictionary, we will read to tha address in dictionary file and show the result.

In order to do that, we need to organize the way we store our values:

- For each word adding, we will need:
  1. **word**: 1024 bytes to store the word
  2. **explanation**: 4 bytes to store the size of explanation
  3. **explanation_size**: sizeof(**explanation**) in bytes

When we read through the file byte to byte, we will read first (1024+4) bytes for **word** and its **explanation_size** and store the address (index) position (**address**) of that word.

- To lookup the word, we will read **explanation_size** bytes from **address**.

- To delete a word, we set **word** and **explanation_size** to **null** (both in files and our in-memory cache).

- To update a word, we check if the new explanation_size is still less or equal than current explanation_size:
    - if it is, we set current **explanation** = new **explanation** without updating the **explanation_size**
    - it it isn't, delete the current word and add a new word

Another problem arises when we update/delete an item by doing this way. It will create holes when we remove an item.

To fix this, let's create a command (called defragment :) ), to move all using words to the start of the file to fill up those holes.

## IMPLEMENTATION

We will separate the tasks into multiple classes/functions:
- FileReader: is used to read a dictionary files
  - ReadFile(): reads the file and returns a list of Word Data
  - ReadAtAddress(start, length): get the contains of file at index `start` with length `length` 

- FileWriter: is used to create/write to a dictionary
  - WriteWord(word, explanation): writes `word` and `explanation` to the end of the file
  - WriteAtAddress(start, word, explanation): updates word at position `start`

- MemoryCache: is used to store Word Data and its position in memory for quick access, the inner cache used map<string, WordData>
  - Get(word)
  - Add(word, explanation)
  - Update(word, explanation)
  - Delete(word)

## DEMO

- Run `go run main.go generate` to generate file `data` from `raw_dict.txt`
- Run `go run main.go` to get explanation of some words in `raw_dict.txt`

## PSEUDO CODE (FOR MISSING PARTS IN DEMO):

- Add a word:
```
func add(word, explanation) {
  append_to_file(file_name, new {
    word: word
    explanation_size: of(explanation)
    explanation: explanation
  })
}
```

- To read entire file:
```
address := 0
while !eof {
  word := read_word()
  explanation_size := read_size()
  dict[word].add(new {
    word: word
    address: address
    explanation_size: explanation_size
  })
  address += explanation_size
}
```

- To lookup a word:
```
if dict[word] == null:
  return null;
print read_from_address_with_size({
  address: dict[word].address + 1024+4
  size: dict[word].explanation_size
})
```
- To delete a word:
```
if dict[word] == null:
  return null;
print clear_to_bit_0({
  address: dict[word].address
  size: 1024+4+dict[word].explanation_size
})
```

- To update a word:
```
if dict[word] == null:
  throw 'no word existing in dict';
old_item := dict[word]
if old_item.explanation_size < sizeof(new_explanation) then
  old_item.explanation = new_explanation
else
  delete(word)
  add(word, new_explanation)
```

## CONCLUSION

The algorithm above should lookup the word in O(1) since we're using data structure dictionary as a memory caching to store word and address of the explanation.

Reading the file and build the cache will only cost us O(N) where N is the size of the file.

Other functions (add, update, delete) will take O(1)

Although the task doesn't ask the about the Defragment Command. A good implementation should take O(N) since we're only iterating and copy all words to fill the empty holes.
