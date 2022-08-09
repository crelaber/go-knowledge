package context

import (
	"go-knowledge/wechat/credential"
	"go-knowledge/wechat/officialaccount/config"
)

type Context struct {
	*config.Config
	credential.AccessTokenHandle
}
