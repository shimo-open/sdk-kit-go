package api

import (
	"fmt"
	"net/http"

	"github.com/gotomicro/ego/client/ehttp"
)

// JudgeTableLockedReq contains parameters for checking spreadsheet lock status.
type JudgeTableLockedReq struct {
	Metadata
	// FileID is the spreadsheet file identifier.
	FileID string
}

// JudgeTableLockedRes contains the spreadsheet lock status.
type JudgeTableLockedRes struct {
	rawRes
	// LockedStatus is the spreadsheet lock status: 0 unlocked, 1 locked, 2 parse failed.
	// LockedStatus 是表格锁定状态：0 未锁定，1 已锁定，2 内容解析失败。
	LockedStatus int `json:"lockedStatus"`
}

// NewJudgeTableLockedApi creates the spreadsheet lock status request.
func NewJudgeTableLockedApi(cli *ehttp.Component, ss SignatureSigner, p JudgeTableLockedReq) *APIConf {
	sign := ss.Sign(ExpireLong, ScopeDefault)
	return &APIConf{
		ss:     ss,
		Client: cli,
		URL:    fmt.Sprintf("/sdk/v2/api/files/%s/judge-table-locked", p.FileID),
		Method: http.MethodGet,
		ReqParams: APIRequestParams{
			Headers: addHeaders(p.Metadata, sign, nil),
		},
		ResParams: APIResParams{
			SuccessHTTPCode: http.StatusOK,
			Body:            JudgeTableLockedRes{},
		},
	}
}
