package config_test

import (
	"fmt"
	"latipe-notification-service/config"
	"os"
	"testing"
)

func TestInjectConfig(t *testing.T) {
	err := os.Setenv("cfgPath", "config")
	if err != nil {
		fmt.Printf("Error setting environment variable: %v\n", err)
		return
	}

	_, err = config.NewConfig()
	if err != nil {
		t.Error(err)
	}
}
