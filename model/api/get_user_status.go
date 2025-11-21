package api

import (
	"net/http"
	"strconv"

	"github.com/shimo-open/sdk-kit-go/model/common"
)

// 获取用户列表和席位状态
// https://open.shimo.im/docs/06API-document/interface-description/system-interface#%E8%8E%B7%E5%8F%96%E7%94%A8%E6%88%B7%E5%88%97%E8%A1%A8%E5%92%8C%E5%B8%AD%E4%BD%8D%E7%8A%B6%E6%80%81

type GetUserAndStatusParams struct {
	Auth
	Page, Size int
}

type GetUserAndStatusRespBody struct {
	UserId    string `json:"userId"`    // 用户 ID
	CreatedAt string `json:"createdAt"` // 创建时间
	Status    int    `json:"status"`    //  用户席位状态
}

func NewGetUserAndStatusApi(params GetUserAndStatusParams) common.ApiConf {
	return common.ApiConf{
		Url:    Prefix + "/sdk/v2/api/license/users",
		Method: http.MethodGet,
		ReqParams: common.ApiRequestParams{
			Headers: GetHeaderParams(params.Auth, nil),
			Query: map[string][]string{
				"page": {strconv.Itoa(params.Page)},
				"size": {strconv.Itoa(params.Size)},
			},
		},
		RespParams: common.ApiResponseParams{
			SuccessHttpCode: http.StatusOK,
			Body:            GetUserAndStatusRespBody{},
		},
	}
}
