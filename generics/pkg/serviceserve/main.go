package serviceserve

import (
	"context"
	"flag"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/ropenttd/tsubasa/generics/pkg/responses"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync/atomic"
	"time"
)

type key int

const (
	requestIDKey key = 0
)

var (
	listenAddr string
	Healthy    int32
	Logger     *log.Logger
)

type loggingResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func NewLoggingResponseWriter(w http.ResponseWriter) *loggingResponseWriter {
	return &loggingResponseWriter{w, http.StatusOK}
}

func (lrw *loggingResponseWriter) WriteHeader(code int) {
	lrw.statusCode = code
	lrw.ResponseWriter.WriteHeader(code)
}

func init() {
	flag.StringVar(&listenAddr, "listen-addr", ":8000", "Service listen address")
	flag.Parse()

	Logger = log.New(os.Stdout, "http: ", log.LstdFlags)
}

func Serve(router *mux.Router) error {
	router.Handle("/healthz", healthz())
	router.NotFoundHandler = responses.HandleNotFound()

	nextRequestID := func() string {
		return fmt.Sprintf("%d", time.Now().UnixNano())
	}

	server := &http.Server{
		Addr:         listenAddr,
		Handler:      tracing(nextRequestID)(logging(Logger)(router)),
		ErrorLog:     Logger,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	done := make(chan bool)
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	go func() {
		<-quit
		Logger.Println("Server is shutting down...")
		atomic.StoreInt32(&Healthy, 0)

		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		server.SetKeepAlivesEnabled(false)
		if err := server.Shutdown(ctx); err != nil {
			Logger.Fatalf("Could not gracefully shutdown the server: %v\n", err)
		}
		close(done)
	}()

	Logger.Println("🚀 Ready for requests on", listenAddr)
	atomic.StoreInt32(&Healthy, 1)
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		Logger.Fatalf("Could not listen on %s: %v\n", listenAddr, err)
		return err
	}

	<-done
	Logger.Println("Server stopped")
	return nil
}

func healthz() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if atomic.LoadInt32(&Healthy) == 1 {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		w.WriteHeader(http.StatusServiceUnavailable)
	})
}

func logging(logger *log.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			lrw := NewLoggingResponseWriter(w)
			next.ServeHTTP(lrw, r)
			requestID, ok := r.Context().Value(requestIDKey).(string)
			if !ok {
				requestID = "unknown"
			}
			statusCode := lrw.statusCode
			logger.Println(requestID, r.Method, r.URL.Path, statusCode, r.RemoteAddr, r.UserAgent())
		})
	}
}

func tracing(nextRequestID func() string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			requestID := r.Header.Get("X-Request-Id")
			if requestID == "" {
				requestID = nextRequestID()
			}
			ctx := context.WithValue(r.Context(), requestIDKey, requestID)
			w.Header().Set("X-Request-Id", requestID)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
