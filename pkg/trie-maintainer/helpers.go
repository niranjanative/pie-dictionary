package trie_maintainer

import (
	"github.com/niranjanative/pie-dictionary/internal/common"
	"github.com/niranjanative/pie-dictionary/internal/models"
	"github.com/niranjanative/pie-dictionary/internal/trie"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
)

// AddToWordList adds given key to wordList
func AddToWordList(key string) (indexOfAddedWord int) {
	mutex.Lock()
	indexOfAddedWord = len(wordList)
	wordList = append(wordList, key)
	mutex.Unlock()
	return
}

// ReadKeyValue gets value for given key from redis server
func ReadKeyValue(context *gin.Context) {

	key := context.Param("key")

	val, err := client.Get(key).Result()
	if err != nil && err == redis.Nil {
		context.JSON(http.StatusNotFound, gin.H{common.DataString: trie.StringSet{}})
		return
	} else if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{common.ErrorString: err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{common.DataString: val})
}

// InsertKeysInTrie is a wrapper function which populates all given keys in Prefix & Postfix tries
func InsertKeysInTrie(keys ...string) {

	for keyIndex, key := range keys {
		wordList = append(wordList, key)
		if err := prefixTrie.InsertInPrefixTrie(key, keyIndex); err != nil {
			log.Panicf("error populating prefixTrie for key: %s. Error:%s", keys[keyIndex], err.Error())
		}
		if err := suffixTrie.InsertInSuffixTrie(key, keyIndex); err != nil {
			log.Panicf("error populating prefixTrie for key: %s. Error:%s", keys[keyIndex], err.Error())
		}
	}

	return
}

// SearchKeys returns all keys present in trie which matches given prefix/postfix
func SearchKeys(c *gin.Context) {

	var (
		matchingKeys trie.StringSet
	)

	prefix := c.Query(common.PrefixString)
	postfix := c.Query(common.PostFixString)

	if (prefix == "" && postfix == "") || (prefix != "" && postfix != "") {
		c.JSON(http.StatusBadRequest, gin.H{common.ErrorString: "Only 1 QueryParams required. " +
			"Supported QueryParams: [prefix, postfix]"})
		return
	} else if prefix != "" {
		matchingKeys = prefixTrie.SearchInTrie(prefix, wordList)
	} else if postfix != "" {
		matchingKeys = suffixTrie.SearchInTrie(postfix, wordList)
	}

	if len(matchingKeys) == 0 {
		c.JSON(http.StatusNotFound, gin.H{common.DataString: nil})
		return
	}

	c.JSON(http.StatusOK, gin.H{common.DataString: matchingKeys})

}

// InsertKeyValue inserts given key-value pair in redis & in both tries
func InsertKeyValue(c *gin.Context) {

	var input models.KeyValue

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{common.ErrorString: err.Error()})
		return
	}

	// write key-value to redis server
	if setErr := client.Set(input.Key, input.Value, 0).Err(); setErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{common.ErrorString: setErr.Error()})
		return
	}

	// add word to wordList
	indexOfAddedWord := AddToWordList(input.Key)

	if err := prefixTrie.InsertInPrefixTrie(input.Key, indexOfAddedWord); err != nil {
		log.Fatal(err)
	}

	if err := suffixTrie.InsertInSuffixTrie(input.Key, indexOfAddedWord); err != nil {
		log.Fatal(err)
	}

	c.JSON(http.StatusOK, gin.H{common.DataString: input})
}
