package config

import "os"

func Host() string {

	host := os.Getenv("HOST_WALLET")

	if host == "" {
		return "0.0.0.0"
	}
	return host
}

func Port() string {

	port := os.Getenv("PORT_WALLET")

	if port == "" {
		return "8080"
	}
	return port

}
