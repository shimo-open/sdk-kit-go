package api

import (
	"fmt"
	"net/http"

	"github.com/gotomicro/ego/client/ehttp"
)

// 获取版本列表
// https://open.shimo.im/docs/06API-document/interface-description/collaborative-editing#%E8%8E%B7%E5%8F%96%E7%89%88%E6%9C%AC%E5%88%97%E8%A1%A8

type GetRevisionListReq struct {
	Metadata
	FileID string
}

type GetRevisionListRes struct {
	Revisions []GetRevisionListRes `json:"revisions"`
	rawRes
}

type GetRevisionListRevision struct {
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

func NewGetRevisionListApi(cli *ehttp.Component, ss SignatureSigner, params GetRevisionListReq) *APIConf {
	sign := ss.Sign(expire4m, ScopeDefault)
	return &APIConf{
		ss:     ss,
		Client: cli,
		URL:    fmt.Sprintf("/sdk/v2/collab-files/%s/revisions", params.FileID),
		Method: http.MethodGet,
		ReqParams: APIRequestParams{
			Headers: addHeaders(params.Metadata, sign, nil),
		},
		ResParams: APIResParams{
			SuccessHTTPCode: http.StatusOK,
			Body:            GetRevisionListRes{},
		},
	}
}
