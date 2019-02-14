package core

import (
	"log"
	"os"
)

var logger *log.Logger

func init() {
	path, _ := os.LookupEnv("ABQ_TRACE")
	if path == "" {
		return
	}

	file, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	logger = log.New(file, "", 0)
}

func debug(r *Reply, err error) {
	if logger == nil {
		return
	}

	logger.Println()
	if err != nil {
		logger.Println("- error : ", err)
	}

	req := r.Request
	logger.Printf(`curl -nv -X%v %v \`, req.Method, req.URL)
	for h := range req.Header {
		logger.Printf(`-H "%v: %v" \`, h, req.Header.Get(h))
	}
	if len(r.payload) > 0 {
		logger.Printf(`-d '%v'`, string(r.payload))
	}

	logger.Println("< status", r.StatusCode)
	for h := range r.Header {
		logger.Println("< header", h, r.Header.Get(h))
	}
	if len(r.result) > 0 {
		logger.Println("< body")
		logger.Println(string(r.result))
	}
}
