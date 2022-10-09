// package xample

// import (
// 	"time"
// 	"fmt"
// 	"net/http"
// 	"github.com/gorilla/mux"
// )

// func index() {
//   router := mux.NewRouter()

//   router.Use(middleware1)

//   wsRouter := router.PathPrefix("/ws").Subrouter()
//   wsRouter.Use(middleware2)
//   wsRouter.Use(middleware3)

//   wsRouter.HandleFunc("/sub", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//     fmt.Println("handling ws /sub")
//     w.Write([]byte("/sub (ws)"))
//   }))

//   chainRouter := router.PathPrefix("/chain").Subrouter()
//   chainRouter.HandleFunc("/sub1", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//     fmt.Println("handling chain /sub1")
//     w.Write([]byte("/sub1"))
//   }))
//   chainRouter.HandleFunc("/sub2", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//     fmt.Println("handling chain /sub2")
//     w.Write([]byte("/sub2"))
//   }))

//   restRouter := router.PathPrefix("/").Subrouter()
//   restRouter.Use(middleware3)

//   restRouter.HandleFunc("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//     fmt.Println("handling rest /")
//     w.Write([]byte("/ (rest)"))
//   }))

//   server := &http.Server{
//         Addr:         "0.0.0.0:3010",
//         // Good practice to set timeouts to avoid Slowloris attacks.
//         WriteTimeout: time.Second * 15,
//         ReadTimeout:  time.Second * 15,
//         IdleTimeout:  time.Second * 60,
//         Handler: router, // Pass our instance of gorilla/mux in.
//     }

//   fmt.Println("starting server")
//   if err := server.ListenAndServe(); err != nil {
//     fmt.Println(err)
//   }

  
// }

// func middleware1(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		fmt.Println("middleware1")
// 		next.ServeHTTP(w, r)
// 	})
// }

// func middleware2(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		fmt.Println("middleware2")
// 		next.ServeHTTP(w, r)
// 	})
// }

// func middleware3(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		fmt.Println("middleware3")
// 		next.ServeHTTP(w, r)
// 	})
// }