package trie

import "strings"

// NewNode this will be used to initialize a new node with 26 children
// each child should first be initialized to nil
func NewNode(char string) *Node {
	node := &Node{Char: char}
	for i := 0; i < 26; i++ {
		node.Children[i] = nil
		node.OccurrenceIndexes = IntegerSet{}
	}
	return node
}

// NewTrie Creates a new trie with a root
func NewTrie() *Trie {
	//we will not use this node, so it can be anything
	root := NewNode("\000")
	return &Trie{RootNode: root}
}

// InsertInPrefixTrie inserts reverse string for prefix search to the trie
func (t *Trie) InsertInPrefixTrie(word string, wordIndex int) error {
	current := t.RootNode
	///remove all spaces from the word & convert it to lowercase
	strippedWord := strings.ToLower(strings.ReplaceAll(word, " ", ""))

	// insert string for prefix search
	for i := 0; i < len(strippedWord); i++ {
		// we are taking the decimal representation of a character and subtract decimal representation of 'a' to get
		// index for that char
		index := strippedWord[i] - 'a'
		if current.Children[index] == nil {
			current.Children[index] = NewNode(string(strippedWord[i]))
		}
		// to remember in what all word indexes this char occurs
		current.Children[index].OccurrenceIndexes.Insert(wordIndex)

		current = current.Children[index]
	}

	return nil
}

// InsertInSuffixTrie inserts reverse string for suffix search to the trie
func (t *Trie) InsertInSuffixTrie(word string, wordIndex int) error {
	current := t.RootNode
	// remove all spaces from the word & convert it to lowercase
	strippedWord := strings.ToLower(strings.ReplaceAll(word, " ", ""))

	// insert reverse string for suffix search
	for i := len(strippedWord) - 1; i >= 0; i-- {
		// we are taking the decimal representation of a character and subtract decimal representation of 'a' to get
		// index for that char
		index := strippedWord[i] - 'a'
		if current.Children[index] == nil {
			current.Children[index] = NewNode(string(strippedWord[i]))
		}
		// to remember in what all word indexes this char occurs
		current.Children[index].OccurrenceIndexes.Insert(wordIndex)

		current = current.Children[index]
	}

	return nil
}

// SearchInTrie will return all words which are present in trie, for given str(prefix/suffix/full-word)
// if no match found for str, it will return empty set
func (t *Trie) SearchInTrie(str string, wordList WordList) StringSet {

	// if no words present in wordList, return empty stringSet
	if len(wordList) == 0 {
		return nil
	}

	strippedStr := strings.ToLower(strings.ReplaceAll(str, " ", ""))
	current := t.RootNode
	for i := 0; i < len(strippedStr); i++ {
		index := strippedStr[i] - 'a'
		// if we have encounter null in the path we were transversing, that means this word is not
		// indexed(present) in this trie
		if current == nil || current.Children[index] == nil {
			return nil
		}
		current = current.Children[index]
	}

	wordSet := StringSet{}
	if current != nil {
		for i := range current.OccurrenceIndexes {
			wordSet.Insert(wordList[i])
		}
	}

	return wordSet
}

// Insert adds items to the set.
func (i IntegerSet) Insert(items ...int) IntegerSet {
	for _, item := range items {
		i[item] = Empty{}
	}
	return i
}

// NewIntegerSet initializes new IntegerSet
func NewIntegerSet() IntegerSet {
	integerSet := IntegerSet{}
	return integerSet
}

// Insert adds items to the set.
func (i StringSet) Insert(items ...string) StringSet {
	for _, item := range items {
		i[item] = Empty{}
	}
	return i
}

// NewStringSet initializes new StringSet
func NewStringSet() StringSet {
	stringSet := StringSet{}
	return stringSet
}
