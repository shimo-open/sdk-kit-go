package api

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/shimo-open/sdk-kit-go/model/common"
)

// 获取历史列表
// https://open.shimo.im/docs/06API-document/interface-description/collaborative-editing#doc-sidebar-info

type GetHistoryListParams struct {
	Auth
	FileId                       string
	PageSize, Count, HistoryType int
}

type GetHistoryListRespBody struct {
	RawResponse
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

func NewGetHistoryListApi(params GetHistoryListParams) common.ApiConf {
	return common.ApiConf{
		Url:    fmt.Sprintf(Prefix+"/sdk/v2/collab-files/%s/doc-sidebar-info", params.FileId),
		Method: http.MethodGet,
		ReqParams: common.ApiRequestParams{
			Headers: GetHeaderParams(params.Auth, nil),
			Query: map[string][]string{
				"pageSize":    {strconv.Itoa(params.PageSize)},
				"count":       {strconv.Itoa(params.Count)},
				"historyType": {strconv.Itoa(params.HistoryType)},
			},
		},
		RespParams: common.ApiResponseParams{
			SuccessHttpCode: http.StatusOK,
			Body:            GetHistoryListRespBody{},
		},
	}
}
