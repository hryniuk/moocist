package api

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/hryniuk/moocist/pkg/coursera"

	"github.com/gomodule/redigo/redis"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

type key int

const (
	requestIDKey key = 0
)

type Server struct {
	Router http.Handler
	dbConn redis.Conn
}

const cacheKeyFormat = "syllabus:%s"
const baseCourseURLFormat = "https://www.coursera.org/learn/%s#syllabus"

func (s *Server) course() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		courseSlug, ok := vars["slug"]
		if !ok && courseSlug == "" {
			log.Errorf("wrong url, slug not found %s", r.URL.Path)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		cacheKey := fmt.Sprintf(cacheKeyFormat, courseSlug)

		syllabusJson, err := redis.Bytes(s.dbConn.Do("GET", cacheKey))
		if err == nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write(syllabusJson)
			return
		}

		courseUrl := fmt.Sprintf(baseCourseURLFormat, courseSlug)
		syllabus, err := coursera.GetCourseInfo(courseUrl)
		if err != nil {
			log.Errorf("cannot download course syllabus for slug %s: ", courseSlug, err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		respJson, err := json.Marshal(syllabus)
		if err != nil {
			log.Error(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		s.dbConn.Do("SET", cacheKey, respJson)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(respJson)
	}
}

func (s *Server) routes(router *mux.Router) {
	router.HandleFunc("/course/{slug}", s.course()).Methods("GET")
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
	conn, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		log.Fatal(err)
	}

	router := mux.NewRouter()
	s := &Server{dbConn: conn}
	s.routes(router)
	s.Router = tracing(logging(handleCORS(router)))

	return s
}
