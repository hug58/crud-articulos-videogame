package server

import (
	"fmt"
	"go-postgres/internal/server/handlers"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi"
)

type Server struct {
	server *http.Server
}

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprint(w, "hello\n")
}

func New(port string) (*Server, error) {

	r := chi.NewRouter()

	r.Mount("/api/v1", handlers.New())

	serv := &http.Server{
		Addr:         ":" + port,
		Handler:      r,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	//hello with name
	r.Get("/hello/{name}", func(w http.ResponseWriter, r *http.Request) {
		name := chi.URLParam(r, "name")
		fmt.Fprintf(w, "Hello %s", name)
	})

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		//welcome
		fmt.Fprintf(w, "<h1> Welcome to the home page! </h1>")
	})

	server := Server{server: serv}
	return &server, nil
}

func (serv *Server) Close() error {
	return nil
}

func (serv *Server) Start() {
	log.Printf("Server running on http://localhost%s", serv.server.Addr)

	log.Fatal(serv.server.ListenAndServe())
}
