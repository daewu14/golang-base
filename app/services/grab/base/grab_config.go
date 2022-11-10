package base

import "os"

type GrabConfig struct {
}

func (config GrabConfig) BaseUrl() string {
	return os.Getenv("GRAB_EXPRESS_BASE_URL")
}

func (config GrabConfig) ClientID() string {
	return os.Getenv("GRAB_EXPRESS_CLIENT_ID")
}

func (config GrabConfig) ClientSecret() string {
	return os.Getenv("GRAB_EXPRESS_CLIENT_SECRET")
}

func (config GrabConfig) APIKey() string {
	return os.Getenv("GRAB_EXPRESS_API_KEY")
}

func (config GrabConfig) SuffixUrl() string {
	return os.Getenv("GRAB_EXPRESS_SUFFIX_URL")
}
