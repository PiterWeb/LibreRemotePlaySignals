package signals

import (
	"fmt"
	"log"

	"github.com/PiterWeb/LibreRemotePlaySignals/v1/src/types"
	"github.com/gorilla/websocket"
)

func SendClientCode(s types.Server, client_code types.ClientCode, ID uint16) (types.HostCode, error) {
	// Implementation of sending the client code to the server
	// and receiving the host code back.
	
	connUrl := fmt.Sprintf("%s?role=client&id=%d",s.GetUrl(), ID)

	log.Printf("Connecting to server at %s\n", connUrl)

	c, _, err := websocket.DefaultDialer.Dial(connUrl, nil)

	if err != nil {
		return types.HostCode{}, err
	}

	defer c.Close()
	
	err = c.WriteJSON(client_code)

	if err != nil {
		return types.HostCode{}, err
	}
	
	// Process the message and convert it to a HostCode
	hostCode := types.HostCode{}

	err = c.ReadJSON(&hostCode)

	if err != nil {
		return types.HostCode{}, err
	}

	return hostCode, nil
}

func SendHostCode(s types.Server, host_code types.HostCode, ID uint16) error {
	// Implementation of sending the host code to the server.
	
	connUrl := fmt.Sprintf("%s?role=host&id=%d", s.GetUrl(), ID)

	log.Printf("Connecting to server at %s\n", connUrl)

	c, _, err := websocket.DefaultDialer.Dial(connUrl, nil)

	if err != nil {
		return err
	}

	defer c.Close()
	
	err = c.WriteJSON(host_code)

	if err != nil {
		return err
	}
	
	return nil
}
