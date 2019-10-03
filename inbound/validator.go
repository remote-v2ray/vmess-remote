// +build !confonly

package inbound

import (
	"context"

	"github.com/apache/thrift/lib/go/thrift"
	"github.com/remote-v2ray/vmess-remote/inbound/gen-go/user"
	"v2ray.com/core/common/protocol"
	"v2ray.com/core/proxy/vmess"
)

type RemoteUserValidator struct {
	client *user.UserSvcClient
}

type RemoteUserValidatorOptions struct {
	RemoteURL string
}

func NewRemoteUserValidator(options RemoteUserValidatorOptions) *RemoteUserValidator {

	if options.RemoteURL == "" {
		panic("thrift remote url must be defined")
	}

	transport, _ := thrift.NewTHttpClient(options.RemoteURL)
	client := user.NewUserSvcClientFactory(transport, thrift.NewTJSONProtocolFactory())
	remoteUserValidator := &RemoteUserValidator{
		client: client,
	}
	return remoteUserValidator
}

func toRemoteMemoryUser(u *protocol.MemoryUser) (remoteUser *user.MemoryUser) {
	a := u.Account.(*vmess.MemoryAccount)
	account := &user.Account{
		ID:      a.ID.String(),
		AlterID: int32(len(a.AlterIDs)),
	}
	remoteUser = &user.MemoryUser{
		Email:   u.Email,
		Level:   int32(u.Level),
		Account: account,
	}
	return
}

func (v *RemoteUserValidator) Add(u *protocol.MemoryUser) (err error) {
	remoteUser := toRemoteMemoryUser(u)
	err = v.client.Add(context.Background(), remoteUser)
	return
}

func toLocalMemoryUser(u *user.MemoryUser) (user *protocol.MemoryUser, err error) {
	a := u.GetAccount()
	va := &vmess.Account{
		Id:      a.GetID(),
		AlterId: uint32(a.GetAlterID()),
	}
	ma, err := va.AsAccount()
	if err != nil {
		return
	}
	user = &protocol.MemoryUser{
		Email:   u.GetEmail(),
		Level:   uint32(u.GetLevel()),
		Account: ma,
	}
	return
}

func (v *RemoteUserValidator) Get(userHash []byte) (*protocol.MemoryUser, protocol.Timestamp, bool) {
	r, err := v.client.Get(context.Background(), userHash)
	if err != nil {
		return nil, 0, false
	}
	if found := r.GetFound(); !found {
		return nil, 0, true
	}
	u := r.GetUser()
	user, err := toLocalMemoryUser(u)
	if err != nil {
		return nil, 0, true
	}
	t := r.GetTime()
	return user, protocol.Timestamp(t), true
}

func (v *RemoteUserValidator) Remove(email string) bool {
	return false
}

// Close implements common.Closable.
func (v *RemoteUserValidator) Close() error {
	return nil
}
