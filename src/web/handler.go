package web

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"log"
	"net"
	"net/http"
	"regexp"

	"runtime/debug"
	"syscall"
	"time"

	"story_writer/src/lib"
)

// each handler can return the data and error, and serveHTTP can chose how to convert this
type HandlerFunc func(rw http.ResponseWriter, r *http.Request) (interface{}, error)

var excludeRegex *regexp.Regexp

func shouldTraceURL(url string) bool {
	if excludeRegex == nil || !excludeRegex.Match([]byte(url)) {
		return true
	}
	return false
}

func (fn HandlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer func() {
		recov := recover()
		if recov != nil {
			w.WriteHeader(500)
			errorStack := fmt.Sprintf("``` %s ```", string(debug.Stack()))
			capture(fmt.Sprintf("Panic happening in http handler %s\n", errorStack))
		}
	}()

	response := Response{}
	response.Base.Status = "OK"

	// CORS
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")

	if r.Method == "OPTIONS" {
		w.WriteHeader(200)
		return
	}

	ctx := r.Context()

	ctx, cancelFn := context.WithTimeout(ctx, 10*time.Second)
	defer cancelFn()

	r = r.WithContext(ctx)

	start := time.Now()

	var data interface{}
	var err error
	errStatus := http.StatusInternalServerError

	w.Header().Set("Content-Type", "application/json")

	data, err = fn(w, r)
	response.Base.ServerProcessTime = time.Since(start).String()

	var buf []byte

	if data != nil && err == nil {
		response.Data = data
		if buf, err = response.MarshalJSON(); err == nil {
			w.Write(buf)
			return
		}
	}

	if err != nil {
		response.Base.ErrorMessage = []string{
			err.Error(),
		}

		switch t := err.(type) {
		case lib.APIError:
			errStatus = t.Status
			response.Data = t
		case net.Error:
			if t.Timeout() {
				response.Base.ErrorMessage = []string{
					ErrConnectivity.Error(),
				}
			}
		case *net.OpError:
			// t.Op == "dial",  "read" , "write"
			response.Base.ErrorMessage = []string{
				ErrConnectivity.Error(),
			}
		case syscall.Errno:
			// t == syscall.ECONNREFUSED
			response.Base.ErrorMessage = []string{
				ErrConnectivity.Error(),
			}
		}

		log.Println("handler error", err.Error(), r.URL.Path)

		w.WriteHeader(errStatus)
	}

	if w.Header().Get("Content-Type") == "text/csv" {
		return
	}

	buf, _ = response.MarshalJSON()
	log.Println(string(buf[:]))
	w.Write(buf)
	return
}

// capture panics
func capture(err string, message ...string) {
	var tmp string
	for i, val := range message {
		if i == 0 {
			tmp += val
		} else {
			tmp += fmt.Sprintf("\n\n%s", val)
		}
	}

	publishError(errors.New(err), []byte(tmp), false)
}

func publishError(errs error, reqBody []byte, withStackTrace bool) {
	var buffer bytes.Buffer
	buffer.WriteString(fmt.Sprintf(`*%s*`, errs.Error()))

	if reqBody != nil {
		buffer.WriteString(fmt.Sprintf(" ```%s``` ", string(reqBody)))
	}
}
