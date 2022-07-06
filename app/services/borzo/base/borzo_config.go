package base

import (
	"os"
)

type BorzoConfig struct {
}

func (config BorzoConfig) BaseUrl() string {
	return os.Getenv("BORZO_BASE_URL")
}

func (config BorzoConfig) AuthToken() string {
	return os.Getenv("BORZO_AUTH_TOKEN")
}
