package main

import (
	"fmt"
	"net/http"
	"strings"
	"io/ioutil"
	"time"
)

func main() {
	http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
		var body []byte

		if accept, ok := r.Header["Accept"]; ok && strings.Contains(accept[0], "text/html") {
			w.Header().Set("Content-Type", "text/html; charset=utf-8")
			body = CreateHTML()
		} else {
			w.Header().Set("Content-Type", "application/json; charset=utf-8")
			body = CreateJSON()
		}

		w.Header().Set("Content-Length", fmt.Sprintf("%v", len(body)))
		w.Write(body)
	})

	fmt.Println("Starting server at 8080")
	http.ListenAndServe("0.0.0.0:8080", nil)
}

func DockerId() string {
	cont, err := ioutil.ReadFile("/proc/self/cpuset")
	if err != nil {
		return "not a docker container?"
	}

	s := fmt.Sprintf("%s", cont)
	ss := strings.Split(s, "/")

	return strings.TrimSpace(ss[2])
}

func CreateHTML() []byte {
	html := make([]string, 0, 1024)
	html = append(html, "<html><body>")
	html = append(html, "<h3>Swarm net test:</h3>")
	html = append(html, fmt.Sprintf("<p>Current container time: %v<br>", time.Now()))
	html = append(html, "<i>(if not changed after reload, then this page is cached)</i></p>")
	html = append(html, fmt.Sprintf("<p>Container id: <b>%v</b><br>", DockerId()))
	html = append(html, "<i>(this id is the same as in `docker ps`)</i></p>")
	html = append(html, "</body></html>")
	return []byte(strings.Join(html, ""))
}

func CreateJSON() []byte {
	json := make([]string, 0, 1024)
	json = append(json, "{")
	json = append(json, fmt.Sprintf("\"timestamp\": \"%v\",", time.Now()))
	json = append(json, fmt.Sprintf("\"container_id\": \"%v\"", DockerId()))
	json = append(json, "}")
	return []byte(strings.Join(json, ""))
}