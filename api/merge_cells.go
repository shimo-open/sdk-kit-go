package api

import (
	"fmt"
	"net/http"

	"github.com/gotomicro/ego/client/ehttp"
)

// MergeCellsReq contains parameters for merging spreadsheet cells.
type MergeCellsReq struct {
	Metadata
	// FileID is the spreadsheet file identifier.
	FileID string
	// SheetName is the target worksheet name.
	SheetName string `json:"sheetName"`
	// Cell is the top-left cell of the merge range.
	Cell string `json:"cell"`
	// Rowspan is the number of rows to merge.
	Rowspan int `json:"rowspan"`
	// ColSpan is the number of columns to merge.
	ColSpan int `json:"colSpan"`
}

// MergeCellsRes is the response for merging spreadsheet cells.
type MergeCellsRes struct{ rawRes }

// NewMergeCellsApi creates the spreadsheet cell merge request.
func NewMergeCellsApi(cli *ehttp.Component, ss SignatureSigner, p MergeCellsReq) *APIConf {
	body := struct {
		SheetName string `json:"sheetName"`
		Cell      string `json:"cell"`
		Rowspan   int    `json:"rowspan"`
		ColSpan   int    `json:"colSpan"`
	}{p.SheetName, p.Cell, p.Rowspan, p.ColSpan}
	sign := ss.Sign(ExpireLong, ScopeDefault)
	extra := map[string]string{"Content-Type": "application/json"}
	return &APIConf{
		ss:     ss,
		Client: cli,
		URL:    fmt.Sprintf("/sdk/v2/api/files/%s/sheets/merge-cells", p.FileID),
		Method: http.MethodPost,
		ReqParams: APIRequestParams{
			Headers: addHeaders(p.Metadata, sign, extra),
			Body:    body,
		},
		ResParams: APIResParams{
			SuccessHTTPCode: http.StatusNoContent,
			Body:            MergeCellsRes{},
		},
	}
}
