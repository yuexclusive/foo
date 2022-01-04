package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

func main() {
	http.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		name, err := os.Hostname()

		if err != nil {
			panic(err)
		}
		fmt.Fprintln(w, struct {
			Host string
			Time string
		}{Host: name, Time: time.Now().Format(time.RFC3339)})
	})

	http.ListenAndServe(":8080", nil)
}
