### 简介

需要指定环境变量 `RemoteUserValidatorAPI` 才能运行

接口定义文件是 [`./inbound/user.thrift`](./inbound/user.thrift)

### 使用方法

#### 第一步: 克隆仓库

`git clone -b features/remote https://github.com/remote-v2ray/v2ray-core` (因为有些 pr 尚未合并到主分支只能使用这个仓库的 remote 分支)

#### 第二步: 修改模块配置文件 `v2ray-core/main/distro/all/all.go`

因为 `v2ray` 是模块化配置的, 所以我们可以很容易地替换其中一个模块.
将其中的 `vmess/inbound` 模块替换成本仓库实现的模块 `gitlab.com/shy-v2ray/vmess-remote/inbound`

```golang
  _ "v2ray.com/core/proxy/vmess/inbound"
```

```golang
	// _ "v2ray.com/core/proxy/vmess/inbound"
	_ "github.com/remote-v2ray/vmess-remote/inbound"
```

### 第三步: 构建 `v2ray`

```sh
cd v2ray-core/main
# 下载依赖到 vendor 文件夹里
go mod vendor
# 构建 v2ray
CGO_ENABLED=0 go build -mod=vendor -ldflags "-s -w" -o v2ray
```

#### 第五步: 服务端实现远程验证接口并启动 v2ray

这个时候所有 `vmess` 类型的 `inbound` 用户验证都会通过远程接口来验证.

服务端暂未完成, 只是测试了用户验证是可以通过和拒绝的

```
ThriftUserValidatorEndpoint=http://127.0.0.1/v2ray/user v2ray -config config.json
```
