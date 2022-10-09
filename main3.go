package main
import (
    "net/http"
    "github.com/gorilla/mux"
	"log"
	
	"time"

)
func main() {
    r := mux.NewRouter()
    r.HandleFunc("/ping", handlePing).Methods(http.MethodGet)
    r.HandleFunc("/hello", handleHello).Methods(http.MethodGet)
	r.Use(RequestLoggerMiddleware(r))
    // http.ListenAndServe(":http", r)
	// handler := handlers.LoggingHandler(os.Stdout, r)
	port := "3010"
    srv := &http.Server{
        Addr:         "0.0.0.0:" + port,
        Handler:      r,
        ReadTimeout:  15 * time.Second,
        WriteTimeout: 15 * time.Second,
    }
	log.Println("Starting server on port", port)
    log.Fatal(srv.ListenAndServe())
}
func handlePing(rw http.ResponseWriter, req *http.Request) {
    rw.Write([]byte("I am working...\n"))
    return
}

func RequestLoggerMiddleware(r *mux.Router) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
        return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
            start := time.Now()
            sw := NewStatusResponseWriter(w)
            defer func() {
                log.Printf(
                    "[%s] [%v] [%d] %s %s %s",
                    req.Method,
                    time.Since(start),
                    sw.statusCode,
                    req.Host,
                    req.URL.Path,
                    req.URL.RawQuery,
                )
            }()
            next.ServeHTTP(sw, req)
        })
    }
}

type ResponseWriter interface {
    // Header() Header
    Write([]byte) (int, error)
    WriteHeader(statusCode int)
}

func handleHello(rw http.ResponseWriter, req *http.Request) {
    rw.WriteHeader(http.StatusOK)
    rw.Write([]byte("Hi, how are you...\n"))
    return
}

type statusResponseWriter struct {
    http.ResponseWriter
    statusCode int
}
// NewStatusResponseWriter returns pointer to a new statusResponseWriter object
func NewStatusResponseWriter(w http.ResponseWriter) *statusResponseWriter {
    return &statusResponseWriter{
        ResponseWriter: w,
        statusCode:     http.StatusOK,
    }
}

func (sw *statusResponseWriter) WriteHeader(statusCode int) {
    sw.statusCode = statusCode
    sw.ResponseWriter.WriteHeader(statusCode)
}