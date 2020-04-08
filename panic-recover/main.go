package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"net/http"
	"runtime/debug"
)

// type ResponseWriter interface {
// 	Header() Header
// 	Write([]byte) (int, error)
// 	WriteHeader(statusCode int)
// }
type rw struct {
	http.ResponseWriter
	writes [][]byte
	status int
}

// func (r *rw) Header() {

// }

func (r *rw) Write(b []byte) (int, error) {
	r.writes = append(r.writes, b)
	return len(b), nil
}

func (r *rw) WriteHeader(statusCode int) {
	r.status = statusCode
}

func (r *rw) Hijack() (net.Conn, *bufio.ReadWriter, error) {
	if hijacker, ok := r.ResponseWriter.(http.Hijacker); !ok {
		return nil, nil, fmt.Errorf("Response writer does not support the Hijacker interface")
	}
}

func (r *rw) flush() error {
	r.ResponseWriter.Write(r.writes[len(r.writes)-1])
	r.ResponseWriter.WriteHeader(r.status)
	return nil
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/panic/", panicDemo)
	mux.HandleFunc("/panic-after/", panicAfterDemo)
	mux.HandleFunc("/", hello)
	fmt.Println("Server running on http://localhost:3000")
	log.Fatal(http.ListenAndServe(":3000", recoverMw(mux)))
}

func recoverMw(app http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if e := recover(); e != nil {
				buf := debug.Stack()
				log.Println(string(buf))
				w.WriteHeader(http.StatusInternalServerError)
				http.Error(w, string(buf), http.StatusInternalServerError)
			}
		}()
		resW := rw{ResponseWriter: w}
		app.ServeHTTP(&resW, r)
		resW.flush()
	}
}

func panicDemo(w http.ResponseWriter, r *http.Request) {
	funcThatPanics()
}

func panicAfterDemo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Hello!</h1>")
	funcThatPanics()
}

func funcThatPanics() {
	panic("Oh no!")
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "<h1>Hello!</h1>")
}
