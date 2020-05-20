package main

import (
	"errors"
	"flag"
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func main() {
	addr := flag.String("addr", ":80", "interface and port to listen on")
	flag.Parse()

	// Initialize logger.
	logger := log.New(os.Stdout, "http: ", log.LstdFlags)

	// Determine the CODE env variable.
	name, ok := os.LookupEnv("CODE")
	if !ok {
		name = "CODE_NOT_FOUND"
	}

	// Set up http server.
	http.HandleFunc("/hello", logRequest(logger, helloHandler(name)))
	http.HandleFunc("/ready", logRequest(logger, readyHandler()))
	http.HandleFunc("/live", logRequest(logger, liveHandler(time.Now())))

	// File server.
	fs := http.FileServer(http.Dir("/tmp"))
	http.HandleFunc("/files/", logRequest(logger, http.StripPrefix("/files/", fs).ServeHTTP))

	// File upload.
	uploadFileDir := "/uploadfiles"
	http.HandleFunc("/upload", logRequest(logger, uploadHandler(uploadFileDir)))
	{
		fs := http.FileServer(http.Dir(uploadFileDir))
		http.HandleFunc(uploadFileDir+"/", logRequest(logger, http.StripPrefix(uploadFileDir+"/", fs).ServeHTTP))
	}

	// A game.
	game := http.FileServer(http.Dir("/game"))
	http.HandleFunc("/", logRequest(logger, game.ServeHTTP))

	// Same game protected with a secret.
	username, password, err := lookupAuthParameters()
	if err == nil {
		game = http.StripPrefix("/secret", http.FileServer(http.Dir("/game")))
		http.HandleFunc("/secret/", logRequest(logger, basicAuth(username, password, game.ServeHTTP)))
	} else {
		logger.Println("not configuring basic-auth endpoint because:", err)
	}

	logger.Printf("server is starting on %s", *addr)
	err = http.ListenAndServe(*addr, nil)
	if err != nil {
		logger.Fatal(err)
	}
}

// logRequest is a middleware that logs some request details
func logRequest(logger *log.Logger, next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logger.Println(r.Method, r.Host, r.URL.Path, r.RemoteAddr, r.UserAgent())
		next.ServeHTTP(w, r)
	}
}

// helloHandler responds to request with given name
func helloHandler(name string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %s\n", name)
	}
}

// readyHandler is used for readiness probe
func readyHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "ok\n")
	}
}

// liveHandler is used for liveness probe and fakes being dead after 20 seconds
func liveHandler(start time.Time) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		duration := time.Since(start)
		if duration.Seconds() > 600 {
			w.WriteHeader(500)
			fmt.Fprintf(w, "error: not alive any more\n")
		} else {
			fmt.Fprintf(w, "ok\n")
		}
	}
}

// lookupAuthParameters fetches basic auth username and password
func lookupAuthParameters() (username, password string, err error) {
	// Username should be found in WORKSHOP_USERNAME env variable.
	name, ok := os.LookupEnv("WORKSHOP_USERNAME")
	if !ok {
		return "", "", errors.New("missing env variable: $WORKSHOP_USERNAME")
	}

	// Password should be found in the file at /opt/password.
	buf, err := ioutil.ReadFile("/opt/password")
	if err != nil {
		return "", "", err
	}

	// Remove any newlines.
	pass := strings.Split(string(buf), "\n")[0]
	return name, pass, nil
}

// basicAuth is a middleware that ensures the user is authenticated
func basicAuth(username, password string, next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		u, p, ok := r.BasicAuth()
		if !ok || u != username || p != password {
			w.Header().Set("WWW-Authenticate", `Basic realm="k8s-workshop"`)
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	}
}

// uploadHandler handles file upload
func uploadHandler(fileDir string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet, http.MethodHead:
			t, err := template.ParseFiles("/templates/upload.gtpl")
			if err != nil {
				log.Fatalln(err)
			}

			t.Execute(w, nil)
		case http.MethodPost:
			err := r.ParseMultipartForm(32 << 20)
			if err != nil {
				log.Fatalln(err)
			}

			headers := r.MultipartForm.File["uploadfile"]
			for _, header := range headers {
				file, err := header.Open()
				if err != nil {
					log.Println(err)
					return
				}
				defer file.Close()

				// Open a file for writing.
				filename := filepath.Join(fileDir, header.Filename)
				f, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE, 0666)
				if err != nil {
					log.Println(err)
					return
				}
				defer f.Close()

				// Write the file.
				_, err = io.Copy(f, file)
				if err != nil {
					log.Fatalln(err)
				}

				// Redirect to files.
				w.Header().Add("Location", fileDir)
				w.WriteHeader(http.StatusSeeOther)
			}
		}
	}
}
