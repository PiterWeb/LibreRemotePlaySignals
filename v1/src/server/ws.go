package server

import (
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

type connectionT struct {
	hostChannel   chan []byte
	clientChannel chan []byte
}

const (
	hostEnum = iota
	clientEnum
)

var connectionsMu sync.Mutex
var connections = make(map[uint16]connectionT)

func handleWebSocket(w http.ResponseWriter, r *http.Request, ID uint16, role int) {
	// Implementation of handling WebSocket connections
	// This function will manage the WebSocket connection for the server
	// and handle incoming messages or events.

	upgrader := websocket.Upgrader{
		ReadBufferSize:  2048,
		WriteBufferSize: 2048,
	}

	conn, err := upgrader.Upgrade(w, r, nil)

	if err != nil {
		http.Error(w, "Failed to upgrade connection", http.StatusInternalServerError)
		return
	}

	defer conn.Close()

	connectionsMu.Lock()

	if _, ok := connections[ID]; !ok {
		connections[ID] = connectionT{
			hostChannel:   make(chan []byte),
			clientChannel: make(chan []byte),
		}
	}

	var wg sync.WaitGroup

	wg.Add(2)

	if role == hostEnum {
		go readMessages(&wg, conn, connections[ID].hostChannel)
		go writeMessages(&wg, conn, connections[ID].clientChannel)
	} else {
		go readMessages(&wg, conn, connections[ID].clientChannel)
		go writeMessages(&wg, conn, connections[ID].hostChannel)
	}

	connectionsMu.Unlock()

	wg.Wait()

}

func readMessages(wg *sync.WaitGroup, conn *websocket.Conn, channel chan<- []byte) {

	defer wg.Done()

	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			break
		}

		if len(msg) == 0 {
			continue
		}

		channel <- msg

	}
}
func writeMessages(wg *sync.WaitGroup, conn *websocket.Conn, channel <-chan []byte) {

	defer wg.Done()

	for msg := range channel {
		err := conn.WriteMessage(websocket.BinaryMessage, msg)
		if err != nil {
			break
		}
	}
}
