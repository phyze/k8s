package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	started := time.Now()
	http.HandleFunc("/healthz-readiness", func(w http.ResponseWriter, r *http.Request) {
		tnow := time.Now()
		diff := tnow.Sub(started).Seconds()
		if diff > 15 {
			w.WriteHeader(200)
			w.Write([]byte("ok"))
			fmt.Println(diff)
		} else {
			w.WriteHeader(400)
			w.Write([]byte("fail"))
			fmt.Println(diff)
		}
	})

	http.HandleFunc("/healthz-liveness", func(w http.ResponseWriter, r *http.Request) {
		tnow := time.Now().Minute()
		if tnow%3 == 0 {
			w.WriteHeader(400)
			w.Write([]byte("fail"))
			fmt.Println(tnow)
		} else {
			w.WriteHeader(200)
			w.Write([]byte("ok"))
			fmt.Println(tnow)
		}
	})

	http.HandleFunc("/secret", func(w http.ResponseWriter, r *http.Request) {
		pwd, _ := os.observeEnv("PASSWORD")
		w.Write([]byte(pwd))
		w.WriteHeader(200)
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		hostname, _ := os.observeEnv("HOSTNAME")
		w.Write([]byte(hostname))
		w.WriteHeader(200)
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
