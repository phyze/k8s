package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Budget struct {
	Id string `json:"id"`
	Year string `json:"year"`
	Receive string `json:"receive"`
}

func main() {
	http.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(200)
			w.Write([]byte("ok"))
			return 
	})

	http.HandleFunc("/financial", func(w http.ResponseWriter, r *http.Request) {
		budgets := make([]Budget,0)
		budget1 := Budget{
			Id: "1",
			Year: "2019",
			Receive: "200,000,000",
		}
		budget2 := Budget{
			Id: "2",
			Year: "2020",
			Receive: "3,000,000",
		}
		budget3 := Budget{
			Id: "3",
			Year: "2021",
			Receive: "33,000,000",
		}
		budgets = append(budgets, budget1, budget2, budget3)
		b,_ := json.Marshal(&budgets)
		w.WriteHeader(200)
		fmt.Fprintln(w, string(b))
		return
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}