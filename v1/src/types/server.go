package types

import (
	net_url "net/url"
)

type Server struct {
	url  string
}

type ServerOptions struct {
	Port uint16
	// Secure struct {
	// 	Domains []string
	// 	Enabled bool
	// }
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
