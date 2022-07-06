package base

import "os"

type GosendConfig struct {
}

func (config GosendConfig) BaseUrl() string {
	return os.Getenv("GOSEND_BASE_URL")
}

func (config GosendConfig) PassKey() string {
	return os.Getenv("GOSEND_PASS_KEY")
}

func (config GosendConfig) ClientID() string {
	return os.Getenv("GOSEND_CLIENT_ID")
}
