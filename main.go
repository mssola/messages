// Copyright (C) 2019 Miquel Sabaté Solà <mikisabate@gmail.com>
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/go-redis/redis"
	"github.com/gorilla/mux"
)

// indexFile contains the exact location of the index file to be rendered. It's
// set by the `setIndexFile` function.
var indexFile = ""

// Page contains the data needed to render the index page.
type Page struct {
	// Messages contains the messages as fetched from redis to be displayed on
	// the main page.
	Messages []string
}

// waitForRedis tries to connect to the redis server over and over again. It
// will fail eventually after a timeout.
func waitForRedis() {
	var addr string
	var err error
	var client *redis.Client

	for i := 0; i < 100; i += 1 {
		client, addr = getClient()
		_, err = client.Ping().Result()
		if err != nil {
			log.Printf("waiting for redis...")
			time.Sleep(5 * time.Second)
		} else {
			return
		}
	}
	log.Fatalf("Could not connect to redis on '%v': %v", addr, err)
}

// setIndexFile sets the exact location if the index file to be rendered by the
// `index` function.
func setIndexFile() {
	indexFile = os.Getenv("MESSAGES_FILE_PATH")
	if indexFile != "" {
		indexFile = filepath.Join(indexFile, "index.html")
	} else {
		indexFile = "index.html"
	}
	log.Printf("file to be served is %v...", indexFile)
}

// getClient fetches a redis client. It also returns a second parameter which
// contains the address of the redis server.
func getClient() (*redis.Client, string) {
	host := os.Getenv("MESSAGES_REDIS_HOST")
	if host == "" {
		host = "localhost"
	}
	port := os.Getenv("MESSAGES_REDIS_PORT")
	if port == "" {
		port = "6379"
	}
	password := os.Getenv("MESSAGES_REDIS_PASSWORD")

	addr := fmt.Sprintf("%v:%v", host, port)
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       0,
	})
	return client, addr
}

// getSafeClient fetches a client for redis. If it cannot be done, it will log
// an error message and quit abruptly.
func getSafeClient() *redis.Client {
	client, addr := getClient()

	_, err := client.Ping().Result()
	if err != nil {
		log.Fatalf("Could not connect to redis on '%v': %v", addr, err)
	}
	return client

}

// postMessage saves the given message into redis and redirects to the main
// page.
func postMessage(w http.ResponseWriter, req *http.Request) {
	log.Print("handling /message")

	req.ParseForm()

	message := req.Form["message"][0]
	client := getSafeClient()
	err := client.RPush("messages", message).Err()
	if err != nil {
		log.Printf("error: %v", err.Error())
		fmt.Fprintf(w, "oops! take a look at the logs for more info...")
		return
	}
	http.Redirect(w, req, "/", http.StatusFound)
}

// index renders the main page.
func index(w http.ResponseWriter, req *http.Request) {
	log.Print("handling /")

	client := getSafeClient()
	result, err := client.LRange("messages", 0, -1).Result()
	if err != nil {
		fmt.Fprintf(w, "oops!")
		http.Redirect(w, req, "/", http.StatusFound)
		return
	}

	t, _ := template.ParseFiles(indexFile)
	t.Execute(w, &Page{Messages: result})
}

// health is a simple handler that just returns 200.
func health(w http.ResponseWriter, req *http.Request) {
	log.Print("handling /health")
}

func main() {
	setIndexFile()
	waitForRedis()

	log.Print("listening on :3000")

	r := mux.NewRouter()
	r.HandleFunc("/message", postMessage).Methods("POST")
	r.HandleFunc("/health", health).Methods("GET")
	r.HandleFunc("/", index).Methods("GET")

	log.Fatal(http.ListenAndServe(":3000", r))
}
