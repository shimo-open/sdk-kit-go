package api

import (
	"net/http"
	"os"

	"github.com/gotomicro/ego/client/ehttp"
)

// 表格上传附件
// https://open.shimo.im/docs/06API-document/interface-description/collaborative-editing#%E8%A1%A8%E6%A0%BC%E4%B8%8A%E4%BC%A0%E9%99%84%E4%BB%B6

type UploadAttachmentReq struct {
	Metadata
	AccessToken string
	Download    []string
	File        *os.File
}

type UploadAttachmentRes struct {
	rawRes
	Code int           `json:"code"`
	Data interface{}   `json:"data"`
	Body AttachmentRes `json:"body"`
}

type AttachmentRes struct {
	GUID     string      `json:"GUID"`     // 附件 GUID
	Filename string      `json:"filename"` // 文件名称
	Key      string      `json:"key"`      // 附件 key
	MimeType string      `json:"mimeType"` // 例如 application/gzip
	Images   string      `json:"images"`
	Url      string      `json:"url"`
	Audio    interface{} `json:"audio"`  // 音频信息 例如
	Video    interface{} `json:"video"`  // 视频信息 {"width":"640","height":"352","duration":"17.333333","fps":"","fileFormat":"","bitrate":"49710"}
	Width    string      `json:"width"`  // 图片宽度
	Height   string      `json:"height"` // 图片高度
	Size     string      `json:"size"`   // 大小
}

func NewUploadAttachmentApi(cli *ehttp.Component, ss SignatureSigner, params UploadAttachmentReq) (apiConf *APIConf) {
	sign := ss.Sign(expire4m, ScopeDefault)
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

type GetTokenReq struct {
	Metadata
	GetUploadTokenReqBody []GetUploadTokenReqBody
}

type GetUploadTokenReqBody struct {
	Bucket   string `json:"bucket"`   // 别名，必须为以下几种之一：attachments,assets,avatar,images
	Filename string `json:"filename"` // 指定上传文件的文件名, 为空时使用 untitled{ext} 作为文件名
	FileGuid string `json:"fileGuid"` // 文件唯一ID
	FileSize int64  `json:"fileSize"` // 文件大小，单位字节
	Encrypt  string `json:"encrypt"`  // 是否开启加密，若开启填 default
}

type GetUploadTokenResToken struct {
	AccessToken  string `json:"accessToken"` // 第二步上传时候使用的 accessToken
	Directory    string `json:"directory"`
	Download     string `json:"download"`     // 控制生成的 url 是否携带 download=1 参数
	Guid         string `json:"guid"`         // 附件 guid
	Images       string `json:"images"`       // 同 url，兼容使用
	FileFieldKey string `json:"fileFieldKey"` // 第二步上传时 file 应该使用的 key name。固定值 file
	Key          string `json:"key"`          // 附件 key
	ServerUrl    string `json:"serverUrl"`    // 下一步上传需要提交 post 表单到的服务器地址
	Url          string `json:"url"`          // 附件访问 url
}

type GetUploadTokenRes struct {
	rawRes
	Tokens []GetUploadTokenResToken `json:"tokens"`
}

func NewGetUploadTokenApi(cli *ehttp.Component, ss SignatureSigner, params GetTokenReq) *APIConf {
	sign := ss.Sign(expire4m, ScopeDefault)
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
