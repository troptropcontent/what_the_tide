package models

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

const PortsConfigFile = "config/ports.json"

type Port struct {
	Name string `json:"name"`
	Id   int    `json:"id"`
}

type PortsConfig struct {
	Ports []Port `json:"ports"`
}

func (portsConfig *PortsConfig) LoadFromJson(path string) error {
	jsonFile, err := os.Open(path)
	if err != nil {
		return err
	}
	byteValue, _ := io.ReadAll(jsonFile)
	json.Unmarshal(byteValue, portsConfig)

	return nil
}

func (portsConfig *PortsConfig) FindPort(id int) (requestedPort Port, err error) {
	// TODO : Use memoization here
	for _, port := range portsConfig.Ports {
		if port.Id == id {
			requestedPort = port

			break
		}
	}

	if requestedPort.Id == 0 {
		return requestedPort, fmt.Errorf("no ports found for this ID : %d", id)
	}

	return requestedPort, nil
}
