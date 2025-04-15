package types

import (
	net_url "net/url"
)

type Server struct {
	port uint16
	url  string
}

func (s *Server) SetPort(port uint16) {
	s.port = port
}

func (s Server) GetPort() uint16 {
	return s.port
}

// SetUrl sets the URL of the server.
// The URL must be a valid URL string.
// If the URL is not valid, an error will be returned.
func (s *Server) SetUrl(url_str string) error {
	url, err := net_url.Parse(url_str)

	if err != nil {
		return err
	}

	s.url = url.String()
	return nil
}

func (s Server) GetUrl() string {
	return s.url
}
