package api

import (
	"context"
	"fmt"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/gotomicro/ego/client/ehttp"
)

const (
	// HeaderShimoToken is the header name for authentication token
	HeaderShimoToken = "X-Shimo-Token"
	// HeaderShimoSignature is the header name for request signature
	HeaderShimoSignature = "X-Shimo-Signature"
	// HeaderShimoCredentialType is the header name for credential type
	HeaderShimoCredentialType = "X-Shimo-Credential-Type"
	// HeaderShimoSdkEvent is the header name for SDK events
	HeaderShimoSdkEvent = "X-Shimo-Sdk-Event"
	// HeaderWebofficeToken is the new shimo token
	HeaderWebofficeToken = "X-Weboffice-Token"
	// HeaderWebofficeUserUuid controller user info cache
	HeaderWebofficeUserUuid = "X-Weboffice-User-Uuid"

	// ME represents the current user in API requests
	ME = "me"
	// ANONYMOUS is the user ID for anonymous users
	ANONYMOUS = -1
	// ANONYMOUSTOKEN is the token used for anonymous users
	ANONYMOUSTOKEN = "pseudonymoustoken"
)

const (
	expire4m  = 4 * time.Minute
	expire24h = 24 * time.Hour
)

type Scope string

const (
	ScopeDefault Scope = ""
	ScopeSystem  Scope = "license"
)

const (
	ApiBase           = "/sdk/v2"
	ApiIframeAssets   = "/sdk/v2/shimo-assets/static-uploader/sdk-iframe-assets.json"
	ApiAiAssets       = "/sdk/v2/api/assets"
	ApiCloudFilesPage = "/sdk/v2/api/cloud-files/%s/page"
)

// Metadata contains authentication information.
// Metadata 包含认证信息。
type Metadata struct {
	// ShimoToken token. 令牌。
	ShimoToken string
	// WebofficeToken token, new shimo token
	WebofficeToken string
	// User's unique identifier. 用户 UUID。
	WebofficeUserUuid string
}

// addHeaders builds HTTP request headers from auth and extra parameters.
// addHeaders 从认证信息和额外参数构建 HTTP 请求头。
func addHeaders(md Metadata, sig string, extraHeaders map[string]string) map[string][]string {
	if md.WebofficeToken == "" {
		md.WebofficeToken = md.ShimoToken
	}
	// 基础 header
	params := map[string][]string{
		HeaderShimoSignature:    {sig},
		HeaderShimoToken:        {md.ShimoToken},
		HeaderWebofficeToken:    {md.WebofficeToken},
		HeaderWebofficeUserUuid: {md.WebofficeUserUuid},
	}
	// 添加额外 header
	for k, v := range extraHeaders {
		params[k] = []string{v}
	}
	return params
}

// Lang represents a language type.
// Lang 表示语言类型。
type Lang string

// Language type constants.
// 语言类型常量。
const (
	LangZhCN Lang = "zh-CN" // Simplified Chinese. 简体中文。
	LangEn   Lang = "en"    // English. 英语。
	LangJp   Lang = "jp"    // Japanese. 日语。
)

// String returns the string representation of the language.
// String 返回语言的字符串表示。
func (l Lang) String() string {
	return string(l)
}

// CollabFileType represents the type of a collaborative file.
// CollabFileType 表示协同文件类型。
type CollabFileType string

// Collaborative file type constants.
// 协同文件类型常量。
const (
	CollabFileTypeDocument     CollabFileType = "document"     // Document file. 文档文件。
	CollabFileTypeSpreadsheet  CollabFileType = "spreadsheet"  // Spreadsheet file. 电子表格文件。
	CollabFileTypeDocumentPro  CollabFileType = "documentPro"  // Professional document file. 专业文档文件。
	CollabFileTypePresentation CollabFileType = "presentation" // Presentation file. 演示文件。
	CollabFileTypeTable        CollabFileType = "table"        // Table file. 表格文件。
)

// APIConf represents the configuration for an API call.
// APIConf 表示 API 调用的配置信息。
type APIConf struct {
	URL       string           // API endpoint URL. API 端点 URL。
	Method    string           // HTTP method. HTTP 方法。
	ReqParams APIRequestParams // Request parameters. 请求参数。
	ResParams APIResParams     // Response parameters. 响应参数。
	Client    *ehttp.Component
	ss        SignatureSigner
	req       *resty.Request
	auth      Metadata
	err       error
}

// APIRequestParams contains parameters for an API request.
// APIRequestParams 包含 API 请求的参数。
type APIRequestParams struct {
	Headers map[string][]string // HTTP request headers. HTTP 请求头。
	Query   map[string][]string // URL query parameters. URL 查询参数。
	Body    interface{}         // Request body. 请求体。
	Form    map[string][]string // Form data. 表单数据。
}

// APIResParams contains parameters for an API response.
// APIResParams 包含 API 响应的参数。
type APIResParams struct {
	SuccessHTTPCode int         // Expected HTTP status code. 期望的 HTTP 状态码。
	Body            interface{} // Response body. 响应体。
}

// Request creates and returns a configured HTTP request.
// Request 创建并返回配置好的 HTTP 请求。
func (ac *APIConf) Request(reqHandlers ...ReqHandler) *APIConf {
	req := ac.Client.R()
	req.URL = ac.URL
	req.Method = ac.Method
	req.Header = ac.ReqParams.Headers
	req.QueryParam = ac.ReqParams.Query
	req.Body = ac.ReqParams.Body
	req.FormData = ac.ReqParams.Form

	for _, h := range reqHandlers {
		h(req)
	}
	ac.req = req
	return ac
}

type apiRes interface {
	SetResponse(resp *resty.Response)
	Response() *resty.Response
}

type rawRes struct {
	res *resty.Response
}

func (r *rawRes) Response() *resty.Response {
	return r.res
}

func (r *rawRes) SetResponse(res *resty.Response) {
	r.res = res
}

type ReqHandler func(req *resty.Request)

func (ac *APIConf) Invoke(ctx context.Context, res apiRes) (err error) {
	var r *resty.Response
	defer func() {
		res.SetResponse(r)
	}()

	r, err = ac.req.SetContext(ctx).SetResult(res).Send()
	if err != nil {
		return fmt.Errorf("request execute failed: %w", err)
	}
	if r.StatusCode() != ac.ResParams.SuccessHTTPCode {
		return fmt.Errorf("request failed: status code %d, message: %s", r.StatusCode(), r.Body())
	}

	return err
}

func (ac *APIConf) SetResult(res interface{}) *APIConf {
	ac.req.SetResult(res)
	return ac
}

func (ac *APIConf) SetFileReader(param, fileName string, reader interface{}) *APIConf {
	if r, ok := reader.(interface{ Read([]byte) (int, error) }); ok {
		ac.req.SetFileReader(param, fileName, r)
	}
	return ac
}

func (ac *APIConf) Send(ctx context.Context) (*resty.Response, error) {
	rawResponse, err := ac.req.SetContext(ctx).Send()
	if err != nil {
		return rawResponse, fmt.Errorf("request execute failed: %w", err)
	}
	if rawResponse.StatusCode() != ac.ResParams.SuccessHTTPCode {
		return rawResponse, fmt.Errorf("request failed: status code %d, message: %s", rawResponse.StatusCode(), rawResponse.Body())
	}
	return rawResponse, nil
}
