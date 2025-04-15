package LBRSignals

import "github.com/PiterWeb/LibreRemotePlaySignals/v1/types"

/*
InitServer initializes a HTTP + WS server on the given port.
The ips_listening channel is used to notify the outside world about the IPs listening
on the server.
The server will listen on all available IPs.

This function is used in the CLI or in the LibreRemotePlay Host APP to start the server.
The server will be started in a goroutine.
*/
func InitServer(port int, ips_listening chan<- []string) error {
	return nil
}

/*
Server creates a new server instance with the given port and URL.
The server instance is needed to send and receive client and host codes.
*/
func Server(port uint32, url string) (types.Server, error) {
	s := types.Server{}
	(&s).SetPort(port)
	err := (&s).SetUrl(url)

	if err != nil {
		return types.Server{}, err
	}

	return s, nil
}

/*
SendClientCode sends the client code to the server with a given connection ID
and returns the host code with the same connection ID.
*/
func SendClientCode(s types.Server, client_code types.ClientCode, ID int) (types.HostCode, error) {

	return types.HostCode{}, nil

}

/*
SendHostCode sends the host code to the server with a given connection ID.
*/
func SendHostCode(s types.Server, host_code types.HostCode, ID int) error {

	return nil

}

/*
ReceiveClientCode receives the client code from the server with a given connection ID.
*/
func ReceiveClientCode(s types.Server, ID int) (types.ClientCode, error) {

	return types.ClientCode{}, nil

}
