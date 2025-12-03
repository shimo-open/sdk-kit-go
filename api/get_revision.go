package api

import (
	"fmt"
	"net/http"

	"github.com/gotomicro/ego/client/ehttp"
)

// GetRevisionListReq contains parameters for getting document revision list.
// GetRevisionListReq 包含获取文档版本列表的参数。
// API Doc: https://open.shimo.im/docs/06API-document/interface-description/collaborative-editing#%E8%8E%B7%E5%8F%96%E7%89%88%E6%9C%AC%E5%88%97%E8%A1%A8
type GetRevisionListReq struct {
	Metadata
	// FileID is the unique identifier of the file.
	// FileID 是文件的唯一标识符。
	FileID string
}

// GetRevisionListRes contains the response for getting document revision list.
// GetRevisionListRes 包含获取文档版本列表的响应。
type GetRevisionListRes struct {
	// Revisions is the list of revision entries.
	// Revisions 是版本条目列表。
	Revisions []GetRevisionListRes `json:"revisions"`
	rawRes
}

// GetRevisionListRevision represents a single revision entry.
// GetRevisionListRevision 表示单个版本条目。
type GetRevisionListRevision struct {
	// Id is the revision ID.
	// Id 是版本 ID。
	Id int `json:"id"`
	// Label is the revision label.
	// Label 是版本标签。
	Label string `json:"label"`
	// Title is the revision title.
	// Title 是版本标题。
	Title string `json:"title"`
	// DocHistoryId is the corresponding sidebar history ID.
	// DocHistoryId 是对应的侧边栏历史 ID。
	DocHistoryId string `json:"docHistoryId"`
	// CreatedAt is the creation time.
	// CreatedAt 是创建时间。
	CreatedAt string `json:"createdAt"`
	// UpdatedAt is the update time.
	// UpdatedAt 是更新时间。
	UpdatedAt string `json:"updatedAt"`
	// User is the user who created the revision.
	// User 是创建版本的用户。
	User User `json:"user"`
}

// User represents a user in the system.
// User 表示系统中的用户。
type User struct {
	// Id is the user ID.
	// Id 是用户 ID。
	Id string `json:"id"`
	// Name is the username.
	// Name 是用户名。
	Name string `json:"name"`
}

// NewGetRevisionListApi creates a new API config for getting document revision list.
// NewGetRevisionListApi 创建用于获取文档版本列表的 API 配置。
func NewGetRevisionListApi(cli *ehttp.Component, ss SignatureSigner, params GetRevisionListReq) *APIConf {
	sign := ss.Sign(ExpireLong, ScopeDefault)
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
