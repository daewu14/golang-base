package base

import "go_base_project/app/base"

type GrabApi struct {
	GrabConfig
}

func (api GrabApi) Get(endpoint string) base.NetClient {
	return base.HttpService().Get().Url(api.BaseUrl()+endpoint).AddHeader("Content-Type", "application/json")
}

func (api GrabApi) Post(endpoint string) base.NetClient {
	return base.HttpService().Post().Url(api.BaseUrl()+api.SuffixUrl()+endpoint).AddHeader("Content-Type", "application/json")
}

func (api GrabApi) PostGetAccessToken(endpoint string) base.NetClient {
	return base.HttpService().Post().Url(api.BaseUrl()+endpoint).AddHeader("Content-Type", "application/json")
}
