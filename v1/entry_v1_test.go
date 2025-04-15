package LRPSignals

import (
	"sync"
	"testing"
)

func TestInitServer(t *testing.T) {
	
	// Test the InitServer function
	// This is a placeholder test and should be replaced with actual test logic
	port := uint16(8080)
	ips_listening := make(chan []string)

	go func () {
		err := InitServer(port, ips_listening)
		if err != nil {
			t.Errorf("InitServer failed: %v", err)
		}
	}()

	<-ips_listening

	// Close the channel after use
	close(ips_listening)

	server, err := Server("ws://localhost:8080/ws")

	if err != nil {
		t.Errorf("Server creation failed: %v", err)
	}

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		
		defer wg.Done()

		clientCode := ClientCodeT{
			Data: "test client code",
		}

		t.Log("Sending client code to server...")
		hostCode, err := SendClientCode(server, clientCode, 1)
		if err != nil {
			t.Fatalf("SendClientCode failed: %v", err)
		}
		t.Log("Client code sent successfully and received host code:", hostCode)
	}()

	hostCode := HostCodeT{
		Data: "test host code",
	}

	clientCode, err := ReceiveClientCode(server, 1)
	
	if err != nil {
		t.Fatalf("ReceiveClientCode failed: %v", err)
	}

	t.Log("Received client code successfully:", clientCode)
	t.Log("Sending host code to server...")

	err = SendHostCode(server, hostCode, 1)
	if err != nil {
		t.Fatalf("SendHostCode failed: %v", err)
	}

	t.Log("Host code sent successfully")

	wg.Wait()

}