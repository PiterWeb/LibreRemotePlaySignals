package signals

import "github.com/PiterWeb/LibreRemotePlaySignals/v1/src/types"

func SendClientCode(s types.Server, client_code types.ClientCode, ID uint16) (types.HostCode, error) {
	// Implementation of sending the client code to the server
	// and receiving the host code back.
	return types.HostCode{}, nil
}

func SendHostCode(s types.Server, host_code types.HostCode, ID uint16) error {
	// Implementation of sending the host code to the server.
	return nil
}
