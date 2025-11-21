package api

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/shimo-open/sdk-kit-go/model/common"
)

// 删除表格行
// https://open.shimo.im/docs/06API-document/interface-description/collaborative-editing#%E5%88%A0%E9%99%A4%E8%A1%A8%E6%A0%BC%E8%A1%8C

type DeleteTableRowParams struct {
	Auth
	FileId, SheetName string
	Index, Count      int
}

func NewDeleteTableRowApi(params DeleteTableRowParams) common.ApiConf {
	extra := map[string]string{
		"Content-Type": "application/json",
	}
	return common.ApiConf{
		Url:    fmt.Sprintf(Prefix+"/sdk/v2/api/files/%s/sheets/%s/rows/%d", params.FileId, params.SheetName, params.Index),
		Method: http.MethodDelete,
		ReqParams: common.ApiRequestParams{
			Headers: GetHeaderParams(params.Auth, extra),
			Query: map[string][]string{
				"count": {strconv.Itoa(params.Count)},
			},
		},
		RespParams: common.ApiResponseParams{
			SuccessHttpCode: http.StatusNoContent,
		},
	}
}
