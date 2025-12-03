package api

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gotomicro/ego/client/ehttp"
)

// 获取历史列表
// https://open.shimo.im/docs/06API-document/interface-description/collaborative-editing#doc-sidebar-info

type GetHistoryListReq struct {
	Metadata
	FileID                       string
	PageSize, Count, HistoryType int
}

type GetHistoryListRes struct {
	rawRes
	Histories  []History   `json:"histories"`  // 侧边栏历史数组
	IsLastPage bool        `json:"isLastPage"` // 是否最后一页
	Limit      int         `json:"limit"`      // 分页大小
	Users      interface{} `json:"users"`      // 接入方用户 ID 对应的用户名映射
}

type History struct {
	Content     string `json:"content"`     // 协作文件格式数据
	CreatedAt   string `json:"createdAt"`   // 本条侧边栏历史创建时间
	HistoryType int    `json:"historyType"` // 侧边栏历史类型，1 为操作历史，2 为编辑产生
	Id          string `json:"id"`          // 侧边栏历史 ID
	Name        string `json:"name"`        // 侧边栏历史名称
	UpdateAt    string `json:"updateAt"`    // 侧边栏历史最后更新时间
	UserId      string `json:"userId"`      // 服务商用户 ID，可能有多个，以英文逗号 "," 分隔
}

func NewGetHistoryListApi(cli *ehttp.Component, ss SignatureSigner, params GetHistoryListReq) *APIConf {
	sign := ss.Sign(expire4m, ScopeDefault)
	return &APIConf{
		ss:     ss,
		Client: cli,
		URL:    fmt.Sprintf("/sdk/v2/collab-files/%s/doc-sidebar-info", params.FileID),
		Method: http.MethodGet,
		ReqParams: APIRequestParams{
			Headers: addHeaders(params.Metadata, sign, nil),
			Query: map[string][]string{
				"pageSize":    {strconv.Itoa(params.PageSize)},
				"count":       {strconv.Itoa(params.Count)},
				"historyType": {strconv.Itoa(params.HistoryType)},
			},
		},
		ResParams: APIResParams{
			SuccessHTTPCode: http.StatusOK,
			Body:            GetHistoryListRes{},
		},
	}
}
