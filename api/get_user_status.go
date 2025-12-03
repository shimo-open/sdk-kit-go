package api

import (
	"net/http"
	"strconv"

	"github.com/gotomicro/ego/client/ehttp"
)

// 获取用户列表和席位状态
// https://open.shimo.im/docs/06API-document/interface-description/system-interface#%E8%8E%B7%E5%8F%96%E7%94%A8%E6%88%B7%E5%88%97%E8%A1%A8%E5%92%8C%E5%B8%AD%E4%BD%8D%E7%8A%B6%E6%80%81

type GetUserAndStatusReq struct {
	Metadata
	Page, Size int
}

type GetUserAndStatusRes struct {
	rawRes
	UserId    string `json:"userId"`    // 用户 ID
	CreatedAt string `json:"createdAt"` // 创建时间
	Status    int    `json:"status"`    //  用户席位状态
}

func NewGetUserAndStatusApi(cli *ehttp.Component, ss SignatureSigner, params GetUserAndStatusReq) *APIConf {
	sign := ss.Sign(expire4m, ScopeSystem)
	return &APIConf{
		ss:     ss,
		Client: cli,
		URL:    "/sdk/v2/api/license/users",
		Method: http.MethodGet,
		ReqParams: APIRequestParams{
			Headers: addHeaders(params.Metadata, sign, nil),
			Query: map[string][]string{
				"page": {strconv.Itoa(params.Page)},
				"size": {strconv.Itoa(params.Size)},
			},
		},
		ResParams: APIResParams{
			SuccessHTTPCode: http.StatusOK,
			Body:            GetUserAndStatusRes{},
		},
	}
}
