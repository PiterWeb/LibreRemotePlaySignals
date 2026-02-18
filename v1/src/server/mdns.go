package server

import (
	"context"
	"log"
	"net"
	"os"
	"strings"
	"sync/atomic"
	"time"

	"github.com/pion/mdns/v2"
	"golang.org/x/net/ipv4"
	"golang.org/x/net/ipv6"
)

var usedLocalName atomic.Value

func announceMDNS(localName string) error {
	
	mdnsConfig := &mdns.Config{}
	
	if strings.TrimSpace(localName) == ""  {
		mdnsConfig.LocalNames = []string{"libreremoteplay-easyconnect-server.local"}
	} else {
		mdnsConfig.LocalNames = []string{localName}
	}
	
	used, err := isLocalBeingUsed(mdnsConfig.LocalNames[0])
	
	if err != nil {
		return err
	}
	
	if used {
		log.Println("Localname already in use, mdns canceled")
		return nil
	}
	
	addr4, err := net.ResolveUDPAddr("udp4", mdns.DefaultAddressIPv4)
	if err != nil {
		return err
	}

	addr6, err := net.ResolveUDPAddr("udp6", mdns.DefaultAddressIPv6)
	if err != nil {
		return err
	}

	l4, err := net.ListenUDP("udp4", addr4)
	if err != nil {
		return err
	}

	l6, err := net.ListenUDP("udp6", addr6)
	if err != nil {
		return err
	}
	
	setMDNSLocalUsed(mdnsConfig.LocalNames[0])
	
	_, err = mdns.Server(
		ipv4.NewPacketConn(l4),
		ipv6.NewPacketConn(l6),
		mdnsConfig,
	)
	
	if err != nil {
		return err
	}
	
	return nil
}

func isLocalBeingUsed(localName string) (bool, error) {
	var useV4, useV6 bool
	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "-v4only":
			useV4 = true
			useV6 = false
		case "-v6only":
			useV4 = false
			useV6 = true
		default:
			useV4 = true
			useV6 = true
		}
	} else {
		useV4 = true
		useV6 = true
	}

	var packetConnV4 *ipv4.PacketConn
	if useV4 {
		addr4, err := net.ResolveUDPAddr("udp4", mdns.DefaultAddressIPv4)
		if err != nil {
			return false, err
		}

		l4, err := net.ListenUDP("udp4", addr4)
		if err != nil {
			return false, err
		}

		packetConnV4 = ipv4.NewPacketConn(l4)
	}

	var packetConnV6 *ipv6.PacketConn
	if useV6 {
		addr6, err := net.ResolveUDPAddr("udp6", mdns.DefaultAddressIPv6)
		if err != nil {
			return false, err
		}

		l6, err := net.ListenUDP("udp6", addr6)
		if err != nil {
			return false, err
		}

		packetConnV6 = ipv6.NewPacketConn(l6)
	}

	server, err := mdns.Server(packetConnV4, packetConnV6, &mdns.Config{})
	if err != nil {
		return false, err
	}

	ctx, cancelFn := context.WithDeadline(context.Background(),time.Now().Add(time.Second * 3))
	
	defer cancelFn()
	
	// perform query
	_, _, err = server.QueryAddr(ctx, localName)
	
	if err != nil {
		return false, nil
	}
	
	return true, nil
	
}

func setMDNSLocalUsed(localName string) {
	usedLocalName.Store(localName)
}

func GetMDNSLocalUsed() string {
	localName, ok := usedLocalName.Load().(string)
	
	if !ok {
		return ""
	}
	
	return localName
}