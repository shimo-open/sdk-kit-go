package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gotomicro/ego/client/ehttp"
)

// GetUserAndStatusReq contains parameters for getting user list and seat status.
// GetUserAndStatusReq 包含获取用户列表和席位状态的参数。
// API Doc: https://open.shimo.im/docs/06API-document/interface-description/system-interface#%E8%8E%B7%E5%8F%96%E7%94%A8%E6%88%B7%E5%88%97%E8%A1%A8%E5%92%8C%E5%B8%AD%E4%BD%8D%E7%8A%B6%E6%80%81
type GetUserAndStatusReq struct {
	Metadata
	// Page is the page number.
	// Page 是页码。
	Page int
	// Size is the page size.
	// Size 是每页大小。
	Size int
}

// GetUserAndStatusRes contains the response for getting user list and seat status.
// GetUserAndStatusRes 包含获取用户列表和席位状态的响应。
type GetUserAndStatusRes struct {
	rawRes
	Users []UserStatus `json:"-"`
}

// UnmarshalJSON implements json.Unmarshaler interface.
// The API returns an array directly, so we need custom unmarshaling.
// UnmarshalJSON 实现 json.Unmarshaler 接口。
// API 直接返回数组，因此需要自定义反序列化。
func (r *GetUserAndStatusRes) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.Users)
}

type UserStatus struct {
	// UserId is the user ID.
	// UserId 是用户 ID。
	UserId string `json:"userId"`
	// CreatedAt is the creation time.
	// CreatedAt 是创建时间。
	CreatedAt string `json:"createdAt"`
	// Status is the user seat status.
	// Status 是用户席位状态。
	Status int `json:"status"`
}

// NewGetUserAndStatusApi creates a new API config for getting user list and seat status.
// NewGetUserAndStatusApi 创建用于获取用户列表和席位状态的 API 配置。
func NewGetUserAndStatusApi(cli *ehttp.Component, ss SignatureSigner, params GetUserAndStatusReq) *APIConf {
	sign := ss.Sign(ExpireShort, ScopeSystem)
	return &APIConf{
		ss:     ss,
		Client: cli,
		URL:    "/sdk/v2/api/license/users",
		Method: http.MethodGet,
		ReqParams: APIRequestParams{
			Headers: addHeaders(params.Metadata, sign, nil),
			Query: map[string][]string{
				"page": {strconv.Itoa(params.Page)},
				"size": {strconv.Itoa(params.Size)},
			},
		},
		ResParams: APIResParams{
			SuccessHTTPCode: http.StatusOK,
			Body:            GetUserAndStatusRes{},
		},
	}
}
