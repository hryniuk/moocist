package api

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/hryniuk/moocist/pkg/mooc"

	"github.com/hryniuk/moocist/pkg/coursera"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
)

type key int

const (
	requestIDKey key = 0
)

type Server struct {
	Router http.Handler
}

func (s *Server) syllabus() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		courseSlug, ok := vars["slug"]
		if !ok && courseSlug == "" {
			log.Errorf("wrong url, slug not found %s", r.URL.Path)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		importer := coursera.SlugImporter{Slug: courseSlug}
		syllabus, err := importer.Import()
		if err != nil {
			log.Errorf("cannot download course syllabus for slug %s: %s", courseSlug, err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		respJson, err := json.Marshal(syllabus)
		if err != nil {
			log.Error(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(respJson)
	}
}

func (s *Server) template() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		courseSlug, ok := vars["slug"]
		if !ok && courseSlug == "" {
			log.Errorf("wrong url, slug not found %s", r.URL.Path)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		importer := coursera.SlugImporter{Slug: courseSlug}
		syllabus, err := importer.Import()
		if err != nil {
			log.Errorf("cannot download course syllabus for slug %s: %s", courseSlug, err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		exporter := mooc.TodoistExporter{mooc.ExportOptions{}}
		csv, err := exporter.Export(syllabus)
		if err != nil {
			log.Errorf("cannot convert course syllabus to template for slug %s: %s", courseSlug, err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/csv")
		w.WriteHeader(http.StatusOK)
		w.Write(csv)
	}
}

func (s *Server) routes(router *mux.Router) {
	router.HandleFunc("/api/syllabus/{slug}", s.syllabus()).Methods("GET")
	router.HandleFunc("/api/template/{slug}", s.template()).Methods("GET")
}

func handleCORS(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set(
			"Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		if (*r).Method == "OPTIONS" {
			// TODO: handle it in a correct manner, set proper status code etc.
			return
		}

		h.ServeHTTP(w, r)
	})
}

func logging(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()
		defer func() {
			requestID, ok := r.Context().Value(requestIDKey).(string)
			if !ok {
				requestID = "unknown"
			}
			duration := time.Now().Sub(startTime) / time.Microsecond
			log.Infof("%s %s %s %s %s %d μs", requestID, r.Method, r.URL.Path, r.RemoteAddr, r.UserAgent(), duration)
		}()
		h.ServeHTTP(w, r)
	})
}

func tracing(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestID := r.Header.Get("X-Request-Id")
		if requestID == "" {
			// requestID = newID()
		}
		ctx := context.WithValue(r.Context(), requestIDKey, requestID)
		w.Header().Set("X-Request-Id", requestID)
		h.ServeHTTP(w, r.WithContext(ctx))
	})
}

func NewServer() *Server {
	router := mux.NewRouter()
	s := &Server{}
	s.routes(router)
	s.Router = tracing(logging(handleCORS(router)))

	return s
}

func initLogging() {
	log.SetLevel(logrus.DebugLevel)
	customFormatter := new(logrus.TextFormatter)
	customFormatter.TimestampFormat = "2006-01-02 15:04:05"
	customFormatter.FullTimestamp = true
	logrus.SetFormatter(customFormatter)
}

func (s *Server) Run(addr string) {
	initLogging()

	httpServer := &http.Server{
		Addr:    addr,
		Handler: s.Router,
	}

	log.Printf("starting server at %s", addr)
	log.Fatal(httpServer.ListenAndServe())
}
