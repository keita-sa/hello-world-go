package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path) // r.URL.Pathはリクエストのパスを返す, %sで文字列の表示
	})

	http.ListenAndServe(":80", nil)
}
