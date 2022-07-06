package base

import "go_base_project/app/base"

type GosendApi struct {
	GosendConfig
}

func (api GosendApi) Get(endpoint string) base.NetClient {
	return base.HttpService().Get().Url(api.BaseUrl()+endpoint).AddHeader("Client-ID", api.ClientID()).AddHeader("Pass-Key", api.PassKey())
}

func (api GosendApi) Post(endpoint string) base.NetClient {
	return base.HttpService().Post().Url(api.BaseUrl()+endpoint).AddHeader("Client-ID", api.ClientID()).AddHeader("Pass-Key", api.PassKey())
}
