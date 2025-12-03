package api

import (
	"fmt"
	"net/http"

	"github.com/gotomicro/ego/client/ehttp"
)

// GetAppDetailReq contains parameters for getting app details.
// GetAppDetailReq 包含获取应用详情的参数。
// API Doc: https://open.shimo.im/docs/06API-document/interface-description/system-interface#%E8%8E%B7%E5%8F%96app%E8%AF%A6%E6%83%85
type GetAppDetailReq struct {
	Metadata
	// AppID is the application identifier.
	// AppID 是应用标识符。
	AppID string
}

// GetAppDetailRes contains the response for getting app details.
// GetAppDetailRes 包含获取应用详情的响应。
type GetAppDetailRes struct {
	rawRes
	// AppName is the application name.
	// AppName 是应用名称。
	AppName string `json:"appName"`
	// AvailableFileTypes is the list of available Shimo suite types.
	// AvailableFileTypes 是可用石墨套件类型列表。
	AvailableFileTypes []string `json:"availableFileTypes"`
	// PremiumFileTypes is the list of premium suite types.
	// PremiumFileTypes 是增值套件类型列表。
	PremiumFileTypes []string `json:"premiumFileTypes"`
	// ActivatedUserCount is the number of activated seat users.
	// ActivatedUserCount 是已激活席位用户数。
	ActivatedUserCount int `json:"activatedUserCount"`
	// UserCount is the total number of users (activated, disabled, and unused).
	// UserCount 是用户总数（包含已激活、已禁用和未使用的用户）。
	UserCount int `json:"userCount"`
	// MemberLimit is the license seat limit.
	// MemberLimit 是许可证席位限制用户数。
	MemberLimit int `json:"memberLimit"`
	// ValidFrom is the license effective date.
	// ValidFrom 是许可证生效时间。
	ValidFrom string `json:"validFrom"`
	// ValidUntil is the license expiration date.
	// ValidUntil 是许可证到期时间。
	ValidUntil string `json:"validUntil"`
	// EndpointUrl is the callback URL.
	// EndpointUrl 是接入方回调地址。
	EndpointUrl string `json:"endpointUrl"`
}

// NewGetAppDetailApi creates a new API config for getting app details.
// NewGetAppDetailApi 创建用于获取应用详情的 API 配置。
func NewGetAppDetailApi(cli *ehttp.Component, ss SignatureSigner, params GetAppDetailReq) *APIConf {
	sign := ss.Sign(ExpireShort, ScopeSystem)
	return &APIConf{
		ss:     ss,
		Client: cli,
		URL:    fmt.Sprintf("/sdk/v2/api/license/apps/%s", params.AppID),
		Method: http.MethodGet,
		ReqParams: APIRequestParams{
			Headers: addHeaders(params.Metadata, sign, nil),
		},
		ResParams: APIResParams{
			SuccessHTTPCode: http.StatusOK,
			Body:            GetAppDetailRes{},
		},
	}
}
