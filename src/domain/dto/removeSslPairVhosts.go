package dto

import "github.com/speedianet/os/src/domain/valueObject"

type RemoveSslPairVhosts struct {
	SslPairId    valueObject.SslId  `json:"sslPairId"`
	VirtualHosts []valueObject.Fqdn `json:"virtualHosts"`
}

func NewRemoveSslPairVhosts(
	sslPairId valueObject.SslId,
	virtualHosts []valueObject.Fqdn,
) RemoveSslPairVhosts {
	return RemoveSslPairVhosts{
		SslPairId:    sslPairId,
		VirtualHosts: virtualHosts,
	}
}