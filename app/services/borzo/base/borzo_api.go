package base

import "go_base_project/app/base"

type BorzoApi struct {
	BorzoConfig
}

func (api BorzoApi) Get(endpoint string) base.NetClient {
	return base.HttpService().Get().Url(api.BaseUrl()+endpoint).AddHeader("X-DV-Auth-Token", api.AuthToken())
}
