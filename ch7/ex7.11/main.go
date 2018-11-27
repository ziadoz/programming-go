package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"
)

var mutex sync.Mutex

type database map[string]dollars

func (db database) list(w http.ResponseWriter, req *http.Request) {
	mutex.Lock()
	for item, price := range db {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}
	mutex.Unlock()
}

func (db database) price(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")

	mutex.Lock()
	price, ok := db[item]
	mutex.Unlock()

	if !ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "no such item: %q\n", item)
		return
	}
	fmt.Fprintf(w, "%s\n", price)
}

// https://stackoverflow.com/questions/3050518/what-http-status-response-code-should-i-use-if-the-request-is-missing-a-required
func (db database) create(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	if item == "" {
		http.Error(w, "item cannot be empty", http.StatusUnprocessableEntity)
		return
	}

	mutex.Lock()
	_, ok := db[item]
	mutex.Unlock()
	if ok {
		http.Error(w, "item already exists", http.StatusUnprocessableEntity)
		return
	}

	price, _ := strconv.Atoi(req.URL.Query().Get("price"))
	if price == 0 {
		http.Error(w, "price cannot be zero", http.StatusUnprocessableEntity)
		return
	}

	mutex.Lock()
	db[item] = dollars(price)
	mutex.Unlock()

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "successfully created %s", item)
}

func (db database) update(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	mutex.Lock()
	_, ok := db[item]
	mutex.Unlock()
	if !ok {
		http.Error(w, "item does not exist", http.StatusUnprocessableEntity)
		return
	}

	price, _ := strconv.Atoi(req.URL.Query().Get("price"))
	if price == 0 {
		http.Error(w, "price cannot be zero", http.StatusUnprocessableEntity)
		return
	}

	mutex.Lock()
	db[item] = dollars(price)
	mutex.Unlock()

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "successfully updated %s", item)
}

type dollars float32

func (d dollars) String() string {
	return fmt.Sprintf("$%.2f", d)
}

func main() {
	db := database{"shoes": 50, "socks": 5}
	http.HandleFunc("/list", db.list)
	http.HandleFunc("/price", db.price)
	http.HandleFunc("/create", db.create)
	http.HandleFunc("/update", db.update)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
