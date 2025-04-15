package signals

import (
	"fmt"

	"github.com/PiterWeb/LibreRemotePlaySignals/v1/src/types"
	"github.com/gorilla/websocket"
)

func ReceiveClientCode(s types.Server, ID uint16) (types.ClientCode, error) {
	// Implementation of receiving the client code from the server

	c, _, err := websocket.DefaultDialer.Dial(fmt.Sprintf("%s?role=host&id=%d",s.GetUrl(), ID), nil)

	if err != nil {
		return types.ClientCode{}, err
	}

	defer c.Close()

	// Process the message and convert it to a ClientCode
	clientCode := types.ClientCode{}
	
	err = c.ReadJSON(&clientCode)

	if err != nil {
		return types.ClientCode{}, err
	}

	return clientCode, nil
}