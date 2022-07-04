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
	return "C5360B9A1595660CEABD1FFE53B31BF6A2822D48"
}
