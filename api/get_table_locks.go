package api

import (
	"fmt"
	"net/http"

	"github.com/gotomicro/ego/client/ehttp"
)

// GetTableLocksReq contains parameters for retrieving spreadsheet locks.
type GetTableLocksReq struct {
	Metadata
	// FileID is the spreadsheet file identifier.
	FileID string
}

// GetTableLocksRes contains spreadsheet lock information.
type GetTableLocksRes struct {
	rawRes
	Locks []TableSheetLock `json:"locks"`
}

type TableSheetLock struct {
	SheetID string      `json:"sheetId"`
	Locks   []TableLock `json:"locks"`
}

type TableLock struct {
	Type              string            `json:"type"`
	LockerID          int64             `json:"lockerId"`
	VisitorPermission string            `json:"visitorPermission"`
	Users             map[string]string `json:"users"`
	Departments       map[string]string `json:"departments"`
	Groups            map[string]string `json:"groups"`
	Ranges            []string          `json:"ranges"`
}

// NewGetTableLocksApi creates the spreadsheet locks request.
func NewGetTableLocksApi(cli *ehttp.Component, ss SignatureSigner, p GetTableLocksReq) *APIConf {
	sign := ss.Sign(ExpireLong, ScopeDefault)
	return &APIConf{
		ss:     ss,
		Client: cli,
		URL:    fmt.Sprintf("/sdk/v2/api/files/%s/sheets/locks", p.FileID),
		Method: http.MethodGet,
		ReqParams: APIRequestParams{
			Headers: addHeaders(p.Metadata, sign, nil),
		},
		ResParams: APIResParams{
			SuccessHTTPCode: http.StatusOK,
			Body:            GetTableLocksRes{},
		},
	}
}
