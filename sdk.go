package sdk_sdk

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/gotomicro/ego/client/ehttp"
	"github.com/pkg/errors"
	"github.com/shimo-open/sdk-kit-go/model/api"
)

var (
	SdkMgr *SdkManager
)

type SdkManager struct {
	httpClient *ehttp.Component
}

type SdkManagerConf struct {
	HttpOptions ClientOptions
}

func SetPrefix(prefix string) {
	api.Prefix = prefix
}

func NewSdkManager(conf SdkManagerConf) *SdkManager {
	return &SdkManager{
		httpClient: NewRestyClientWithOptions(conf.HttpOptions),
	}
}

func DefaultManagerConf() SdkManagerConf {
	return SdkManagerConf{
		HttpOptions: ClientOptions{},
	}
}

func DefaultSdkManager() *SdkManager {
	return NewSdkManager(DefaultManagerConf())
}

// CreateFile 创建协同文档
func (sm *SdkManager) CreateFile(params api.CreateFileParams) (res api.RawResponse, err error) {
	apiConf := api.NewCreateFileApi(params)
	req := apiConf.Request()
	resp, err := req.Send()

	// 数据归位
	res.Resp = resp

	if err != nil {
		err = errors.Wrap(err, "CreateFile request failed")
		return
	}

	if resp.StatusCode() != apiConf.RespParams.SuccessHttpCode {
		err = errors.Errorf("CreateFile request failed, status code: %d, message: %s", resp.StatusCode(), resp.Body())
		return
	}

	return
}

// CreateFileCopy 创建协同文档副本
func (sm *SdkManager) CreateFileCopy(params api.CreateFileCopyParams) (res api.RawResponse, err error) {
	apiConf := api.NewCreateFileCopyApi(params)
	req := apiConf.Request()
	resp, err := req.Send()

	// 数据归位
	res.Resp = resp

	if err != nil {
		err = errors.Wrap(err, "CreateFileCopy request failed")
		return
	}

	if resp.StatusCode() != apiConf.RespParams.SuccessHttpCode {
		err = errors.Errorf("CreateFileCopy request failed, status code: %d, message: %s", resp.StatusCode(), resp.Body())
		return
	}

	return
}

// DeleteFile 删除协同文档
func (sm *SdkManager) DeleteFile(params api.DeleteFileParams) (res api.RawResponse, err error) {
	apiConf := api.NewDeleteFileApi(params)
	req := apiConf.Request()
	resp, err := req.Send()

	// 数据归位
	res.Resp = resp

	if err != nil {
		err = errors.Wrap(err, "DeleteFile request failed")
		return
	}

	if resp.StatusCode() != apiConf.RespParams.SuccessHttpCode {
		err = errors.Errorf("DeleteFile request failed, status code: %d, message: %s", resp.StatusCode(), resp.Body())
		return
	}

	return
}

// GetHistoryList 获取历史列表 body
func (sm *SdkManager) GetHistoryList(params api.GetHistoryListParams) (res api.GetHistoryListRespBody, err error) {
	apiConf := api.NewGetHistoryListApi(params)
	req := apiConf.Request()
	resp, err := req.Send()

	// 数据归位
	res.Resp = resp

	if err != nil {
		err = errors.Wrap(err, "GetHistoryList request failed")
		return
	}

	if resp.StatusCode() != apiConf.RespParams.SuccessHttpCode {
		err = errors.Errorf("GetHistoryList request failed, status code: %d, message: %s", resp.StatusCode(), resp.Body())
		return
	}

	// 反序列化 JSON 数据到结构体
	_ = json.Unmarshal(resp.Body(), &res)

	return
}

// GetRevisionList 获取版本列表 body
func (sm *SdkManager) GetRevisionList(params api.GetRevisionListParams) (res api.RawResponse, resBody []api.GetRevisionListRespBody, err error) {
	apiConf := api.NewGetRevisionListApi(params)
	req := apiConf.Request()
	resp, err := req.Send()

	// 数据归位
	res.Resp = resp

	if err != nil {
		err = errors.Wrap(err, "GetRevisionList request failed")
		return
	}

	if resp.StatusCode() != apiConf.RespParams.SuccessHttpCode {
		err = errors.Errorf("GetRevisionList request failed, status code: %d, message: %s", resp.StatusCode(), resp.Body())
		return
	}

	err = json.Unmarshal(resp.Body(), &resBody)
	if err != nil {
		err = errors.Wrap(err, "GetRevisionList request failed")
		return
	}

	return
}

// GetPlainText 获取文件纯文本内容 body
func (sm *SdkManager) GetPlainText(params api.GetPlainTextParams) (res api.GetPlainTextRespBody, err error) {
	apiConf := api.NewGetPlainTextApi(params)
	req := apiConf.Request()
	resp, err := req.Send()

	// 数据归位
	res.Resp = resp

	if err != nil {
		err = errors.Wrap(err, "GetPlainText request failed")
		return
	}

	if resp.StatusCode() != apiConf.RespParams.SuccessHttpCode {
		err = errors.Errorf("GetPlainText request failed, status code: %d, message: %s", resp.StatusCode(), resp.Body())
		return
	}

	_ = json.Unmarshal(resp.Body(), &res)
	return
}

// GetPlainTextWC 文件纯文本字数统计 body
func (sm *SdkManager) GetPlainTextWC(params api.GetPlainTextWCParams) (res api.GetPlainTextWCRespBody, err error) {
	apiConf := api.NewGetPlainTextWCApi(params)
	req := apiConf.Request()
	resp, err := req.Send()

	// 数据归位
	res.Resp = resp

	if err != nil {
		err = errors.Wrap(err, "GetPlainTextWC request failed")
		return
	}

	if resp.StatusCode() != apiConf.RespParams.SuccessHttpCode {
		err = errors.Errorf("GetPlainTextWC request failed, status code: %d, message: %s", resp.StatusCode(), resp.Body())
		return
	}

	_ = json.Unmarshal(resp.Body(), &res)
	return
}

// GetMentionAt 获取文件内容中所有的 at 人信息列表 body
func (sm *SdkManager) GetMentionAt(params api.GetMentionAtParams) (res api.GetMentionAtRespBody, err error) {
	apiConf := api.NewGetMentionAtApi(params)
	req := apiConf.Request()
	resp, err := req.Send()

	// 数据归位
	res.Resp = resp

	if err != nil {
		err = errors.Wrap(err, "GetMentionAt request failed")
		return
	}

	if resp.StatusCode() != apiConf.RespParams.SuccessHttpCode {
		err = errors.Errorf("GetMentionAt request failed, status code: %d, message: %s", resp.StatusCode(), resp.Body())
		return
	}

	_ = json.Unmarshal(resp.Body(), &res)
	return
}

// GetCommentCount 获取文件中的评论数 body
func (sm *SdkManager) GetCommentCount(params api.GetCommentCountParams) (res api.GetCommentCountRespBody, err error) {
	apiConf := api.NewGetCommentCountApi(params)
	req := apiConf.Request()
	resp, err := req.Send()

	// 数据归位
	res.Resp = resp

	if err != nil {
		err = errors.Wrap(err, "GetCommentCount request failed")
		return
	}

	if resp.StatusCode() != apiConf.RespParams.SuccessHttpCode {
		err = errors.Errorf("GetCommentCount request failed, status code: %d, message: %s", resp.StatusCode(), resp.Body())
		return
	}

	_ = json.Unmarshal(resp.Body(), &res)
	return
}

// GetTableContent 获取表格内容
func (sm *SdkManager) GetTableContent(params api.GetTableContentParams) (res api.GetTableContentRespBody, err error) {
	apiConf := api.NewGetTableContentApi(params)
	req := apiConf.Request()
	resp, err := req.Send()

	// 数据归位
	res.Resp = resp

	if err != nil {
		err = errors.Wrap(err, "GetTableContent request failed")
		return
	}

	if resp.StatusCode() != apiConf.RespParams.SuccessHttpCode {
		err = errors.Errorf("GetTableContent request failed, status code: %d, message: %s", resp.StatusCode(), resp.Body())
		return
	}

	_ = json.Unmarshal(resp.Body(), &res)
	return
}

// UpdateTableContent 更新表格内容
func (sm *SdkManager) UpdateTableContent(params api.UpdateTableContentParams) (res api.RawResponse, err error) {
	apiConf := api.NewUpdateTableContentApi(params)
	req := apiConf.Request()
	resp, err := req.Send()

	// 数据归位
	res.Resp = resp

	if err != nil {
		err = errors.Wrap(err, "UpdateTableContent request failed")
		return
	}

	if resp.StatusCode() != apiConf.RespParams.SuccessHttpCode {
		err = errors.Errorf("UpdateTableContent request failed, status code: %d, message: %s", resp.StatusCode(), resp.Body())
		return
	}

	return
}

// AppendTableContent 追加表格内容
func (sm *SdkManager) AppendTableContent(params api.AppendTableContentParams) (res api.RawResponse, err error) {
	apiConf := api.NewAppendTableContentApi(params)
	req := apiConf.Request()
	resp, err := req.Send()

	// 数据归位
	res.Resp = resp

	if err != nil {
		err = errors.Wrap(err, "AppendTableContent request failed")
		return
	}

	if resp.StatusCode() != apiConf.RespParams.SuccessHttpCode {
		err = errors.Errorf("AppendTableContent request failed, status code: %d, message: %s", resp.StatusCode(), resp.Body())
		return
	}

	return
}

// DeleteTableRow 删除表格行
func (sm *SdkManager) DeleteTableRow(params api.DeleteTableRowParams) (res api.RawResponse, err error) {
	apiConf := api.NewDeleteTableRowApi(params)
	req := apiConf.Request()
	resp, err := req.Send()

	// 数据归位
	res.Resp = resp

	if err != nil {
		err = errors.Wrap(err, "DeleteTableRow request failed")
		return
	}

	if resp.StatusCode() != apiConf.RespParams.SuccessHttpCode {
		err = errors.Errorf("DeleteTableRow request failed, status code: %d, message: %s", resp.StatusCode(), resp.Body())
		return
	}

	return
}

// AddTableSheet 新增表格工作表
func (sm *SdkManager) AddTableSheet(params api.AddTableSheetParams) (res api.RawResponse, err error) {
	apiConf := api.NewAddTableSheetApi(params)
	req := apiConf.Request()
	resp, err := req.Send()

	// 数据归位
	res.Resp = resp

	if err != nil {
		err = errors.Wrap(err, "AddTableSheet request failed")
		return
	}

	if resp.StatusCode() != apiConf.RespParams.SuccessHttpCode {
		err = errors.Errorf("AddTableSheet request failed, status code: %d, message: %s", resp.StatusCode(), resp.Body())
		return
	}

	return
}

// UploadAttachment 表格上传附件
func (sm *SdkManager) UploadAttachment(params api.GetTokenParams, file *os.File) (res api.UploadAttachmentRespBody, err error) {
	tokenApiConf := api.NewGetUploadTokenApi(params)
	tokenReq := tokenApiConf.Request()
	tokenResp, err := tokenReq.Send()
	if err != nil {
		err = errors.Wrap(err, "UploadAttachment request failed")
		return
	}
	var TokenRespBody []api.GetUploadTokenRespBody
	if tokenResp.StatusCode() != tokenApiConf.RespParams.SuccessHttpCode {
		err = errors.Errorf("UploadAttachment request failed, status code: %d, message: %s", tokenResp.StatusCode(), tokenResp.Body())
		return
	}
	var downloads []string
	for _, v := range TokenRespBody {
		downloads = append(downloads, v.Download)
	}
	err = json.Unmarshal(tokenResp.Body(), &TokenRespBody)
	// 构造上传请求body
	uploadParams := api.UploadAttachmentParams{
		Auth:        params.Auth,
		AccessToken: TokenRespBody[0].AccessToken,
		Download:    downloads,
		File:        file,
	}
	uploadApiConf := api.NewUploadAttachmentApi(uploadParams)
	uploadReq := uploadApiConf.Request()
	uploadReq.SetFileReader("file", filepath.Base(file.Name()), file)

	uploadResp, err := uploadReq.Send()
	if err != nil {
		err = errors.Wrap(err, "UploadAttachment request failed")
		return
	}

	if uploadResp.StatusCode() != uploadApiConf.RespParams.SuccessHttpCode {
		err = errors.Errorf("UploadAttachment request failed, status code: %d", uploadResp.StatusCode())
		return
	}

	err = json.Unmarshal(uploadResp.Body(), &res)
	if err != nil {
		err = errors.Wrap(err, "UploadAttachment request failed")
		return
	}
	return
}

// ReadBookmarkContent 读取传统文档书签内容 body
func (sm *SdkManager) ReadBookmarkContent(params api.ReadBookmarkContentParams) (res api.ReadBookmarkContentRespBody, err error) {
	apiConf := api.NewReadBookmarkContentApi(params)
	req := apiConf.Request()
	resp, err := req.Send()

	// 数据归位
	res.Resp = resp

	if err != nil {
		err = errors.Wrap(err, "ReadBookmarkContent request failed")
		return
	}

	if resp.StatusCode() != apiConf.RespParams.SuccessHttpCode {
		err = errors.Errorf("ReadBookmarkContent request failed, status code: %d, message: %s", resp.StatusCode(), resp.Body())
		return
	}

	_ = json.Unmarshal(resp.Body(), &res)
	return
}

// ReplaceBookmarkContent 替换传统文档书签内容
func (sm *SdkManager) ReplaceBookmarkContent(params api.RepBookmarkContentParams) (res api.RawResponse, err error) {
	apiConf := api.NewReplaceBookmarkContentApi(params)
	req := apiConf.Request()
	resp, err := req.Send()

	// 数据归位
	res.Resp = resp

	if err != nil {
		err = errors.Wrap(err, "ReplaceBookmarkContent request failed")
		return
	}

	if resp.StatusCode() != apiConf.RespParams.SuccessHttpCode {
		err = errors.Errorf("ReplaceBookmarkContent request failed, status code: %d, message: %s", resp.StatusCode(), resp.Body())
		return
	}

	return
}

// 系统接口

// GetAppDetail 获取App详情 body
func (sm *SdkManager) GetAppDetail(params api.GetAppDetailParams) (res api.GetAppDetailRespBody, err error) {
	apiConf := api.NewGetAppDetailApi(params)
	req := apiConf.Request()
	resp, err := req.Send()
	err = json.Unmarshal(resp.Body(), &res)

	// 数据归位
	res.Resp = resp

	if err != nil {
		err = errors.Wrap(err, "GetAppDetail request failed")
		return
	}

	if resp.StatusCode() != apiConf.RespParams.SuccessHttpCode {
		err = errors.Errorf("GetAppDetail request failed, status code: %d, message: %s", resp.StatusCode(), resp.Body())
		return
	}

	return
}

// UpdateCallbackUrl 更新app回调地址
func (sm *SdkManager) UpdateCallbackUrl(params api.UpdateCallbackUrlParams) (res api.RawResponse, err error) {
	apiConf := api.NewUpdateCallbackUrlApi(params)
	req := apiConf.Request()
	resp, err := req.Send()

	// 数据归位
	res.Resp = resp

	if err != nil {
		err = errors.Wrap(err, "UpdateCallbackUrl request failed")
		return
	}

	if resp.StatusCode() != apiConf.RespParams.SuccessHttpCode {
		err = errors.Errorf("UpdateCallbackUrl request failed, status code: %d, message: %s", resp.StatusCode(), resp.Body())
		return
	}

	return
}

// GetUserAndStatus 获取用户列表和席位状态 body
func (sm *SdkManager) GetUserAndStatus(params api.GetUserAndStatusParams) (res api.RawResponse, resBody []api.GetUserAndStatusRespBody, err error) {
	apiConf := api.NewGetUserAndStatusApi(params)
	req := apiConf.Request()
	resp, err := req.Send()

	// 数据归位
	res.Resp = resp

	if err != nil {
		err = errors.Wrap(err, "GetUserAndStatus request failed")
		return
	}

	if resp.StatusCode() != apiConf.RespParams.SuccessHttpCode {
		err = errors.Errorf("GetUserAndStatus request failed, status code: %d, message: %s", resp.StatusCode(), resp.Body())
		return
	}

	err = json.Unmarshal(resp.Body(), &resBody)
	if err != nil {
		err = errors.Wrap(err, "GetUserAndStatus request failed")
		return
	}
	return
}

// ActivateUserSeat 激活用户席位
func (sm *SdkManager) ActivateUserSeat(params api.ActivateUserSeatParams) (res api.RawResponse, err error) {
	apiConf := api.NewActivateUserSeatApi(params)
	req := apiConf.Request()
	resp, err := req.Send()

	// 数据归位
	res.Resp = resp

	if err != nil {
		err = errors.Wrap(err, "ActivateUserSeat request failed")
		return
	}

	if resp.StatusCode() != apiConf.RespParams.SuccessHttpCode {
		err = errors.Errorf("ActivateUserSeat request failed, status code: %d, message: %s", resp.StatusCode(), resp.Body())
		return
	}

	return
}

// CancelUserSeat 取消用户席位
func (sm *SdkManager) CancelUserSeat(params api.CancelUserSeatParams) (res api.RawResponse, err error) {
	apiConf := api.NewCancelUserSeatApi(params)
	req := apiConf.Request()
	resp, err := req.Send()

	// 数据归位
	res.Resp = resp

	if err != nil {
		err = errors.Wrap(err, "CancelUserSeat request failed")
		return
	}

	if resp.StatusCode() != apiConf.RespParams.SuccessHttpCode {
		err = errors.Errorf("CancelUserSeat request failed, status code: %d, message: %s", resp.StatusCode(), resp.Body())
		return
	}

	return
}

// BatchSetUserSeat 批量设置用户席位
func (sm *SdkManager) BatchSetUserSeat(params api.BatchSetUserSeatParams) (res api.RawResponse, err error) {
	apiConf := api.NewBatchSetUserSeatApi(params)
	req := apiConf.Request()
	resp, err := req.Send()

	// 数据归位
	res.Resp = resp

	if err != nil {
		err = errors.Wrap(err, "BatchSetUserSeat request failed")
		return
	}

	if resp.StatusCode() != apiConf.RespParams.SuccessHttpCode {
		err = errors.Errorf("BatchSetUserSeat request failed, status code: %d, message: %s", resp.StatusCode(), resp.Body())
		return
	}

	return
}

// ImportFile 文件导入
func (sm *SdkManager) ImportFile(params api.ImportFileParams) (res api.ImportFileRespBody, err error) {
	apiConf, err := api.NewImportFileApi(params)
	req := apiConf.Request()

	if params.File != nil {
		req.SetFileReader("file", filepath.Base(params.File.Name()), params.File)
	}

	resp, err := req.Send()

	// 数据归位
	res.Resp = resp

	if err != nil {
		err = errors.Wrap(err, "ImportFile request failed")
		return
	}

	if resp.StatusCode() != apiConf.RespParams.SuccessHttpCode {
		err = errors.Errorf("ImportFile request failed, status code: %d, message: %s", resp.StatusCode(), resp.Body())
		return
	}

	err = json.Unmarshal(resp.Body(), &res)
	fmt.Println(res.Data.TaskId)

	if err != nil {
		err = errors.Wrap(err, "ImportFile request failed")
		return
	}

	return
}

// GetImportProgress 获取导入进度
func (sm *SdkManager) GetImportProgress(params api.GetImportProgParams) (res api.GetImportProgRespBody, err error) {
	apiConf := api.NewGetImportProgressApi(params)
	req := apiConf.Request()
	resp, err := req.Send()

	// 数据归位
	res.Resp = resp

	if err != nil {
		err = errors.Wrap(err, "GetImportProgress request failed")
		return
	}

	if resp.StatusCode() != apiConf.RespParams.SuccessHttpCode {
		err = errors.Errorf("GetImportProgress request failed, status code: %d, message: %s", resp.StatusCode(), resp.Body())
		return
	}

	_ = json.Unmarshal(resp.Body(), &res)
	return
}

// ImportV2File 文件导入v2
func (sm *SdkManager) ImportV2File(params api.ImportFileParams) (res api.ImportFileRespBody, err error) {
	apiConf, err := api.NewImportV2FileApi(params)
	req := apiConf.Request()

	if params.File != nil {
		req.SetFileReader("ImportV2File file", filepath.Base(params.File.Name()), params.File)
	}

	resp, err := req.Send()

	// 数据归位
	res.Resp = resp

	if err != nil {
		err = errors.Wrap(err, "ImportV2File request failed")
		return
	}

	if resp.StatusCode() != apiConf.RespParams.SuccessHttpCode {
		err = errors.Errorf("ImportV2File request failed, status code: %d, message: %s", resp.StatusCode(), resp.Body())
		return
	}

	err = json.Unmarshal(resp.Body(), &res)
	fmt.Println(res.Data.TaskId)

	if err != nil {
		err = errors.Wrap(err, "ImportV2File request failed")
		return
	}

	return
}

// GetImportV2Progress 获取导入进度v2
func (sm *SdkManager) GetImportV2Progress(params api.GetImportProgParams) (res api.GetImportProgRespBody, err error) {
	apiConf := api.NewGetImportV2ProgressApi(params)
	req := apiConf.Request()
	resp, err := req.Send()

	// 数据归位
	res.Resp = resp

	if err != nil {
		err = errors.Wrap(err, "GetImportV2Progress request failed")
		return
	}

	if resp.StatusCode() != apiConf.RespParams.SuccessHttpCode {
		err = errors.Errorf("GetImportV2Progress request failed, status code: %d, message: %s", resp.StatusCode(), resp.Body())
		return
	}

	_ = json.Unmarshal(resp.Body(), &res)
	return
}

// ExportFile 导出文件 body
func (sm *SdkManager) ExportFile(params api.ExportFileParams) (res api.ExportFileRespBody, err error) {
	apiConf := api.NewExportFileApi(params)
	req := apiConf.Request()
	resp, err := req.Send()

	// 数据归位
	res.Resp = resp

	if err != nil {
		err = errors.Wrap(err, "ExportFile request failed")
		return
	}

	if resp.StatusCode() != apiConf.RespParams.SuccessHttpCode {
		err = errors.Errorf("ExportFile request failed, status code: %d, message: %s", resp.StatusCode(), resp.Body())
		return
	}

	err = json.Unmarshal(resp.Body(), &res)
	fmt.Println(res.Data.TaskId)
	if err != nil {
		err = errors.Wrap(err, "ExportFile request failed")
		return
	}

	return
}

// GetExportProgress 获取导出进度
func (sm *SdkManager) GetExportProgress(params api.GetExportProgParams) (res api.GetExportProgRespBody, err error) {
	apiConf := api.NewGetExportProgressApi(params)
	req := apiConf.Request()
	resp, err := req.Send()

	// 数据归位
	res.Resp = resp

	if err != nil {
		err = errors.Wrap(err, "GetExportProgress request failed")
		return
	}

	if resp.StatusCode() != apiConf.RespParams.SuccessHttpCode {
		err = errors.Errorf("GetExportProgress request failed, status code: %d, message: %s", resp.StatusCode(), resp.Body())
		return
	}

	_ = json.Unmarshal(resp.Body(), &res)
	return
}

// ExportTableSheets 导出应用表格为 Excel
func (sm *SdkManager) ExportTableSheets(params api.ExportTableSheetsParams) (res api.ExportTableSheetsRespBody, err error) {
	apiConf := api.NewExportTableSheetsApi(params)
	req := apiConf.Request()
	resp, err := req.Send()

	// 数据归位
	res.Resp = resp

	if err != nil {
		err = errors.Wrap(err, "ExportTableSheets request failed")
		return
	}

	if resp.StatusCode() != apiConf.RespParams.SuccessHttpCode {
		err = errors.Errorf("ExportTableSheets request failed, status code: %d, message: %s", resp.StatusCode(), resp.Body())
		return
	}

	_ = json.Unmarshal(resp.Body(), &res)
	return
}

// CreatePreview 创建预览
func (sm *SdkManager) CreatePreview(params api.CreatePreviewParams) (res api.CreatePreviewRespBody, err error) {
	apiConf := api.NewCreatePreviewApi(params)
	req := apiConf.Request()
	resp, err := req.Send()

	// 数据归位
	res.Resp = resp

	if err != nil {
		err = errors.Wrap(err, "CreatePreview request failed")
		return
	}

	if resp.StatusCode() != apiConf.RespParams.SuccessHttpCode {
		err = errors.Errorf("CreatePreview request failed, status code: %d, message: %s", resp.StatusCode(), resp.Body())
		return
	}

	return
}

// AccessPreview 访问预览
func (sm *SdkManager) AccessPreview(params api.AccessPreviewParams) (res api.RawResponse, err error) {
	apiConf := api.NewAccessPreviewApi(params)
	req := apiConf.Request()
	resp, err := req.Send()

	// 数据归位
	res.Resp = resp

	if err != nil {
		err = errors.Wrap(err, "AccessPreview request failed")
		return
	}

	if resp.StatusCode() != apiConf.RespParams.SuccessHttpCode {
		err = errors.Errorf("AccessPreview request failed, status code: %d, message: %s", resp.StatusCode(), resp.Body())
		return
	}

	return
}

// ImportFileToAiKnowledgeBase 导入文件到知识库
func (sm *SdkManager) ImportFileToAiKnowledgeBase(params api.ImportFileToAiKnowledgeBaseParams) (res api.RawResponse, err error) {
	apiConf := api.NewImportFileToAiKnowledgeBaseApi(params)
	req := apiConf.Request()
	resp, err := req.Send()

	// 数据归位
	res.Resp = resp

	if err != nil {
		err = errors.Wrap(err, "ImportFileToAiKnowledgeBase request failed")
		return
	}

	if resp.StatusCode() != apiConf.RespParams.SuccessHttpCode {
		err = errors.Errorf("ImportFileToAiKnowledgeBase request failed, status code: %d, message: %s", resp.StatusCode(), resp.Body())
		return
	}

	return
}

// DeleteFileFromAiKnowledgeBase 从知识库中删除文档
func (sm *SdkManager) DeleteFileFromAiKnowledgeBase(params api.DeleteFileFromAiKnowledgeBaseParams) (res api.RawResponse, err error) {
	apiConf := api.NewDeleteFileFromAiKnowledgeBaseApi(params)
	req := apiConf.Request()
	resp, err := req.Send()

	// 数据归位
	res.Resp = resp

	if err != nil {
		err = errors.Wrap(err, "DeleteFileFromAiKnowledgeBase request failed")
		return
	}

	if resp.StatusCode() != apiConf.RespParams.SuccessHttpCode {
		err = errors.Errorf("DeleteFileFromAiKnowledgeBase request failed, status code: %d, message: %s", resp.StatusCode(), resp.Body())
		return
	}

	return
}

// ImportFileToAiKnowledgeBaseV2 导入文件到知识库v2
func (sm *SdkManager) ImportFileToAiKnowledgeBaseV2(params api.ImportFileToAiKnowledgeBaseV2Params) (res api.ImportFileToAiV2RespBody, err error) {
	apiConf := api.NewImportFileToAiKnowledgeBaseV2Api(params)
	req := apiConf.Request()
	resp, err := req.Send()

	// 数据归位
	res.Resp = resp

	if err != nil {
		err = errors.Wrap(err, "ImportFileToAiKnowledgeBaseV2 request failed")
		return
	}

	if resp.StatusCode() != apiConf.RespParams.SuccessHttpCode {
		err = errors.Errorf("ImportFileToAiKnowledgeBaseV2 request failed, status code: %d, message: %s", resp.StatusCode(), resp.Body())
		return
	}

	err = json.Unmarshal(resp.Body(), &res)
	if err != nil {
		err = errors.Wrap(err, "ImportFileToAiKnowledgeBaseV2 request failed")
		return
	}

	return
}

// GetImportFileToAiProgressV2 获取导入进度v2
func (sm *SdkManager) GetImportFileToAiProgressV2(params api.GetImportFileToAiProgressV2Params) (res api.GetImportFileToAiProgressV2RespBody, err error) {
	apiConf := api.NewGetImportFileToAiProgressV2Api(params)
	req := apiConf.Request()
	resp, err := req.Send()

	// 数据归位
	res.Resp = resp

	if err != nil {
		err = errors.Wrap(err, "GetImportFileToAiProgressV2 request failed")
		return
	}

	if resp.StatusCode() != apiConf.RespParams.SuccessHttpCode {
		err = errors.Errorf("GetImportFileToAiProgressV2 request failed, status code: %d, message: %s", resp.StatusCode(), resp.Body())
		return
	}

	err = json.Unmarshal(resp.Body(), &res)
	if err != nil {
		err = errors.Wrap(err, "GetImportFileToAiProgressV2 request failed")
		return
	}

	return
}
