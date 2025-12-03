package api

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/gotomicro/ego/core/econf"
)

type TokenSigner interface {
	Sign(subjects ...any) string
}

type SignatureSigner interface {
	Sign(expireAfter time.Duration, scope Scope) string
}

// signatureSinger provides signature generation functionality
type signatureSinger struct {
	appid     string
	appsecret string
}

// NewSignatureSigner initializes a new signatureSinger instance
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

// ClientInfo holds application client credentials
type ClientInfo struct {
	AppID     string `json:"appId"`
	AppSecret string `json:"appSecret"`
}

// CustomClaims extends JWT standard claims with additional fields
type CustomClaims struct {
	*jwt.StandardClaims
	Scope   string `json:"scope,omitempty"`
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

// UserClaims represents JWT claims for user authentication
type UserClaims struct {
	*jwt.StandardClaims
	UserId int64  `json:"userId"`
	Mode   string `json:"mode"`
}

// SignUserJWT issues a user token
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

// SignUserJWTWithMode generates a user JWT token with a specific mode
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

// SDKClaims represents JWT claims for SDK operations
type SDKClaims struct {
	*jwt.StandardClaims
	FileId string `json:"fileId"`
	UserId string `json:"userId"`
}
