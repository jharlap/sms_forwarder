package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

var forwardTo string

func main() {
	port := os.Getenv("PORT")
	forwardTo = os.Getenv("FORWARD_TO")

	log.Printf("Starting server. PORT: %s FORWARD_TO: %s", port, forwardTo)
	http.ListenAndServe("0.0.0.0:"+port, http.HandlerFunc(handle))
}

func handle(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Println("Error parsing form:", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	from := r.Form.Get("From")
	body := r.Form.Get("Body")

	log.Printf("Message from %s: %s", from, body)

	w.Header().Set("Content-Type", "application/xml")
	tpl := `<?xml version="1.0" encoding="UTF-8" ?>
<Response>
    <Message>This number does not support texts. Please call.</Message>
    <Message to="%s">From %s:
    %s</Message>
</Response>
`
	fmt.Fprintf(w, tpl, forwardTo, from, body)
}
