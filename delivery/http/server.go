package http

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/robertotambunan/gdg-jakarta-2025-demo/entity"
	"github.com/robertotambunan/gdg-jakarta-2025-demo/usecase/search"
)

// PageData is the view model for the HTML template.
type PageData struct {
	Title   string
	Results []entity.Fruit
	Query   string
}

// Server holds HTTP dependencies and routes.
type Server struct {
	addr      string
	usecase   search.Usecase
	templates *template.Template
}

// NewServer creates a new HTTP server bound to addr and using the given templates directory.
func NewServer(addr, templatesDir string, uc search.Usecase) (*Server, error) {
	tmplPath := filepath.Join(templatesDir, "index.html")
	tmpl, err := template.ParseFiles(tmplPath)
	if err != nil {
		return nil, err
	}

	return &Server{
		addr:      addr,
		usecase:   uc,
		templates: tmpl,
	}, nil
}

// routes registers the handlers.
func (s *Server) routes() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/", s.indexHandler)
	mux.HandleFunc("/search", s.searchHandler)
	mux.HandleFunc("/api/autocomplete", s.autocompleteHandler)
	return mux
}

// Start runs the HTTP server.
func (s *Server) Start() error {
	log.Printf("HTTP server listening on %s", s.addr)
	return http.ListenAndServe(s.addr, s.routes())
}

func (s *Server) indexHandler(w http.ResponseWriter, r *http.Request) {
	data := PageData{
		Title:   "Cari Buah",
		Results: []entity.Fruit{},
		Query:   "",
	}
	if err := s.templates.Execute(w, data); err != nil {
		log.Printf("error rendering template: %v", err)
	}
}

func (s *Server) searchHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("q")
	if query == "" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	results, err := s.usecase.Autocomplete(r.Context(), query)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := PageData{
		Title:   "Cari Buah",
		Results: results,
		Query:   query,
	}
	if err := s.templates.Execute(w, data); err != nil {
		log.Printf("error rendering template: %v", err)
	}
}

func (s *Server) autocompleteHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("q")
	if query == "" {
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(map[string]interface{}{
			"results": []entity.Fruit{},
			"query":   "",
		})
		return
	}

	results, err := s.usecase.Autocomplete(r.Context(), query)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(map[string]interface{}{
		"results": results,
		"query":   query,
	})
}
