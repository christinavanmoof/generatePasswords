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
	a.once.Do(func() {
		mux := http.NewServeMux()
		mux.HandleFunc("/generate", func(rw http.ResponseWriter, _ *http.Request) {
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			rw.Write([]byte(a.Gen.Generate())) // Write the generated password to the response writer
		})
		mux.ServeHTTP(w, r)
	})
}
