package server

import (
	"fmt"
	"net"
	"net/http"
	"strconv"
)

const allAddresses = "0.0.0.0"

const (
	hostRole   = "host"
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

		if !r.URL.Query().Has("role") {
			http.Error(w, "Missing host parameter", http.StatusBadRequest)
			return
		}

		role := r.URL.Query().Get("role")

		if role != hostRole && role != clientRole {
			http.Error(w, "Invalid role parameter", http.StatusBadRequest)
		}

		if role == hostRole {
			handleWebSocket(w, r, uint16(id), hostEnum)
		} else {
			handleWebSocket(w, r, uint16(id), clientEnum)
		}

		return
		// Handle the WebSocket connection

	})

	ips, err := getIps()

	if err != nil {
		return err
	}

	ips_listening <- ips

	for _, ip := range ips {
		fmt.Printf("Listening on %s:%d\n", ip, port)
	}

	err = http.ListenAndServe(fmt.Sprintf("%s:%d", allAddresses, port), nil)

	return err
}

func getIps() ([]string, error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return nil, err
	}

	var ips []string

	for _, addr := range addrs {
		ips = append(ips, addr.String())
	}

	return ips, nil
}
