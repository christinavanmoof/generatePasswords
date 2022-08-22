package API

import (
	"net/http"
	"sync"
)

type Generator interface {
	Generate() string
}

type API struct {
	Gen Generator

	once sync.Once
}

func (a *API) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	mux := http.NewServeMux()
	mux.HandleFunc("/generate", func(rw http.ResponseWriter, _ *http.Request) {
		rw.Write([]byte(a.Gen.Generate())) // Write the generated password to the response writer
	})
	mux.ServeHTTP(w, r)
}
