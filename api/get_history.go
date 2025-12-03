package api

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gotomicro/ego/client/ehttp"
)

// GetHistoryListReq contains parameters for getting document history list.
// GetHistoryListReq 包含获取文档历史列表的参数。
// API Doc: https://open.shimo.im/docs/06API-document/interface-description/collaborative-editing#doc-sidebar-info
type GetHistoryListReq struct {
	Metadata
	// FileID is the unique identifier of the file.
	// FileID 是文件的唯一标识符。
	FileID string
	// PageSize is the number of items per page.
	// PageSize 是每页的项目数。
	PageSize int
	// Count is the total count offset.
	// Count 是总计数偏移量。
	Count int
	// HistoryType is the type of history (1: operation history, 2: edit history).
	// HistoryType 是历史类型（1：操作历史，2：编辑历史）。
	HistoryType int
}

// GetHistoryListRes contains the response for getting document history list.
// GetHistoryListRes 包含获取文档历史列表的响应。
type GetHistoryListRes struct {
	rawRes
	// Histories is the array of sidebar history items.
	// Histories 是侧边栏历史数组。
	Histories []History `json:"histories"`
	// IsLastPage indicates whether this is the last page.
	// IsLastPage 表示是否是最后一页。
	IsLastPage bool `json:"isLastPage"`
	// Limit is the page size.
	// Limit 是分页大小。
	Limit int `json:"limit"`
	// Users is the mapping of user IDs to usernames.
	// Users 是用户 ID 到用户名的映射。
	Users interface{} `json:"users"`
}

// History represents a single history entry.
// History 表示单个历史记录条目。
type History struct {
	// Content is the collaborative file format data.
	// Content 是协作文件格式数据。
	Content string `json:"content"`
	// CreatedAt is the creation time of this history entry.
	// CreatedAt 是本条历史记录的创建时间。
	CreatedAt string `json:"createdAt"`
	// HistoryType is the type (1: operation history, 2: edit history).
	// HistoryType 是类型（1：操作历史，2：编辑历史）。
	HistoryType int `json:"historyType"`
	// Id is the history entry ID.
	// Id 是历史记录 ID。
	Id string `json:"id"`
	// Name is the history entry name.
	// Name 是历史记录名称。
	Name string `json:"name"`
	// UpdateAt is the last update time.
	// UpdateAt 是最后更新时间。
	UpdateAt string `json:"updateAt"`
	// UserId is the user ID(s), multiple IDs separated by comma.
	// UserId 是用户 ID，多个 ID 以逗号分隔。
	UserId string `json:"userId"`
}

// NewGetHistoryListApi creates a new API config for getting document history list.
// NewGetHistoryListApi 创建用于获取文档历史列表的 API 配置。
func NewGetHistoryListApi(cli *ehttp.Component, ss SignatureSigner, params GetHistoryListReq) *APIConf {
	sign := ss.Sign(ExpireLong, ScopeDefault)
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
