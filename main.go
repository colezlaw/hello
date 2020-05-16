package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-redis/redis"
)

func handleRoot(c *redis.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}

		inc := c.Incr("hello:counter")
		if inc.Err() != nil {
			http.Error(w, inc.Err().Error(), http.StatusInternalServerError)
			return
		}

		fmt.Fprintf(w, "this page has been visited %d times", inc.Val())
	}
}

func main() {
	client := redis.NewClient(&redis.Options{
		Addr: "redis:6379",
	})
	_, err := client.Ping().Result()
	if err != nil {
		log.Fatalf("unable to ping redis: %v", err)
	}

	http.HandleFunc("/", handleRoot(client))

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("listen and serve: %v", err)
	}
}
