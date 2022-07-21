package main

import (
	"github.com/niranjanative/pie-dictionary/pkg/trie-maintainer"
	"log"
)

func main() {
	err := trie_maintainer.IntializeService()
	if err != nil {
		log.Fatal(err)
	}
}
