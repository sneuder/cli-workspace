package wsDependency

import (
	"strings"
	"workspace/internal/docker/network"
)

// networks

func Networks(networks string) {
	if networks == "" {
		return
	}

	collectionNetwork := strings.Split(networks, " ")

	for _, networkName := range collectionNetwork {
		if network.Exists(networkName) {
			continue
		}

		network.Create(networkName)
	}
}

// databases
