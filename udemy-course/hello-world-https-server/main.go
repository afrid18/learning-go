package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		n, err := fmt.Fprintf(w, "Hello World!")
		if err != nil {
			fmt.Println("Error is: ", err)
		}
		fmt.Println(fmt.Sprintf("Total number of bytes written are: %d", n))
	})

	_ = http.ListenAndServe(":8080", nil)
}

