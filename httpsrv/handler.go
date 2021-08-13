package httpsrv

import "net/http"

func (s *Server) listMapel(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world"))
}

func (s *Server) createMapel(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world"))
}

func (s *Server) deleteMapel(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world"))
}
