// Here it starts
package main // import "github.com/sirile/go-image-test"

import (
	"fmt"
	"net/http"
	"os"

	"github.com/op/go-logging"
)

var colour string
var hostname string
var log = logging.MustGetLogger("go-image-test")
var format = logging.MustStringFormatter(
	`%{color}%{time:15:04:05.000} %{shortfunc} ▶ %{level:.4s} %{id:03x}%{color:reset} %{message}`,
)

func main() {
	logging.SetBackend(logging.NewBackendFormatter(logging.NewLogBackend(os.Stderr, "", 0), format))
	// Get the hostname and calculate a hash based on it
	name, err := os.Hostname()
	hostname = name
	hash := 0
	for i := 0; i < len(name); i++ {
		hash = int(name[i]) + ((hash << 5) - hash)
	}
	// Generate a colour based on the hostname hash
	colour = "#"
	for i := uint(0); i < 3; i++ {
		hex := "00" + fmt.Sprintf("%x", ((hash>>i*8)&0xFF))
		colour += hex[len(hex)-2:]
	}

	if err == nil {
		log.Debugf("Serving image, host: %s, colour: %s", name, colour)
		http.HandleFunc("/", handler)
		http.ListenAndServe(":80", nil)
	}
}

// Serve the content
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<!DOCTYPE html><html><head>"+
		"<title>Node scaling demo</title></head><body>"+
		"<h1>%s</h1>"+
		"<svg xmlns=\"http://www.w3.org/2000/svg\" viewBox=\"0 0 400 400\">"+
		"<circle cx=\"50\" cy=\"50\" r=\"48\" fill=\"%s\" stroke=\"#000\"/>"+
		"<path d=\"M50,2a48,48 0 1 1 0,96a24 24 0 1 1 0-48a24 24 0 1 0 0-48\" fill=\"#000\"/>"+
		"<circle cx=\"50\" cy=\"26\" r=\"6\" fill=\"#000\"/>"+
		"<circle cx=\"50\" cy=\"74\" r=\"6\" fill=\"#FFF\"/>"+
		"</svg>"+
		"</body></html>", hostname, colour)
}
