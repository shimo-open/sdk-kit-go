package api

import (
	"fmt"
	"net/http"

	"github.com/shimo-open/sdk-kit-go/model/common"
)

// 获取App详情
// https://open.shimo.im/docs/06API-document/interface-description/system-interface#%E8%8E%B7%E5%8F%96app%E8%AF%A6%E6%83%85

type GetAppDetailParams struct {
	Auth
	AppId string
}

type GetAppDetailRespBody struct {
	RawResponse
	AppName            string   `json:"appName"`            // 应用名
	AvailableFileTypes []string `json:"availableFileTypes"` // 可用石墨套件列表
	PremiumFileTypes   []string `json:"premiumFileTypes"`   // 增值套件列表
	ActivatedUserCount int      `json:"activatedUserCount"` // 已激活席位用户数
	UserCount          int      `json:"userCount"`          // 用户数总数，包含已激活、已禁用和未使用的用户总数，仅“已激活 ”数量占用席位。
	MemberLimit        int      `json:"memberLimit"`        // license 席位限制用户数，即“已激活 ”用户数量最大限制
	ValidFrom          string   `json:"validFrom"`          // license 生效时间
	ValidUntil         string   `json:"validUntil"`         // license 到期时间
	EndpointUrl        string   `json:"endpointUrl"`        // 接入方回调地址
}

func NewGetAppDetailApi(params GetAppDetailParams) common.ApiConf {
	return common.ApiConf{
		Url:    fmt.Sprintf(Prefix+"/sdk/v2/api/license/apps/%s", params.AppId),
		Method: http.MethodGet,
		ReqParams: common.ApiRequestParams{
			Headers: GetHeaderParams(params.Auth, nil),
		},
		RespParams: common.ApiResponseParams{
			SuccessHttpCode: http.StatusOK,
			Body:            GetAppDetailRespBody{},
		},
	}
}
