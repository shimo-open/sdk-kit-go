package api

import (
	"fmt"
	"net/http"

	"github.com/gotomicro/ego/client/ehttp"
)

// GetMentionAtReq contains parameters for getting all @mention information in a file.
// GetMentionAtReq 包含获取文件中所有 @提及信息的参数。
// API Doc: https://open.shimo.im/docs/06API-document/interface-description/collaborative-editing#%E8%8E%B7%E5%8F%96%E6%96%87%E4%BB%B6%E5%86%85%E5%AE%B9%E4%B8%AD%E6%89%80%E6%9C%89%E7%9A%84-at-%E4%BA%BA%E4%BF%A1%E6%81%AF%E5%88%97%E8%A1%A8
type GetMentionAtReq struct {
	Metadata
	// FileID is the unique identifier of the file.
	// FileID 是文件的唯一标识符。
	FileID string
}

// GetMentionAtRes contains the response for getting all @mention information in a file.
// GetMentionAtRes 包含获取文件中所有 @提及信息的响应。
type GetMentionAtRes struct {
	rawRes
	// MentionAtList is the list of all @mention information in the file.
	// MentionAtList 是文件中所有 @提及信息的列表。
	MentionAtList []MentionAt `json:"mentionAtList"`
}

// MentionAt represents a single @mention entry.
// MentionAt 表示单个 @提及条目。
type MentionAt struct {
	// UserId is the user ID being mentioned.
	// UserId 是被提及的用户 ID。
	UserId string `json:"userId"`
	// AtGuid is the position marker of the @mention in the file.
	// AtGuid 是 @提及在文件中的位置标记。
	AtGuid string `json:"atGuid"`
}

// NewGetMentionAtApi creates a new API config for getting all @mention information in a file.
// NewGetMentionAtApi 创建用于获取文件中所有 @提及信息的 API 配置。
func NewGetMentionAtApi(cli *ehttp.Component, ss SignatureSigner, params GetMentionAtReq) *APIConf {
	sign := ss.Sign(ExpireLong, ScopeDefault)
	return &APIConf{
		ss:     ss,
		Client: cli,
		URL:    fmt.Sprintf("/sdk/v2/collab-files/%s/mention-at-list", params.FileID),
		Method: http.MethodGet,
		ReqParams: APIRequestParams{
			Headers: addHeaders(params.Metadata, sign, nil),
		},
		ResParams: APIResParams{
			SuccessHTTPCode: http.StatusOK,
			Body:            GetMentionAtRes{},
		},
	}
}
