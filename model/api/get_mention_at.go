package api

import (
	"fmt"
	"net/http"

	"github.com/shimo-open/sdk-kit-go/model/common"
)

// 获取文件内容中所有的 at 人信息列表
// https://open.shimo.im/docs/06API-document/interface-description/collaborative-editing#%E8%8E%B7%E5%8F%96%E6%96%87%E4%BB%B6%E5%86%85%E5%AE%B9%E4%B8%AD%E6%89%80%E6%9C%89%E7%9A%84-at-%E4%BA%BA%E4%BF%A1%E6%81%AF%E5%88%97%E8%A1%A8

type GetMentionAtParams struct {
	Auth
	FileId string
}

type GetMentionAtRespBody struct {
	RawResponse
	MentionAtList []MentionAt `json:"mentionAtList"` // 根据指定文件 ID 获取的石墨文件中的所有 at 信息列表
}
type MentionAt struct {
	UserId string `json:"userId"` // at 提及的用户 ID
	AtGuid string `json:"atGuid"` // at 提及在文件中对应的位置标记
}

func NewGetMentionAtApi(params GetMentionAtParams) common.ApiConf {
	return common.ApiConf{
		Url:    fmt.Sprintf(Prefix+"/sdk/v2/collab-files/%s/mention-at-list", params.FileId),
		Method: http.MethodGet,
		ReqParams: common.ApiRequestParams{
			Headers: GetHeaderParams(params.Auth, nil),
		},
		RespParams: common.ApiResponseParams{
			SuccessHttpCode: http.StatusOK,
			Body:            GetMentionAtRespBody{},
		},
	}
}
