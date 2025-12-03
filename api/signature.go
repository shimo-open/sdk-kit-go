package api

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/gotomicro/ego/core/econf"
)

// SignatureSigner is the interface for generating API request signatures.
// SignatureSigner 是用于生成 API 请求签名的接口。
type SignatureSigner interface {
	// Sign generates a signature with the given expiration and scope.
	// Sign 使用给定的过期时间和作用域生成签名。
	Sign(expireAfter time.Duration, scope Scope) string
}

// signatureSinger provides signature generation functionality
type signatureSinger struct {
	appid     string
	appsecret string
}

// NewSignatureSigner creates a new SignatureSigner instance with the given app credentials.
// NewSignatureSigner 使用给定的应用凭据创建新的 SignatureSigner 实例。
func NewSignatureSigner(appid, appsecret string) SignatureSigner {
	s := &signatureSinger{
		appid:     appid,
		appsecret: appsecret,
	}
	return s
}

// Sign generates a JWT signature with the given credentials and expiration policy
func (s *signatureSinger) Sign(expireAfter time.Duration, scope Scope) string {
	nowTime := time.Now()
	exp := nowTime.Add(expireAfter).Unix()
	return signJWT(s.appid, s.appsecret, exp, string(scope))
}

// ClientInfo holds application client credentials.
// ClientInfo 保存应用客户端凭据。
type ClientInfo struct {
	// AppID is the application identifier.
	// AppID 是应用标识符。
	AppID string `json:"appId"`
	// AppSecret is the application secret key.
	// AppSecret 是应用密钥。
	AppSecret string `json:"appSecret"`
}

// CustomClaims extends JWT standard claims with additional fields.
// CustomClaims 扩展 JWT 标准声明并添加额外字段。
type CustomClaims struct {
	*jwt.StandardClaims
	// Scope is the permission scope for the token.
	// Scope 是令牌的权限作用域。
	Scope string `json:"scope,omitempty"`
	// Version is the callback version.
	// Version 是回调版本。
	Version string `json:"version,omitempty"`
}

// signJWT generates a JWT token with the given parameters
func signJWT(kid, secret string, expires int64, scope string) (signature string) {
	var token *jwt.Token
	fmt.Printf("scope--------------->"+"%+v\n", scope)
	if scope != "" {
		token = jwt.NewWithClaims(
			jwt.SigningMethodHS256, &CustomClaims{
				StandardClaims: &jwt.StandardClaims{
					ExpiresAt: expires,
				},
				Scope:   scope,
				Version: econf.GetString("shimoSDK.callbackVersion"),
			},
		)
	} else {
		token = jwt.NewWithClaims(
			jwt.SigningMethodHS256, &CustomClaims{
				StandardClaims: &jwt.StandardClaims{
					ExpiresAt: expires,
				},
				Version: econf.GetString("shimoSDK.callbackVersion"),
			},
		)
	}

	token.Header["kid"] = kid
	tokenStr, err := token.SignedString([]byte(secret))
	if err != nil {
		panic(err)
	}

	return tokenStr
}

// UserClaims represents JWT claims for user authentication.
// UserClaims 表示用户认证的 JWT 声明。
type UserClaims struct {
	*jwt.StandardClaims
	// UserId is the user's unique identifier.
	// UserId 是用户的唯一标识符。
	UserId int64 `json:"userId"`
	// Mode is the authentication mode.
	// Mode 是认证模式。
	Mode string `json:"mode"`
}

// SignUserJWT generates a JWT token for user authentication.
// SignUserJWT 为用户认证生成 JWT 令牌。
func SignUserJWT(userId int64, expr ...time.Duration) string {
	var expires time.Duration
	if len(expr) > 0 {
		expires = expr[0]
	} else {
		expires = 24 * time.Hour
	}
	secret := econf.GetString("jwt.secret")
	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256, &UserClaims{
			StandardClaims: &jwt.StandardClaims{
				ExpiresAt: time.Now().Add(expires).Unix(),
			},
			UserId: userId,
		})
	tokenStr, err := token.SignedString([]byte(secret))
	if err != nil {
		panic(err)
	}

	return tokenStr
}

// SignUserJWTWithMode generates a user JWT token with a specific authentication mode.
// SignUserJWTWithMode 使用指定的认证模式生成用户 JWT 令牌。
func SignUserJWTWithMode(userId int64, mode string) string {
	secret := econf.GetString("jwt.secret")
	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256, &UserClaims{
			StandardClaims: &jwt.StandardClaims{
				ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
			},
			UserId: userId,
			Mode:   mode,
		})
	token.Header["kid"] = econf.GetString("shimoSDK.appId")
	tokenStr, err := token.SignedString([]byte(secret))
	if err != nil {
		panic(err)
	}

	return tokenStr
}

// SDKClaims represents JWT claims for SDK operations.
// SDKClaims 表示 SDK 操作的 JWT 声明。
type SDKClaims struct {
	*jwt.StandardClaims
	// FileId is the file identifier.
	// FileId 是文件标识符。
	FileId string `json:"fileId"`
	// UserId is the user identifier.
	// UserId 是用户标识符。
	UserId string `json:"userId"`
}
