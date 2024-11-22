package api

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type Server struct {
	*mux.Router
	receipts map[string]int // Map to store receipt ID -> points
}

func NewServer() *Server {
	s := &Server{
		Router:   mux.NewRouter(),
		receipts: make(map[string]int),
	}
	s.routes()
	return s
}

func (s *Server) routes() {
	s.HandleFunc("/receipts/process", s.processReceipt()).Methods("POST")
	s.HandleFunc("/receipts/{id}/points", s.getPoints()).Methods("GET")
}

func (s *Server) processReceipt() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var receipt Receipt
		if err := json.NewDecoder(r.Body).Decode(&receipt); err != nil {
			http.Error(w, "Invalid receipt format", http.StatusBadRequest)
			return
		}

		// Calculate points for the receipt
		points, err := receipt.CalculatePoints()
		if err != nil {
			http.Error(w, "Error processing receipt: "+err.Error(), http.StatusBadRequest)
			return
		}

		// Generate unique ID and store receipt points
		id := uuid.New().String()
		s.receipts[id] = points

		// Return the ID in the response
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"id": id})
	}
}

func (s *Server) getPoints() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := mux.Vars(r)["id"]
		points, exists := s.receipts[id]
		if !exists {
			http.Error(w, "Receipt not found", http.StatusNotFound)
			return
		}

		// Return points in the response
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]int{"points": points})
	}
}
