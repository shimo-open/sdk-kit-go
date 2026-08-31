package api

import (
	"fmt"
	"net/http"

	"github.com/gotomicro/ego/client/ehttp"
)

// GetMentionAtPositionReq contains parameters for locating a spreadsheet mention.
type GetMentionAtPositionReq struct {
	Metadata
	// FileID is the spreadsheet file identifier.
	FileID string
	// AtGuid is the mention identifier.
	AtGuid string
}

// GetMentionAtPositionRes contains the mention position response.
type GetMentionAtPositionRes struct {
	rawRes
	// PositionInfo contains the worksheet position of the mention.
	// PositionInfo 包含 @ 信息所在的工作表位置。
	PositionInfo MentionAtPosition `json:"positionInfo"`
}

// MentionAtPosition represents the worksheet position of a mention.
// MentionAtPosition 表示 @ 信息所在的工作表位置。
type MentionAtPosition struct {
	SheetID   string `json:"sheetId"`
	SheetName string `json:"sheetName"`
	Row       int    `json:"row"`
	Col       int    `json:"col"`
	Position  string `json:"position"`
}

// NewGetMentionAtPositionApi creates the spreadsheet mention position request.
func NewGetMentionAtPositionApi(cli *ehttp.Component, ss SignatureSigner, p GetMentionAtPositionReq) *APIConf {
	sign := ss.Sign(ExpireLong, ScopeDefault)
	return &APIConf{
		ss:     ss,
		Client: cli,
		URL:    fmt.Sprintf("/sdk/v2/api/files/%s/sheets/mention-at-position", p.FileID),
		Method: http.MethodGet,
		ReqParams: APIRequestParams{
			Headers: addHeaders(p.Metadata, sign, nil),
			Query: map[string][]string{
				"atGuid": {p.AtGuid},
			},
		},
		ResParams: APIResParams{
			SuccessHTTPCode: http.StatusOK,
			Body:            GetMentionAtPositionRes{},
		},
	}
}
