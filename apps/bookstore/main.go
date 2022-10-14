package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Book struct {
	Id string `json:"id"`
	Name string `json:"name"`
	Author string `json:"author"`
}

func main() {
	http.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(200)
			w.Write([]byte("ok"))
			return 
	})

	http.HandleFunc("/bookstore", func(w http.ResponseWriter, r *http.Request) {
		books := make([]Book,0)
		book1 := Book{
			Id: "1",
			Name: "Sapiens: A Brief History of Humankind",
			Author: "Ûwan",
		}
		book2 := Book{
			Id: "2",
			Name: "Clean code",
			Author: "Robert",
		}
		book3 := Book{
			Id: "3",
			Name: "Domain Driven Design",
			Author: "Eric",
		}
		books = append(books, book1, book2, book3)
		b,_ := json.Marshal(&books)
		w.WriteHeader(200)
		fmt.Fprintln(w, string(b))
		return
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}