package common

import (
	"bytes"
	"douyin/constants"
	"github.com/patrickmn/go-cache"
	"math/rand"
)

var c *cache.Cache

func init() {
	c = cache.New(constants.CacheTokenExpiration, constants.CacheCleanInterval)
}

func SetToken(token string, userID int64) {
	c.Set(token, userID, 0)
}

func CheckToken(token string, userID int64) error {
	if checkID, found := GetUserIDFromToken(token); !found || checkID != userID {
		return constants.AuthTokenFail
	}
	return nil
}

func GenerateToken() string {
	buf := bytes.Buffer{}
	buf.Grow(constants.TokenLength)
	for i := 0; i < constants.TokenLength; i++ {
		buf.WriteByte(constants.TokenCharacterDictionary[rand.Intn(constants.TokenDictionaryLength)])
	}
	return buf.String()
}

func GetUserIDFromToken(token string) (userID int64, found bool) {
	inter, found := c.Get(token)
	if !found {
		return 0, false
	}
	userID = inter.(int64)
	return userID, true
}
