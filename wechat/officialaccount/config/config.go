package config

import "go-knowledge/wechat/cache"

type Config struct {
	AppID          string `json:"app_id"`
	AppSecret      string `json:"app_secret"`
	Token          string `json:"token"`
	EncodingAESKey string `json:"encoding_aes_key"`
	Cache          cache.Cache
}
