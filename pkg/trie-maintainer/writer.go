package trie_maintainer

import (
	"github.com/niranjanative/pie-dictionary/internal/common"
	"github.com/niranjanative/pie-dictionary/internal/trie"
	"os"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
)

var (
	prefixTrie = trie.NewTrie()
	suffixTrie = trie.NewTrie()
	mutex      = &sync.Mutex{}

	client             *redis.Client
	redisIP, redisPort string
	wordList           trie.WordList
)

func init() {
	redisIP = os.Getenv(common.RedisIPEnvName)
	redisPort = os.Getenv(common.RedisPortEnvName)
}

func IntializeService() error {

	client = redis.NewClient(&redis.Options{
		Addr:     redisIP + ":" + redisPort, //"localhost:6379",
		Password: "",
		DB:       0,
	})

	// get all keys from redisServer, usefull for cases when service restarts
	keys := client.Keys("*").Val()

	//populate prefix & postfix trie's
	InsertKeysInTrie(keys...)

	// start web server
	router := gin.Default()

	router.GET("/search", SearchKeys)
	router.GET("/get/:key", ReadKeyValue)
	router.POST("/set", InsertKeyValue)

	// Run the server
	err := router.Run(":4321")

	return err
}
