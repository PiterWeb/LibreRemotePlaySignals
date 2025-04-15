/*
LibreRemotePlaySignals is a Go library used in the LibreRemotePlay project.

It provides a set of functions and types to handle signals and communication
between the LibreRemotePlay client and server.

The library is designed to be used in the LibreRemotePlay APP and CLI.
*/
package LRPSignals

import (
	"github.com/PiterWeb/LibreRemotePlaySignals/v1/src/server"
	"github.com/PiterWeb/LibreRemotePlaySignals/v1/src/signals"
	"github.com/PiterWeb/LibreRemotePlaySignals/v1/src/types"
)

type (
	// ClientCodeT represents the client code sent to the server.
	ClientCodeT = types.ClientCode
	// HostCodeT represents the host code sent to the server.
	HostCodeT = types.HostCode
	// ServerT represents the server instance used to send and receive codes.
	ServerT = types.Server
)

/*
InitServer initializes a HTTP + WS server on the given port.
The ips_listening channel is used to notify the outside world about the IPs listening
on the server.
The server will listen on all available IPs.

This function is used in the CLI or in the LibreRemotePlay Host APP to start the server.
The server will be started in a goroutine.
*/
func InitServer(port uint16, ips_listening chan<- []string) error {
	return server.Init(port, ips_listening)
}

/*
Server creates a new server instance with the given port and URL.
The server instance is needed to send and receive client and host codes.
*/
func Server(url string) (ServerT, error) {
	return server.Server(url)
}

/*
SendClientCode sends the client code to the server with a given connection ID
and returns the host code with the same connection ID.
*/
func SendClientCode(s ServerT, client_code ClientCodeT, ID uint16) (HostCodeT, error) {
	return signals.SendClientCode(s, client_code, ID)
}

/*
SendHostCode sends the host code to the server with a given connection ID.
*/
func SendHostCode(s ServerT, host_code HostCodeT, ID uint16) error {
	return signals.SendHostCode(s, host_code, ID)
}

/*
ReceiveClientCode receives the client code from the server with a given connection ID.
*/
func ReceiveClientCode(s ServerT, ID uint16) (ClientCodeT, error) {
	return signals.ReceiveClientCode(s, ID)
}
