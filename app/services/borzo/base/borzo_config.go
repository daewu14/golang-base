package base

import (
	"fmt"
	"os"
)

type BorzoConfig struct {
}

func (config BorzoConfig) BaseUrl() string {
	fmt.Println("ini enve", os.Getenv("BORZO_BASE_URL"))
	return os.Getenv("BORZO_BASE_URL")
}

func (config BorzoConfig) AuthToken() string {
	return os.Getenv("BORZO_AUTH_TOKEN")
}
