package main

// Inspired by the noaa firehose sample script
// https://github.com/cloudfoundry/noaa/blob/master/firehose_sample/main.go

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func hello(w http.ResponseWriter, req *http.Request) {
	time.Sleep(time.Minute * 15)
	fmt.Fprintln(w, "done")
}

func setupHTTP() {
	http.HandleFunc("/", hello)

	err := http.ListenAndServe(":"+os.Getenv("PORT"), nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}

}

func main() {
	setupHTTP()
}
