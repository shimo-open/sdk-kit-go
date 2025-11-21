package api

import (
	"fmt"
	"net/http"

	"github.com/shimo-open/sdk-kit-go/model/common"
)

// 获取版本列表
// https://open.shimo.im/docs/06API-document/interface-description/collaborative-editing#%E8%8E%B7%E5%8F%96%E7%89%88%E6%9C%AC%E5%88%97%E8%A1%A8

type GetRevisionListParams struct {
	Auth
	FileId string
}

type GetRevisionListRespBody struct {
	Id           int    `json:"id"`           // 版本 ID
	Label        string `json:"label"`        // 版本 Label
	Title        string `json:"title"`        // 标题
	DocHistoryId string `json:"docHistoryId"` // 对应侧边栏历史 ID
	CreatedAt    string `json:"createdAt"`    // 侧边栏历史创建时间
	UpdatedAt    string `json:"updatedAt"`    // 侧边栏历史更新时间
	User         User   `json:"user"`         // 服务商用户
}

type User struct {
	Id   string `json:"id"`   // 服务商用户 ID
	Name string `json:"name"` // 服务商用户 用户名
}

func NewGetRevisionListApi(params GetRevisionListParams) common.ApiConf {
	return common.ApiConf{
		Url:    fmt.Sprintf(Prefix+"/sdk/v2/collab-files/%s/revisions", params.FileId),
		Method: http.MethodGet,
		ReqParams: common.ApiRequestParams{
			Headers: GetHeaderParams(params.Auth, nil),
		},
		RespParams: common.ApiResponseParams{
			SuccessHttpCode: http.StatusOK,
			Body:            GetRevisionListRespBody{},
		},
	}
}
