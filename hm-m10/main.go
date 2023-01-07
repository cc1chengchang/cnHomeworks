package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/prometheus/client_golang/prometheus/promhttp"

	_ "net/http/pprof"
)

func main() {
	flag.Set("v", "4")
	fmt.Println("Starting http server...")

	log.SetOutput(os.Stdout)

	Register()

	mux := &http.ServeMux{}
	mux.HandleFunc("/", rootHandler)
	mux.HandleFunc("/healthz", healthz)
	mux.HandleFunc("/ready", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
	mux.Handle("/metrics", promhttp.Handler())
	var handler http.Handler = mux
	handler = RequestLogger(handler)

	//http.HandleFunc("/", rootHandler)
	//http.HandleFunc("/healthz", healthz)
	//err := http.ListenAndServe(":8090", nil)
	err := http.ListenAndServe(":8080", handler)
	if err != nil {
		log.Fatal(err)
	}

}

func healthz(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, "200\n")
}

func randFloat64(min float64, max float64) float64 {
	return min + rand.Float64()*(max-min)
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("entering root handler")

	//add random delay
	timer := NewTimer()
	defer timer.ComputeTotal()
	delay := randFloat64(0, 2000)
	time.Sleep(time.Duration(delay) * time.Millisecond)

	user := r.URL.Query().Get("user")
	if user != "" {
		io.WriteString(w, fmt.Sprintf("hello [%s]\n", user))
	} else {
		io.WriteString(w, "hello [stranger]\n")
	}
	io.WriteString(w, "===================Details of the http request header:============\n")
	for k, v := range r.Header {
		io.WriteString(w, fmt.Sprintf("%s=%s\n", k, v))
	}
	io.WriteString(w, "===================Server go version:============\n")
	//ver := os.Getenv("GOVERSION")
	//fmt.Printf("ver: %s", ver)
	io.WriteString(w, fmt.Sprintf("Version=%s\n", os.Getenv("GOVERSION")))
}

type ResponseWithRecorder struct {
	http.ResponseWriter
	statusCode int
	body       bytes.Buffer
}

func (rec *ResponseWithRecorder) WriteHeader(statusCode int) {
	rec.ResponseWriter.WriteHeader(statusCode)
	rec.statusCode = statusCode
}

func (rec *ResponseWithRecorder) Write(d []byte) (n int, err error) {
	n, err = rec.ResponseWriter.Write(d)
	if err != nil {
		return
	}
	rec.body.Write(d)

	return
}

func RequestLogger(targetMux http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("entering logging")

		wc := &ResponseWithRecorder{
			ResponseWriter: w,
			statusCode:     http.StatusOK,
			body:           bytes.Buffer{},
		}

		targetMux.ServeHTTP(wc, r)

		// log request by who(IP address)
		requesterIP := r.RemoteAddr

		log.Printf(
			"%s\t\t%s\t\t%s\t\t%v",
			r.Method,
			r.RequestURI,
			requesterIP,
			wc.statusCode,
		)
	}
	return http.HandlerFunc(fn)
}
