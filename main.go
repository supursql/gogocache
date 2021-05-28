package main

import (
	"fmt"
	"gocache"
	"log"
	"net/http"
)

var db = map[string]string{
	"Tom":  "630",
	"Jack": "589",
	"Sam":  "567",
}

func main() {
	gocache.NewGroup("scores", 2<<10, gocache.GetterFunc(func(key string) ([]byte, error) {
		log.Println("[SlowDB] search key", key)
		if v, ok := db[key]; ok {
			return []byte(v), nil
		}
		return nil, fmt.Errorf("%s not exist", key)
	}))

	addr := "127.0.0.1:9998"
	peers := gocache.NewHTTPPool(addr)
	log.Println("gogocache is running", addr)
	log.Fatal(http.ListenAndServe(":9998", peers))
}
