package api

import (
	"fmt"
	"net/http"

	"github.com/gotomicro/ego/client/ehttp"
)

// SetCellAttrsReq contains parameters for updating spreadsheet cell attributes.
type SetCellAttrsReq struct {
	Metadata
	// FileID is the spreadsheet file identifier.
	FileID string
	// Range is the worksheet cell range to update.
	Range string `json:"range"`
	// Attrs contains the attributes to apply to the range.
	Attrs map[string]interface{} `json:"attrs"`
}

// SetCellAttrsRes is the response for updating cell attributes.
type SetCellAttrsRes struct{ rawRes }

// NewSetCellAttrsApi creates the spreadsheet cell attributes request.
func NewSetCellAttrsApi(cli *ehttp.Component, ss SignatureSigner, p SetCellAttrsReq) *APIConf {
	body := struct {
		Range string                 `json:"range"`
		Attrs map[string]interface{} `json:"attrs"`
	}{p.Range, p.Attrs}
	sign := ss.Sign(ExpireLong, ScopeDefault)
	extra := map[string]string{"Content-Type": "application/json"}
	return &APIConf{
		ss:     ss,
		Client: cli,
		URL:    fmt.Sprintf("/sdk/v2/api/files/%s/sheets/attrs", p.FileID),
		Method: http.MethodPost,
		ReqParams: APIRequestParams{
			Headers: addHeaders(p.Metadata, sign, extra),
			Body:    body,
		},
		ResParams: APIResParams{
			SuccessHTTPCode: http.StatusOK,
			Body:            SetCellAttrsRes{},
		},
	}
}
