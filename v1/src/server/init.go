package server

import (
	// "crypto/tls"
	"fmt"
	"log"
	"net"
	"net/http"
	"strconv"
	"time"

	"github.com/PiterWeb/LibreRemotePlaySignals/v1/src/types"
	// "golang.org/x/crypto/acme/autocert"
)

const allAddresses = "0.0.0.0"

const (
	hostRole   = "host"
	clientRole = "client"
)

func Init(options types.ServerOptions, ips_listening chan<- []string) error {

	serveMux := http.NewServeMux()

	serveMux.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {

		// Handle WebSocket connection
		log.Printf("WebSocket connection established from %s\n", r.RemoteAddr)

		if !r.URL.Query().Has("id") {
			http.Error(w, "Missing id parameter", http.StatusBadRequest)
			return
		}

		id_str := r.URL.Query().Get("id")

		id, err := strconv.Atoi(id_str)

		if err != nil {
			http.Error(w, "Invalid id parameter", http.StatusBadRequest)
			return
		}

		if !r.URL.Query().Has("role") {
			http.Error(w, "Missing host parameter", http.StatusBadRequest)
			return
		}

		role := r.URL.Query().Get("role")

		if role != hostRole && role != clientRole {
			http.Error(w, "Invalid role parameter", http.StatusBadRequest)
			return
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
		log.Printf("Listening on %s:%d\n", ip, options.Port)
	}

	// autocertManager := autocert.Manager{
	// 	Prompt:     autocert.AcceptTOS,
	// 	HostPolicy: autocert.HostWhitelist(options.Secure.Domains...),
	// }

	// tlsConfig := &tls.Config{
	// 	GetCertificate:           autocertManager.GetCertificate,
	// 	PreferServerCipherSuites: true,
	// 	CurvePreferences:         []tls.CurveID{tls.X25519, tls.CurveP256},
	// }

	server := &http.Server{
		Addr: fmt.Sprintf("%s:%d", allAddresses, options.Port),
		// TLSConfig: tlsConfig,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		Handler:      serveMux,
	}

	err = server.ListenAndServe()

	return err
}

func getIps() ([]string, error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return nil, err
	}

	var ips []string

	for _, addr := range addrs {
		if ipnet, ok := addr.(*net.IPNet); ok {
			ips = append(ips, ipnet.IP.String())
		}
	}

	return ips, nil
}
