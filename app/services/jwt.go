package services

import (
	"chatbot-backend/global"
	"chatbot-backend/utils"
	"context"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type jwtService struct {
}

var JwtService = new(jwtService)

type JwtUser interface {
	GetUid() string
}

// CustomClaims 自定义 Claims
type CustomClaims struct {
	jwt.StandardClaims
}

const (
	TokenType = "bearer"
)

type TokenOutPut struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
	TokenType   string `json:"token_type"`
}

func (jwtService *jwtService) getBlackListKey(tokenStr string) string {
    return "jwt_black_list:" + utils.MD5([]byte(tokenStr))
}

// JoinBlackList 
func (jwtService *jwtService) JoinBlackList(token *jwt.Token) (err error) {
    nowUnix := time.Now().Unix()
    timer := time.Duration(token.Claims.(*CustomClaims).ExpiresAt - nowUnix) * time.Second
   
    err = global.App.Redis.SetNX(context.Background(), jwtService.getBlackListKey(token.Raw), nowUnix, timer).Err()
    return
}

// IsInBlacklist
func (jwtService *jwtService) IsInBlacklist(tokenStr string) bool {
    joinUnixStr, err := global.App.Redis.Get(context.Background(), jwtService.getBlackListKey(tokenStr)).Result()
    joinUnix, err := strconv.ParseInt(joinUnixStr, 10, 64)
    if joinUnixStr == "" || err != nil {
        return false
    }

    if time.Now().Unix()-joinUnix < global.App.Config.Jwt.JwtBlacklistGracePeriod {
        return false
    }
    return true
}

func (jwtService *jwtService) CreateToken(user JwtUser) (tokenData TokenOutPut, err error, token *jwt.Token) {
	token = jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		CustomClaims{
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: time.Now().Unix() + global.App.Config.Jwt.JwtTTL,
				Id:        user.GetUid(),
				Issuer:    "web",
				NotBefore: time.Now().Unix() - 1000,
			},
		},
	)

	tokenStr, err := token.SignedString([]byte(global.App.Config.Jwt.Secret))

	tokenData = TokenOutPut{
		tokenStr,
		int(global.App.Config.Jwt.JwtTTL),
		TokenType,
	}
	return
}
