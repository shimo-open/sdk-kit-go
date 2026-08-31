package api

import (
	"fmt"
	"net/http"

	"github.com/gotomicro/ego/client/ehttp"
)

// InsertTableRowsReq contains parameters for inserting spreadsheet rows.
type InsertTableRowsReq struct {
	Metadata
	// FileID is the spreadsheet file identifier.
	FileID string
	// SheetName is the target worksheet name.
	SheetName string
	// Index is the zero-based insertion row index.
	Index int
	// Count is the number of rows to insert.
	Count int
}

// InsertTableRowsRes is the response for inserting spreadsheet rows.
type InsertTableRowsRes struct{ rawRes }

// NewInsertTableRowsApi creates the spreadsheet row insertion request.
func NewInsertTableRowsApi(cli *ehttp.Component, ss SignatureSigner, p InsertTableRowsReq) *APIConf {
	sign := ss.Sign(ExpireLong, ScopeDefault)
	return &APIConf{
		ss:     ss,
		Client: cli,
		URL:    fmt.Sprintf("/sdk/v2/api/files/%s/sheets/%s/rows/%d", p.FileID, p.SheetName, p.Index),
		Method: http.MethodPost,
		ReqParams: APIRequestParams{
			Headers: addHeaders(p.Metadata, sign, nil),
			Query: map[string][]string{
				"count": {fmt.Sprintf("%d", p.Count)},
			},
		},
		ResParams: APIResParams{
			SuccessHTTPCode: http.StatusNoContent,
			Body:            InsertTableRowsRes{},
		},
	}
}
