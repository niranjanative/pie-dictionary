package trie

// Empty is a void struct
type Empty struct{}

// IntegerSet is a set implementation for int type
type IntegerSet map[int]Empty

// StringSet is a set implementation for string type
type StringSet map[string]Empty

// WordList is list of words whose representations are all present in Trie
type WordList []string

//Node represent each character
type Node struct {

	//this is a single letter stored for example letter a,b,c,d,etc (only a-z lowercase)
	Char string

	// store all children of a node(from a-z lowercase only)
	// a slice of Nodes(and each child will also have 26 children)
	Children [26]*Node

	// contains indexes of words in WordList, in which this char occurs as postfix or prefix
	OccurrenceIndexes IntegerSet
}

// Trie  is our actual tree that will hold all of our nodes, the Root node will be nil
type Trie struct {
	RootNode *Node
}
