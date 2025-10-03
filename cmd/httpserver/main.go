package main

import (
	"fmt"
	"geecache"
	"log"
	"net/http"
)

var db = map[string]string{
	"Tom":      "2233",
	"Jerry":    "998",
	"takeboat": "99999",
}

func main() {
	geecache.NewGroup("test-scores", 2<<10, geecache.GetterFunc(func(key string) ([]byte, error) {
		log.Println("[mock-db] search key", key)
		if v, ok := db[key]; ok {
			return []byte(v), nil
		}
		return nil, fmt.Errorf("%s not exist", key)
	}))

	addr := "localhost:8899"
	peers := geecache.NewHttpPool(addr)
	log.Println("geecaceh is running at", addr)
	log.Fatal(http.ListenAndServe(addr, peers))
}
