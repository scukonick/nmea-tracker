package main

import (
	"encoding/json"
	"log"
	"net/http"
)

// APIServer serves all the requests from the frontend
type APIServer struct {
	informer Informer
}

// NewAPIServer returns a pointer to newly initialized APIServer
func NewAPIServer(informer Informer) *APIServer {
	i := &APIServer{
		informer: informer,
	}
	return i
}

// ServeHTTP implements http.Handler interface
func (s *APIServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Printf("Request %s from %s", r.URL, r.RemoteAddr)
	token := r.URL.Query().Get("token")
	if len(token) != 25 {
		log.Printf("Wrong token: %s", token)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	points, err := s.informer.GetPoints(token)
	if err != nil {
		log.Printf("Error getting points from database: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	log.Printf("Got %d points", len(points))

	bJSON, err := json.Marshal(points)
	if err != nil {
		log.Printf("Error json encoding points: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
	_, err = w.Write(bJSON)
	if err != nil {
		log.Printf("Error writing json to client")
	}

}
