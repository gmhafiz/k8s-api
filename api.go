package app

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func (s *Server) Healthz(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func (s *Server) Ready(w http.ResponseWriter, r *http.Request) {
	_, err := s.DB.Exec(r.Context(), "SELECT true;")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		msg := fmt.Sprintf(`{"message": "%s"}`, err)
		log.Println(msg)
		_, _ = w.Write([]byte(msg))
		return
	}
	w.WriteHeader(http.StatusOK)
}

type Random struct {
	ID   int    `db:"id" json:"id"`
	Name string `db:"name" json:"name"`
}

func (s *Server) Randoms(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		s.get(w, r)
	} else if r.Method == http.MethodPost {
		s.add(w, r)
	}

}

func (s *Server) get(w http.ResponseWriter, r *http.Request) {
	rows, err := s.DB.Query(r.Context(), "SELECT * FROM randoms ORDER BY id DESC LIMIT 10;")
	if err != nil {
		msg := fmt.Sprintf(`{"message": "%s"}`, err)
		log.Printf("error with query: %v", err)
		log.Println(msg)
		_, _ = w.Write([]byte(msg))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var randoms []Random
	for rows.Next() {
		var r Random
		if err := rows.Scan(&r.ID, &r.Name); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Printf("error with database: %v", err)
			_, _ = w.Write([]byte(`{"message": "error with database"}`))
			return
		}
		randoms = append(randoms, r)
	}

	resp, err := json.Marshal(randoms)
	if err != nil {
		log.Printf("error with json: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(`{"message": "error with json"}`))
		return
	}

	_, _ = w.Write(resp)
}

func (s *Server) add(w http.ResponseWriter, r *http.Request) {
	_, err := s.DB.Exec(r.Context(), "INSERT INTO randoms values(gen_random_uuid()); ")
	if err != nil {
		return
	}
}
