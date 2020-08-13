package main

import (
	"fmt"
	pluginsimple "github.com/sdm2345/go-monitor/plugin-simple"
	"net/http"
)

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "hello\n")
	})
	h := pluginsimple.Init(mux)
	http.ListenAndServe(":8290", h)
}
