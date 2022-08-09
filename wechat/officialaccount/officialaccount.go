package officialaccount

import (
	"go-knowledge/wechat/credential"
	"go-knowledge/wechat/officialaccount/config"
	"go-knowledge/wechat/officialaccount/context"
)

type OfficialAccount struct {
	ctx *context.Context
}

func NewOfficialAccount(cfg *config.Config) *OfficialAccount {
	defaultAkHandle := credential.NewDefaultAccessToken(cfg.AppID, cfg.AppSecret, credential.CacheKeyOfficialAccountPrefix, cfg.Cache)
	ctx := &context.Context{
		Config:            cfg,
		AccessTokenHandle: defaultAkHandle,
	}
	return &OfficialAccount{
		ctx: ctx,
	}
}

func (officialAccount *OfficialAccount) SetAccessTokenHandle(accessTokenHandle credential.AccessTokenHandle) {
	officialAccount.ctx.AccessTokenHandle = accessTokenHandle
}

func (officialAccount *OfficialAccount) GetContext() *context.Context {
	return officialAccount.ctx
}
