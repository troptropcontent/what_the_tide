package config

import "testing"

func TestRoot(t *testing.T) {
	t.Run("It should return the project root path", func(t *testing.T) {
		expected := "/workspaces/what_the_tide"
		if got := Root(); got != expected {
			t.Errorf("Root() = %v, want %v", got, expected)
		}
	})
}
