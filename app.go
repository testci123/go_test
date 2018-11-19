package main

import (
    "fmt"
    "net/http"
	"errors"
)

func Division(a, b float64) (float64, error) {
    if b == 0 {
		return 0, errors.New("除数不能为0..12")
	}

	return a / b, nil
}

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "46: aa Hello123 123 World")
    })

    http.ListenAndServe(":8080", nil)
}
