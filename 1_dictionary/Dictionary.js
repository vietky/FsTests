
class DictionaryData extends Iterator {
    constructor(word, definition) {
        this.word = word;
        this.definition = definition;
    }
}

class DictionaryLookup {
  get(word) {

  }
}

class DictionaryReader {
    constructor() {
        this.fileName = '';
        // this.iterator = new DictionaryIterator();
    }
    readFile(fileName) {
        this.fileName = fileName;
    }
    readWord() {

    }
}

class DictionaryWriter {
    write(data) {

    }
    save() {
      
    }
}

class Dictionary {
    constructor() {
        this.fileName = '';
    }
    create(fileName) {
        this.fileName = fileName;
    }
    open(fileName) {
        this.fileName = fileName;
    }
    search(word) {

    }
    add(word, definition) {

    }
    delete(word) {

    }
    update(word, definition) {

    }
}

export default Dictionary;
