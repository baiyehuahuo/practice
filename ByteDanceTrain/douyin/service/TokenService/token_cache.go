package TokenService

import (
	"bytes"
	"douyin/constants"
	"douyin/model/dyerror"
	"github.com/patrickmn/go-cache"
	"math/rand"
)

var c *cache.Cache

func init() {
	c = cache.New(constants.CacheTokenExpiration, constants.CacheCleanInterval)
	c.Set("gyo7VknhMe5oQf28rVJm6mXoNzsdgVMa", 2, cache.DefaultExpiration) // use for test, 2 is test.TestUserID
}

func SetToken(token string, userID int64) {
	c.Set(token, userID, cache.DefaultExpiration)
}

func GenerateToken() string {
	buf := bytes.Buffer{}
	buf.Grow(constants.TokenLength)
	for i := 0; i < constants.TokenLength; i++ {
		buf.WriteByte(constants.TokenCharacterDictionary[rand.Intn(constants.TokenDictionaryLength)])
	}
	return buf.String()
}

func GetUserIDFromToken(token string) (userID int64, err *dyerror.DouyinError) {
	inter, found := c.Get(token)
	if !found {
		return 0, dyerror.AuthTokenFailError // userID forever not zero
	}
	userID = inter.(int64)
	return userID, nil
}

func CheckToken(token string, userID int64) *dyerror.DouyinError {
	tokenID, err := GetUserIDFromToken(token)
	if err != nil || userID != tokenID {
		return dyerror.AuthTokenFailError
	}
	return nil
}
