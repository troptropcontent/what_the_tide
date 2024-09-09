package models

import (
	"testing"

	"github.com/troptropcontent/what_the_tide/config"
)

func TestPorts_Load(t *testing.T) {
	t.Run("It should load the ports into the struct", func(t *testing.T) {
		portConfig := PortsConfig{}
		err := portConfig.LoadFromJson(config.Root() + "/" + PortsConfigFile)
		expected := 1
		if err != nil {
			t.Errorf("Root() expected no error, got : %v", err)
		}
		if got := len(portConfig.Ports); got != expected {
			t.Errorf("Root() expected %d ports loaded, got: %d", expected, got)
		}
	})
}
