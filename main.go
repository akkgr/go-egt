package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Error reading body: %v", err)
		http.Error(w, "can't read body", http.StatusBadRequest)
		return
	}

	s := string(body[:1])

	switch s {
	case "H":
		d, err := Heartbeat(body)
		if err != nil {
			log.Fatal("Heartbeat Error:", err.Error())
			fmt.Fprintf(w, "%s", err.Error())
		}
		fmt.Fprintf(w, "%s", d)

	case "D":
		d, err := Data(body)
		if err != nil {
			log.Fatal("Heartbeat Error:", err.Error())
			fmt.Fprintf(w, "%s", err.Error())
		}
		fmt.Fprintf(w, "%s", d)
	case "C":
		fmt.Fprintf(w, "%s", s)
	case "U":
		fmt.Fprintf(w, "%s", s)
	default:
		fmt.Fprintf(w, "Invalid Message Type.")
	}
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
