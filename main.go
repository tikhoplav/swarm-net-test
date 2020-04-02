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
		w.Header().Set("Content-Type", "text/html; charset=utf-8")

		html := make([]string, 0, 1024)
		html = append(html, "<html><body>")
		html = append(html, "<h3>Swarm net test:</h3>")
		html = append(html, fmt.Sprintf("<p>Current container time: %v<br>", time.Now()))
		html = append(html, "<i>(if not changed after reload, then this page is cached)</i></p>")
		html = append(html, fmt.Sprintf("<p>Container id: <b>%v</b><br>", DockerId()))
		html = append(html, "<i>(this id is the same as in `docker ps`)</i><br>")
		html = append(html, "<i>(Also this id can be found running `docker node ps $node-name`)</i></p>")
		html = append(html, "</body></html>")

		w.Write([]byte(strings.Join(html, "\n")))
	})

	fmt.Println("Starting server at 8080")
	http.ListenAndServe("0.0.0.0:8080", nil)
}

func DockerId() string {
	cont, err := ioutil.ReadFile("/proc/self/cgroup")
	if err != nil {
		return "not a docker container"
	}

	return fmt.Sprintf("%s", cont[20:31])
}