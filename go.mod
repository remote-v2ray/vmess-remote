module github.com/remote-v2ray/vmess-remote

go 1.13

require (
	github.com/apache/thrift v0.13.0
	v2ray.com/core v4.19.1+incompatible
)

replace v2ray.com/core => ./modules/v2ray
