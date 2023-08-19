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
	c.Set("gyo7VknhMe5oQf28rVJm6mXoNzsdgVMa", int64(2), cache.DefaultExpiration) // use for test, 2 is test.TestUserID
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

func GetUserIDFromToken(token string) (userID int64, err error) {
	inter, found := c.Get(token)
	if !found {
		return 0, dyerror.AuthTokenFailError // userID forever not zero
	}
	userID = inter.(int64)
	return userID, nil
}

func CheckToken(token string, userID int64) error {
	tokenID, dyerr := GetUserIDFromToken(token)
	if dyerr != nil || userID != tokenID {
		return dyerror.AuthTokenFailError
	}
	return nil
}
