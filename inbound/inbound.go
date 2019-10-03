// +build !confonly

package inbound

import (
	"context"
	"log"
	"os"

	"v2ray.com/core/common"
	"v2ray.com/core/proxy/vmess/inbound"
)

func init() {
	// 取消 vmess inbound 的注册
	common.UnRegisterConfig((*inbound.Config)(nil))
	common.Must(common.RegisterConfig((*inbound.Config)(nil), func(ctx context.Context, config interface{}) (interface{}, error) {
		endpoint := os.Getenv("ThriftUserValidatorEndpoint")
		if endpoint == "" {
			log.Fatal("env ThriftUserValidatorEndpoint is required")
		}
		remoteUserValidator := NewRemoteUserValidator(RemoteUserValidatorOptions{
			RemoteURL: endpoint,
		})
		options := inbound.HandlerOptions{
			Clients: remoteUserValidator,
		}
		// 返回修改了用户验证模块后的 vmess inbound
		return inbound.New(ctx, config.(*inbound.Config), options)
	}))
}
