package api

import (
	"fmt"
	"net/http"

	"github.com/gotomicro/ego/client/ehttp"
)

// GetCellsImageReq contains parameters for retrieving cell image URLs.
type GetCellsImageReq struct {
	Metadata
	// FileID is the spreadsheet file identifier.
	FileID string
	// SheetName is the target worksheet name.
	SheetName string
	// CellsPosition contains the requested cell coordinates in JSON format.
	CellsPosition string
}

// GetCellsImageRes contains the cell image response.
type GetCellsImageRes struct {
	rawRes
	// Images contains image URLs corresponding to the requested cell coordinates.
	// Images 是与请求单元格坐标一一对应的图片地址。
	Images []*string `json:"images"`
}

// NewGetCellsImageApi creates the cell image URL request.
func NewGetCellsImageApi(cli *ehttp.Component, ss SignatureSigner, p GetCellsImageReq) *APIConf {
	sign := ss.Sign(ExpireLong, ScopeDefault)
	return &APIConf{
		ss:     ss,
		Client: cli,
		URL:    fmt.Sprintf("/sdk/v2/api/files/%s/sheets/%s/cells-image", p.FileID, p.SheetName),
		Method: http.MethodGet,
		ReqParams: APIRequestParams{
			Headers: addHeaders(p.Metadata, sign, nil),
			Query: map[string][]string{
				"cellsPosition": {p.CellsPosition},
			},
		},
		ResParams: APIResParams{
			SuccessHTTPCode: http.StatusOK,
			Body:            GetCellsImageRes{},
		},
	}
}
