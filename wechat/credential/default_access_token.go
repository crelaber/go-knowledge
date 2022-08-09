package credential

import (
	"encoding/json"
	"fmt"
	"go-knowledge/wechat/cache"
	"go-knowledge/wechat/util"
	"sync"
	"time"
)

const (
	accessTokenUrl = "https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=%s&secret=%s"
	//企业微信的token获取地址
	workAccessTokenURL = "https://qyapi.weixin.qq.com/cgi-bin/gettoken?corpid=%s&corpsecret=%s"
	// CacheKeyOfficialAccountPrefix 微信公众号cache key前缀
	CacheKeyOfficialAccountPrefix = "go_wechat_official_account_"
)

type DefaultAccessToken struct {
	appID           string
	appSecret       string
	cacheKeyPrefix  string
	cache           cache.Cache
	accessTokenLock *sync.Mutex
}

func NewDefaultAccessToken(appID, appSecret, cacheKeyPrefix string, cache cache.Cache) AccessTokenHandle {
	if cache == nil {
		panic("cache is need")
	}
	return &DefaultAccessToken{
		appID:           appID,
		appSecret:       appSecret,
		cache:           cache,
		cacheKeyPrefix:  cacheKeyPrefix,
		accessTokenLock: new(sync.Mutex),
	}
}

type ResAccessToken struct {
	util.CommonErr

	AccessToken string `json:"access_token"`
	ExpireIn    int64  `json:"expire_in"`
}

func (ak *DefaultAccessToken) GetAccessToken() (accessToken string, err error) {
	accessTokenCacheKey := fmt.Sprintf("%s_access_token_%s", ak.cacheKeyPrefix, ak.appID)
	if val := ak.cache.Get(accessTokenCacheKey); val != nil {
		return val.(string), nil
	}
	ak.accessTokenLock.Lock()
	defer ak.accessTokenLock.Unlock()

	if val := ak.cache.Get(accessTokenCacheKey); val != nil {
		return val.(string), nil
	}

	var resAccessToken ResAccessToken
	resAccessToken, err = GetTokenFromServer(fmt.Sprintf(accessTokenUrl, ak.appID, ak.appSecret))
	if err != nil {
		return
	}

	expires := resAccessToken.ExpireIn - 1500
	err = ak.cache.Set(accessTokenCacheKey, resAccessToken.AccessToken, time.Duration(expires)*time.Second)
	if err != nil {
		return
	}

	accessToken = resAccessToken.AccessToken
	return
}

func GetTokenFromServer(url string) (resAccessToken ResAccessToken, err error) {
	var body []byte
	body, err = util.HTTPGet(url)
	if err != nil {
		return
	}
	err = json.Unmarshal(body, &resAccessToken)
	if err != nil {
		return
	}
	if resAccessToken.ErrCode != 0 {
		err = fmt.Errorf("get access_token error : errcode = %v, errmsg = %v", resAccessToken.ErrCode, resAccessToken.ErrMsg)
	}
	return
}

type WorkAccessToken struct {
	CorpID          string
	CorpSecret      string
	cacheKeyPrefix  string
	cache           cache.Cache
	accessTokenLock *sync.Mutex
}

func NewWorkAccessToken(corpID, corpSecret, cacheKeyPrefix string, cache cache.Cache) AccessTokenHandle {
	if cache == nil {
		panic("cache not exist")
	}

	return &WorkAccessToken{
		CorpID:          corpID,
		CorpSecret:      corpSecret,
		cache:           cache,
		cacheKeyPrefix:  cacheKeyPrefix,
		accessTokenLock: new(sync.Mutex),
	}
}

func (ak WorkAccessToken) GetAccessToken() (accessToken string, err error) {
	ak.accessTokenLock.Lock()
	defer ak.accessTokenLock.Unlock()
	accessTokenCacheKey := fmt.Sprintf("%s_access_token_%s", ak.cacheKeyPrefix, ak.CorpID)
	val := ak.cache.Get(accessTokenCacheKey)
	if val != nil {
		accessToken = val.(string)
		return
	}

	var resAccessToken ResAccessToken
	resAccessToken, err = GetTokenFromServer(fmt.Sprintf(workAccessTokenURL, ak.CorpID, ak.CorpSecret))
	if err != nil {
		return
	}
	accessToken = resAccessToken.AccessToken
	return
}
