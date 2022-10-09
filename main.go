// package main

// import (
// 	"net/http"
// 	"github.com/gorilla/mux"
// 	"github.com/gorilla/handlers"
// 	"api-kafka/src/helper"
// 	"os"
// 	"fmt"
// 	"time"
// 	"log"
// )

// type Route struct {
//     Name        string
//     Method      string
//     Pattern     string
//     Secure      bool
//     HandlerFunc http.HandlerFunc
// }

// type Routes []Route

// var routes = Routes{
//     Route{
//         Name:        "GetUserByName",
//         Method:      "GET",
//         Pattern:     "/v2/user/{username}",
//         HandlerFunc: GetUserByName,
//         Secure:      true,
//     },
// }

// func GetUserByName(w http.ResponseWriter, r *http.Request) {
// 	data := []map[string]interface{}{
// 		{
// 			"id":           1,
// 			"nama_product": "Kemeja",
// 			"stok":         1000,
// 		},
// 		{
// 			"id":           2,
// 			"nama_product": "Celana",
// 			"stok":         10000,
// 		},
// 		{
// 			"id":           1,
// 			"nama_product": "Sepatu",
// 			"stok":         500,
// 		},
// 	}

// 	helper.ResponseJSON(w, http.StatusOK, data)
// }

// func NewRouter() *mux.Router {
//     router := mux.NewRouter().StrictSlash(true)
//     router.NotFoundHandler = http.HandlerFunc(notFound)
//     router.MethodNotAllowedHandler = http.HandlerFunc(notAllowed)
//     for _, route := range routes {
//         var handler http.Handler
//         if route.Secure {
//             handler = AuthMiddleware(route.HandlerFunc)
//         } else {
//             handler = route.HandlerFunc
//         }

// 			handler = handlers.LoggingHandler(os.Stdout, router)

//         router.
//             Methods(route.Method).
//             Path(route.Pattern).
//             Name(route.Name).
//             Handler(handler)
//     }

//     return router
// }

// func ApplicationRecovery(next http.Handler) http.Handler {
//     return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//         defer func() {
//             if err := recover(); err != nil {
//                 fmt.Fprintln(os.Stderr, "Recovered from application error occurred")
//                 _, _ = fmt.Fprintln(os.Stderr, err)
//                 w.WriteHeader(http.StatusInternalServerError)
                
//             }
//         }()
//         next.ServeHTTP(w, r)
//     })
// }

// func Middleware(next http.Handler) http.Handler {
//     return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//         w.Header().Add("Content-Type", "application/json")
//         next.ServeHTTP(w, r)
//     })
// }

// func AuthMiddleware(next http.Handler) http.Handler {
//     return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//         //TODO: Add authentication
//         log.Println("Authentication required")
//         next.ServeHTTP(w, r)
//     })
// }

// func Logger(inner http.Handler, name string) http.Handler {
//     return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//         start := time.Now()

//         log.Printf(
//             "%s %s %s %s",
//             r.Method,
//             r.RequestURI,
//             name,
//             time.Since(start),
//         )

//         inner.ServeHTTP(w, r)
//     })
// }

// func notFound(w http.ResponseWriter, r *http.Request) {
//     w.WriteHeader(http.StatusNotFound)
// }

// func notAllowed(w http.ResponseWriter, r *http.Request) {
//     w.WriteHeader(http.StatusMethodNotAllowed)
// }

// func main() {
// 	port := "3010"
//     srv := &http.Server{
//         Addr:         "0.0.0.0:" + port,
//         Handler:      ApplicationRecovery(Middleware(NewRouter())),
//         ReadTimeout:  15 * time.Second,
//         WriteTimeout: 15 * time.Second,
//     }
// 	log.Println("Starting server on port", port)
//     log.Fatal(srv.ListenAndServe())
// }