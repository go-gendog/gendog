package main

import (
	"encoding/json"
	"log"
	"net/http"

	"example/gen"
)

func main() {
	server := Server{}
	h := gen.Handler(&server)

	s := &http.Server{
		Handler: h,
		Addr:    "0.0.0.0:8080",
	}

	if err := s.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}

var _ gen.ServerInterface = (*Server)(nil)

type Server struct{}

func (s Server) GetPing(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(gen.PingResponse{
		Message: "pong",
	})
}
