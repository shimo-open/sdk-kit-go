package api

import (
	"fmt"
	"net/http"

	"github.com/gotomicro/ego/client/ehttp"
)

// UploadCellImageReq contains parameters for uploading an image to a cell.
type UploadCellImageReq struct {
	Metadata
	// FileID is the spreadsheet file identifier.
	FileID string
	// SheetName is the target worksheet name.
	SheetName string `json:"sheetName"`
	// Cell is the target cell address.
	Cell string `json:"cell"`
	// Base64 is the image content encoded as base64.
	Base64 string `json:"base64,omitempty"`
	// SrcURL is the source URL of the image.
	SrcURL string `json:"srcUrl,omitempty"`
}

// UploadCellImageRes is the response for uploading a cell image.
type UploadCellImageRes struct{ rawRes }

// NewUploadCellImageApi creates the cell image upload request.
func NewUploadCellImageApi(cli *ehttp.Component, ss SignatureSigner, p UploadCellImageReq) *APIConf {
	body := struct {
		SheetName string `json:"sheetName"`
		Cell      string `json:"cell"`
		Base64    string `json:"base64,omitempty"`
		SrcURL    string `json:"srcUrl,omitempty"`
	}{p.SheetName, p.Cell, p.Base64, p.SrcURL}
	sign := ss.Sign(ExpireLong, ScopeDefault)
	extra := map[string]string{"Content-Type": "application/json"}
	return &APIConf{
		ss:     ss,
		Client: cli,
		URL:    fmt.Sprintf("/sdk/v2/api/files/%s/sheets/upload-cell-image", p.FileID),
		Method: http.MethodPost,
		ReqParams: APIRequestParams{
			Headers: addHeaders(p.Metadata, sign, extra),
			Body:    body,
		},
		ResParams: APIResParams{
			SuccessHTTPCode: http.StatusNoContent,
			Body:            UploadCellImageRes{},
		},
	}
}
