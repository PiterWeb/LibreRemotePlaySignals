package server

import (
	"fmt"
	"net/http"
	"strconv"
)

const allAddresses = "0.0.0.0"

const (
	hostRole = "host"
	clientRole = "client"
)

func Init(port uint16, ips_listening chan<- []string) error {

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {

		// Handle WebSocket connection
		fmt.Printf("WebSocket connection established from %s\n", r.RemoteAddr)

		if !r.URL.Query().Has("id") {
			http.Error(w, "Missing id parameter", http.StatusBadRequest)
			return
		}

		id_str := r.URL.Query().Get("id")

		id, err := strconv.Atoi(id_str)

		if err != nil {
			http.Error(w, "Invalid id parameter", http.StatusBadRequest)
		}

		idUint16 := uint16(id)

		if !r.URL.Query().Has("role") {
			http.Error(w, "Missing host parameter", http.StatusBadRequest)
			return
		}

		role := r.URL.Query().Get("role")

		if role != hostRole && role != clientRole {
			http.Error(w, "Invalid role parameter", http.StatusBadRequest)
		}

		if role == hostRole {
			HandleWebSocket(w, r, idUint16, hostEnum)
		} else {
			HandleWebSocket(w, r, idUint16, clientEnum)
		}

		return
		// Handle the WebSocket connection

	})

	err := http.ListenAndServe(fmt.Sprintf("%s:%d", allAddresses, port), nil)

	return err
}
