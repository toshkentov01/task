package utils

import (
	"fmt"

	"github.com/toshkentov01/task/api_gateway/config"
)

// ConnectionURLBuilder func for building URL connection.
func ConnectionURLBuilder(n string) (string, error) {
	// Define URL to connection.
	conf := config.Config()
	var url string

	// Switch given names.
	switch n {
	case "fiber":
		// URL for Fiber connection.
		url = fmt.Sprintf(
			":%d",
			conf.ServerPort,
		)
	default:
		// Return error message.
		return "", fmt.Errorf("connection name '%v' is not supported", n)
	}

	// Return connection URL.
	return url, nil
}
