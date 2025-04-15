package server

import "github.com/PiterWeb/LibreRemotePlaySignals/v1/src/types"

func Server(url string) (types.Server, error) {
	s := types.Server{}
	err := (&s).SetUrl(url)

	if err != nil {
		return types.Server{}, err
	}

	return s, nil
}