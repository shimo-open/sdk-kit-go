package api

import (
	"fmt"
	"net/http"

	"github.com/gotomicro/ego/client/ehttp"
)

// 获取App详情
// https://open.shimo.im/docs/06API-document/interface-description/system-interface#%E8%8E%B7%E5%8F%96app%E8%AF%A6%E6%83%85

type GetAppDetailReq struct {
	Metadata
	AppID string
}

type GetAppDetailRes struct {
	rawRes
	AppName            string   `json:"appName"`            // 应用名
	AvailableFileTypes []string `json:"availableFileTypes"` // 可用石墨套件列表
	PremiumFileTypes   []string `json:"premiumFileTypes"`   // 增值套件列表
	ActivatedUserCount int      `json:"activatedUserCount"` // 已激活席位用户数
	UserCount          int      `json:"userCount"`          // 用户数总数，包含已激活、已禁用和未使用的用户总数，仅"已激活 "数量占用席位。
	MemberLimit        int      `json:"memberLimit"`        // license 席位限制用户数，即"已激活 "用户数量最大限制
	ValidFrom          string   `json:"validFrom"`          // license 生效时间
	ValidUntil         string   `json:"validUntil"`         // license 到期时间
	EndpointUrl        string   `json:"endpointUrl"`        // 接入方回调地址
}

func NewGetAppDetailApi(cli *ehttp.Component, ss SignatureSigner, params GetAppDetailReq) *APIConf {
	sign := ss.Sign(expire4m, ScopeSystem)
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
