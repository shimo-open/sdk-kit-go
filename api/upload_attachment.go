package api

import (
	"net/http"
	"os"

	"github.com/gotomicro/ego/client/ehttp"
)

// UploadAttachmentReq contains parameters for uploading a table attachment.
// UploadAttachmentReq 包含上传表格附件的参数。
// API Doc: https://open.shimo.im/docs/06API-document/interface-description/collaborative-editing#%E8%A1%A8%E6%A0%BC%E4%B8%8A%E4%BC%A0%E9%99%84%E4%BB%B6
type UploadAttachmentReq struct {
	Metadata
	// AccessToken is the upload access token.
	// AccessToken 是上传访问令牌。
	AccessToken string
	// Download is the list of download parameters.
	// Download 是下载参数列表。
	Download []string
	// File is the file to upload.
	// File 是要上传的文件。
	File *os.File
}

// UploadAttachmentRes contains the response for uploading an attachment.
// UploadAttachmentRes 包含上传附件的响应。
type UploadAttachmentRes struct {
	rawRes
	// Code is the response code.
	// Code 是响应代码。
	Code int `json:"code"`
	// Data is the response data.
	// Data 是响应数据。
	Data interface{} `json:"data"`
	// Body is the attachment result.
	// Body 是附件结果。
	Body AttachmentRes `json:"body"`
}

// AttachmentRes contains the attachment upload result.
// AttachmentRes 包含附件上传结果。
type AttachmentRes struct {
	// GUID is the attachment GUID.
	// GUID 是附件 GUID。
	GUID string `json:"GUID"`
	// Filename is the file name.
	// Filename 是文件名称。
	Filename string `json:"filename"`
	// Key is the attachment key.
	// Key 是附件 key。
	Key string `json:"key"`
	// MimeType is the MIME type (e.g., application/gzip).
	// MimeType 是 MIME 类型（例如：application/gzip）。
	MimeType string `json:"mimeType"`
	// Images is the images URL.
	// Images 是图片 URL。
	Images string `json:"images"`
	// Url is the attachment URL.
	// Url 是附件 URL。
	Url string `json:"url"`
	// Audio is the audio information.
	// Audio 是音频信息。
	Audio interface{} `json:"audio"`
	// Video is the video information.
	// Video 是视频信息。
	Video interface{} `json:"video"`
	// Width is the image width.
	// Width 是图片宽度。
	Width string `json:"width"`
	// Height is the image height.
	// Height 是图片高度。
	Height string `json:"height"`
	// Size is the file size.
	// Size 是文件大小。
	Size string `json:"size"`
}

// NewUploadAttachmentApi creates a new API config for uploading a table attachment.
// NewUploadAttachmentApi 创建用于上传表格附件的 API 配置。
func NewUploadAttachmentApi(cli *ehttp.Component, ss SignatureSigner, params UploadAttachmentReq) (apiConf *APIConf) {
	sign := ss.Sign(ExpireLong, ScopeDefault)
	return &APIConf{
		ss:     ss,
		Client: cli,
		URL:    "/uploader/upload",
		Method: http.MethodPost,
		ReqParams: APIRequestParams{
			Headers: addHeaders(params.Metadata, sign, nil),
			Form: map[string][]string{
				"accessToken": {params.AccessToken},
				"download":    params.Download,
			},
		},
		ResParams: APIResParams{
			SuccessHTTPCode: http.StatusOK,
			Body:            UploadAttachmentRes{},
		},
	}
}

// GetTokenReq contains parameters for getting upload token.
// GetTokenReq 包含获取上传令牌的参数。
type GetTokenReq struct {
	Metadata
	// GetUploadTokenReqBody is the list of upload token request bodies.
	// GetUploadTokenReqBody 是上传令牌请求体列表。
	GetUploadTokenReqBody []GetUploadTokenReqBody
}

// GetUploadTokenReqBody contains the request body for getting upload token.
// GetUploadTokenReqBody 包含获取上传令牌的请求体。
type GetUploadTokenReqBody struct {
	// Bucket is the storage bucket alias (attachments, assets, avatar, images).
	// Bucket 是存储桶别名（attachments, assets, avatar, images）。
	Bucket string `json:"bucket"`
	// Filename is the upload file name (default: untitled{ext}).
	// Filename 是上传文件名（默认：untitled{ext}）。
	Filename string `json:"filename"`
	// FileGuid is the file unique ID.
	// FileGuid 是文件唯一 ID。
	FileGuid string `json:"fileGuid"`
	// FileSize is the file size in bytes.
	// FileSize 是文件大小（字节）。
	FileSize int64 `json:"fileSize"`
	// Encrypt indicates whether to enable encryption (use "default" to enable).
	// Encrypt 表示是否开启加密（填 "default" 开启）。
	Encrypt string `json:"encrypt"`
}

// GetUploadTokenResToken contains the upload token response data.
// GetUploadTokenResToken 包含上传令牌响应数据。
type GetUploadTokenResToken struct {
	// AccessToken is the token for the upload step.
	// AccessToken 是上传步骤使用的令牌。
	AccessToken string `json:"accessToken"`
	// Directory is the upload directory.
	// Directory 是上传目录。
	Directory string `json:"directory"`
	// Download controls whether the URL includes download=1 parameter.
	// Download 控制 URL 是否携带 download=1 参数。
	Download string `json:"download"`
	// Guid is the attachment GUID.
	// Guid 是附件 GUID。
	Guid string `json:"guid"`
	// Images is the same as URL (for compatibility).
	// Images 与 URL 相同（兼容使用）。
	Images string `json:"images"`
	// FileFieldKey is the key name for file in the upload form (fixed: "file").
	// FileFieldKey 是上传表单中文件的 key 名称（固定值："file"）。
	FileFieldKey string `json:"fileFieldKey"`
	// Key is the attachment key.
	// Key 是附件 key。
	Key string `json:"key"`
	// ServerUrl is the server URL for the next upload step.
	// ServerUrl 是下一步上传的服务器地址。
	ServerUrl string `json:"serverUrl"`
	// Url is the attachment access URL.
	// Url 是附件访问 URL。
	Url string `json:"url"`
}

// GetUploadTokenRes contains the response for getting upload token.
// GetUploadTokenRes 包含获取上传令牌的响应。
type GetUploadTokenRes struct {
	rawRes
	// Tokens is the list of upload tokens.
	// Tokens 是上传令牌列表。
	Tokens []GetUploadTokenResToken `json:"tokens"`
}

// NewGetUploadTokenApi creates a new API config for getting upload token.
// NewGetUploadTokenApi 创建用于获取上传令牌的 API 配置。
func NewGetUploadTokenApi(cli *ehttp.Component, ss SignatureSigner, params GetTokenReq) *APIConf {
	sign := ss.Sign(ExpireLong, ScopeDefault)
	extra := map[string]string{
		"Content-Type": "application/json",
	}
	return &APIConf{
		ss:     ss,
		Client: cli,
		URL:    "/uploader/token",
		Method: http.MethodPost,
		ReqParams: APIRequestParams{
			Headers: addHeaders(params.Metadata, sign, extra),
			Body:    params.GetUploadTokenReqBody,
		},
		ResParams: APIResParams{
			SuccessHTTPCode: http.StatusOK,
			Body:            GetUploadTokenRes{},
		},
	}
}
