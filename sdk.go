package sdk

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/gotomicro/ego/client/ehttp"

	"github.com/shimo-open/sdk-kit-go/api"
)

// Manager manages the SDK core functionality.
// Manager 管理 SDK 的核心功能。
type Manager struct {
	appID      string
	appSecret  string
	httpClient *ehttp.Component
	ss         api.SignatureSigner
}

// Option defines a function type for configuring Manager.
// Option 定义用于配置 Manager 的函数类型。
type Option func(c *Manager)

// WithHTTPClient sets the http client.
// WithHTTPClient 设置 http 请求客户端。
func WithHTTPClient(httpClient *ehttp.Component) Option {
	return func(m *Manager) {
		m.httpClient = httpClient
	}
}

func WithAppID(appID string) Option {
	return func(m *Manager) {
		m.appID = appID
	}
}

func WithAppSecret(appSecret string) Option {
	return func(m *Manager) {
		m.appSecret = appSecret
	}
}

// NewManager creates a new SDK manager instance with the given configuration.
// NewManager 使用给定配置创建新的 SDK 管理器实例。
func NewManager(options ...Option) *Manager {
	m := &Manager{}
	for _, option := range options {
		option(m)
	}
	m.ss = api.NewSignatureSigner(m.appID, m.appSecret)
	return m
}

func (m *Manager) Sign(d time.Duration, scope api.Scope) string {
	return m.ss.Sign(d, scope)
}

// CreateFile creates a collaborative document.
// CreateFile 创建协同文档。
func (m *Manager) CreateFile(ctx context.Context, req api.CreateFileReq) (res api.CreateFileRes, err error) {
	err = api.NewCreateFileApi(m.httpClient, m.ss, req).Request().Invoke(ctx, &res)
	if err != nil {
		return res, fmt.Errorf("create file request failed: %w", err)
	}

	return res, nil
}

// CreateFileCopy creates a copy of a collaborative document.
// CreateFileCopy 创建协同文档副本。
func (m *Manager) CreateFileCopy(ctx context.Context, req api.CreateFileCopyReq) (res api.CreateFileCopyRes, err error) {
	err = api.NewCreateFileCopyApi(m.httpClient, m.ss, req).Request().Invoke(ctx, &res)
	if err != nil {
		return res, fmt.Errorf("create file copy request failed: %w", err)
	}

	return res, nil
}

// DeleteFile deletes a collaborative document.
// DeleteFile 删除协同文档。
func (m *Manager) DeleteFile(ctx context.Context, req api.DeleteFileReq) (res api.DeleteFileRes, err error) {
	err = api.NewDeleteFileApi(m.httpClient, m.ss, req).Request().Invoke(ctx, &res)
	if err != nil {
		return res, fmt.Errorf("delete file request failed: %w", err)
	}

	return res, nil
}

// GetHistoryList retrieves the history list of a document.
// GetHistoryList 获取历史列表。
func (m *Manager) GetHistoryList(ctx context.Context, req api.GetHistoryListReq) (res api.GetHistoryListRes, err error) {
	err = api.NewGetHistoryListApi(m.httpClient, m.ss, req).Request().Invoke(ctx, &res)
	if err != nil {
		return res, fmt.Errorf("get history list request failed: %w", err)
	}

	return res, nil
}

// GetRevisionList retrieves the revision list of a document.
// GetRevisionList 获取版本列表。
func (m *Manager) GetRevisionList(ctx context.Context, req api.GetRevisionListReq) (res api.GetRevisionListRes, err error) {
	var revisions []api.GetRevisionListRes
	rawRes, err := api.NewGetRevisionListApi(m.httpClient, m.ss, req).Request().SetResult(&revisions).Send(ctx)
	r := api.GetRevisionListRes{Revisions: revisions}
	r.SetResponse(rawRes)
	if err != nil {
		return r, fmt.Errorf("get revision list request failed: %w", err)
	}

	return r, nil
}

// GetPlainText retrieves the plain text content of a file.
// GetPlainText 获取文件纯文本内容。
func (m *Manager) GetPlainText(ctx context.Context, req api.GetPlainTextReq) (res api.GetPlainTextRes, err error) {
	err = api.NewGetPlainTextApi(m.httpClient, m.ss, req).Request().Invoke(ctx, &res)
	if err != nil {
		return res, fmt.Errorf("get plain text request failed: %w", err)
	}

	return res, nil
}

// GetPlainTextWC retrieves the word count of a file's plain text content.
// GetPlainTextWC 获取文件纯文本字数统计。
func (m *Manager) GetPlainTextWC(ctx context.Context, req api.GetPlainTextWCReq) (res api.GetPlainTextWCRes, err error) {
	err = api.NewGetPlainTextWCApi(m.httpClient, m.ss, req).Request().Invoke(ctx, &res)
	if err != nil {
		return res, fmt.Errorf("get plain text word count request failed: %w", err)
	}

	return res, nil
}

// GetMentionAt retrieves all @mention information in the file content.
// GetMentionAt 获取文件内容中所有的 @ 人信息列表。
func (m *Manager) GetMentionAt(ctx context.Context, req api.GetMentionAtReq) (res api.GetMentionAtRes, err error) {
	err = api.NewGetMentionAtApi(m.httpClient, m.ss, req).Request().Invoke(ctx, &res)
	if err != nil {
		return res, fmt.Errorf("get mention at request failed: %w", err)
	}

	return res, nil
}

// GetCommentCount retrieves the comment count of a file.
// GetCommentCount 获取文件中的评论数。
func (m *Manager) GetCommentCount(ctx context.Context, req api.GetCommentCountReq) (res api.GetCommentCountRes, err error) {
	err = api.NewGetCommentCountApi(m.httpClient, m.ss, req).Request().Invoke(ctx, &res)
	if err != nil {
		return res, fmt.Errorf("get comment count request failed: %w", err)
	}

	return res, nil
}

// GetTableContent retrieves the content of a table.
// GetTableContent 获取表格内容。
func (m *Manager) GetTableContent(ctx context.Context, req api.GetTableContentReq) (res api.GetTableContentRes, err error) {
	err = api.NewGetTableContentApi(m.httpClient, m.ss, req).Request().Invoke(ctx, &res)
	if err != nil {
		return res, fmt.Errorf("get table content request failed: %w", err)
	}

	return res, nil
}

// UpdateTableContent updates the content of a table.
// UpdateTableContent 更新表格内容。
func (m *Manager) UpdateTableContent(ctx context.Context, req api.UpdateTableContentReq) (res api.UpdateTableContentRes, err error) {
	err = api.NewUpdateTableContentApi(m.httpClient, m.ss, req).Request().Invoke(ctx, &res)
	if err != nil {
		return res, fmt.Errorf("update table content request failed: %w", err)
	}

	return res, nil
}

// AppendTableContent appends content to a table.
// AppendTableContent 追加表格内容。
func (m *Manager) AppendTableContent(ctx context.Context, req api.AppendTableContentReq) (res api.AppendTableContentRes, err error) {
	err = api.NewAppendTableContentApi(m.httpClient, m.ss, req).Request().Invoke(ctx, &res)
	if err != nil {
		return res, fmt.Errorf("append table content request failed: %w", err)
	}

	return res, nil
}

// DeleteTableRow deletes a row from a table.
// DeleteTableRow 删除表格行。
func (m *Manager) DeleteTableRow(ctx context.Context, req api.DeleteTableRowReq) (res api.DeleteTableRowRes, err error) {
	err = api.NewDeleteTableRowApi(m.httpClient, m.ss, req).Request().Invoke(ctx, &res)
	if err != nil {
		return res, fmt.Errorf("delete table row request failed: %w", err)
	}

	return res, nil
}

// AddTableSheet adds a new sheet to a table.
// AddTableSheet 新增表格工作表。
func (m *Manager) AddTableSheet(ctx context.Context, req api.AddTableSheetReq) (res api.AddTableSheetRes, err error) {
	err = api.NewAddTableSheetApi(m.httpClient, m.ss, req).Request().Invoke(ctx, &res)
	if err != nil {
		return res, fmt.Errorf("add table sheet request failed: %w", err)
	}

	return res, nil
}

// UploadAttachment uploads an attachment to a table.
// UploadAttachment 上传表格附件。
func (m *Manager) UploadAttachment(ctx context.Context, req api.GetTokenReq, file *os.File) (res api.UploadAttachmentRes, err error) {
	// Get upload token.
	// 获取上传 Token。
	var tokenRes []api.GetUploadTokenResToken
	if _, err = api.NewGetUploadTokenApi(m.httpClient, m.ss, req).Request().SetResult(&tokenRes).Send(ctx); err != nil {
		return api.UploadAttachmentRes{}, fmt.Errorf("get upload token request failed: %w", err)
	}

	downloads := make([]string, 0, len(tokenRes))
	for _, v := range tokenRes {
		downloads = append(downloads, v.Download)
	}

	// Build upload request.
	// 构造上传请求。
	uploadreq := api.UploadAttachmentReq{
		Metadata:    req.Metadata,
		AccessToken: tokenRes[0].AccessToken,
		Download:    downloads,
		File:        file,
	}
	err = api.NewUploadAttachmentApi(m.httpClient, m.ss, uploadreq).Request().SetFileReader("file", filepath.Base(file.Name()), file).Invoke(ctx, &res)
	if err != nil {
		return api.UploadAttachmentRes{}, fmt.Errorf("upload attachment request failed: %w", err)
	}

	return res, nil
}

// ReadBookmarkContent reads the content of a bookmark in a traditional document.
// ReadBookmarkContent 读取传统文档书签内容。
func (m *Manager) ReadBookmarkContent(ctx context.Context, req api.ReadBookmarkContentReq) (res api.ReadBookmarkContentRes, err error) {
	err = api.NewReadBookmarkContentApi(m.httpClient, m.ss, req).Request().Invoke(ctx, &res)
	if err != nil {
		return res, fmt.Errorf("read bookmark content request failed: %w", err)
	}

	return res, nil
}

// ReplaceBookmarkContent replaces the content of a bookmark in a traditional document.
// ReplaceBookmarkContent 替换传统文档书签内容。
func (m *Manager) ReplaceBookmarkContent(ctx context.Context, req api.RepBookmarkContentReq) (res api.RepBookmarkContentRes, err error) {
	err = api.NewReplaceBookmarkContentApi(m.httpClient, m.ss, req).Request().Invoke(ctx, &res)
	if err != nil {
		return res, fmt.Errorf("replace bookmark content request failed: %w", err)
	}

	return res, nil
}

// GetAppDetail retrieves the details of an app.
// GetAppDetail 获取 App 详情。
func (m *Manager) GetAppDetail(ctx context.Context, req api.GetAppDetailReq) (res api.GetAppDetailRes, err error) {
	err = api.NewGetAppDetailApi(m.httpClient, m.ss, req).Request().Invoke(ctx, &res)
	if err != nil {
		return res, fmt.Errorf("get app detail request failed: %w", err)
	}

	return res, nil
}

// UpdateCallbackURL updates the callback URL of an app.
// UpdateCallbackURL 更新 app 回调地址。
func (m *Manager) UpdateCallbackURL(ctx context.Context, req api.UpdateCallbackURLReq) (res api.UpdateCallbackURLRes, err error) {
	err = api.NewUpdateCallbackURLApi(m.httpClient, m.ss, req).Request().Invoke(ctx, &res)
	if err != nil {
		return res, fmt.Errorf("update callback url request failed: %w", err)
	}

	return res, nil
}

// GetUserAndStatus retrieves the user list and seat status.
// GetUserAndStatus 获取用户列表和席位状态。
func (m *Manager) GetUserAndStatus(ctx context.Context, req api.GetUserAndStatusReq) (res api.GetUserAndStatusRes, err error) {
	err = api.NewGetUserAndStatusApi(m.httpClient, m.ss, req).Request().Invoke(ctx, &res)
	if err != nil {
		return res, fmt.Errorf("get user and status request failed: %w", err)
	}

	return res, nil
}

// ActivateUserSeat activates a user seat.
// ActivateUserSeat 激活用户席位。
func (m *Manager) ActivateUserSeat(ctx context.Context, req api.ActivateUserSeatReq) (res api.ActivateUserSeatRes, err error) {
	err = api.NewActivateUserSeatApi(m.httpClient, m.ss, req).Request().Invoke(ctx, &res)
	if err != nil {
		return res, fmt.Errorf("activate user seat request failed: %w", err)
	}

	return res, nil
}

// CancelUserSeat cancels a user seat.
// CancelUserSeat 取消用户席位。
func (m *Manager) CancelUserSeat(ctx context.Context, req api.CancelUserSeatReq) (res api.CancelUserSeatRes, err error) {
	err = api.NewCancelUserSeatApi(m.httpClient, m.ss, req).Request().Invoke(ctx, &res)
	if err != nil {
		return res, fmt.Errorf("cancel user seat request failed: %w", err)
	}

	return res, nil
}

// BatchSetUserSeat sets user seats in batch.
// BatchSetUserSeat 批量设置用户席位。
func (m *Manager) BatchSetUserSeat(ctx context.Context, req api.BatchSetUserSeatReq) (res api.BatchSetUserSeatRes, err error) {
	err = api.NewBatchSetUserSeatApi(m.httpClient, m.ss, req).Request().Invoke(ctx, &res)
	if err != nil {
		return res, fmt.Errorf("batch set user seat request failed: %w", err)
	}

	return res, nil
}

// ImportFile imports a file.
// ImportFile 导入文件。
func (m *Manager) ImportFile(ctx context.Context, req api.ImportFileReq) (res api.ImportFileRes, err error) {
	rh := func(r *resty.Request) {
		if req.File != nil {
			r.SetFileReader("file", filepath.Base(req.File.Name()), req.File)
		}
	}

	err = api.NewImportFileApi(m.httpClient, m.ss, req).Request(rh).Invoke(ctx, &res)
	if err != nil {
		return res, fmt.Errorf("import file request failed: %w", err)
	}

	return res, nil
}

// GetImportProgress retrieves the import progress.
// GetImportProgress 获取导入进度。
func (m *Manager) GetImportProgress(ctx context.Context, req api.GetImportProgReq) (res api.GetImportProgRes, err error) {
	err = api.NewGetImportProgressApi(m.httpClient, m.ss, req).Request().Invoke(ctx, &res)
	if err != nil {
		return res, fmt.Errorf("get import progress request failed: %w", err)
	}

	return res, nil
}

// ImportV2File imports a file using v2 API.
// ImportV2File 使用 v2 接口导入文件。
func (m *Manager) ImportV2File(ctx context.Context, req api.ImportFileReq) (res api.ImportFileRes, err error) {
	rh := func(r *resty.Request) {
		if req.File != nil {
			r.SetFileReader("file", filepath.Base(req.File.Name()), req.File)
		}
	}
	err = api.NewImportV2FileApi(m.httpClient, m.ss, req).Request(rh).Invoke(ctx, &res)
	if err != nil {
		return res, fmt.Errorf("import v2 file request failed: %w", err)
	}

	return res, nil
}

// GetImportV2Progress retrieves the import progress using v2 API.
// GetImportV2Progress 使用 v2 接口获取导入进度。
func (m *Manager) GetImportV2Progress(ctx context.Context, req api.GetImportProgReq) (res api.GetImportProgRes, err error) {
	err = api.NewGetImportV2ProgressApi(m.httpClient, m.ss, req).Request().Invoke(ctx, &res)
	if err != nil {
		return res, fmt.Errorf("get import v2 progress request failed: %w", err)
	}

	return res, nil
}

// ExportFile exports a file.
// ExportFile 导出文件。
func (m *Manager) ExportFile(ctx context.Context, req api.ExportFileReq) (res api.ExportFileRes, err error) {
	err = api.NewExportFileApi(m.httpClient, m.ss, req).Request().Invoke(ctx, &res)
	if err != nil {
		return res, fmt.Errorf("export file request failed: %w", err)
	}

	return res, nil
}

// GetExportProgress retrieves the export progress.
// GetExportProgress 获取导出进度。
func (m *Manager) GetExportProgress(ctx context.Context, req api.GetExportProgReq) (res api.GetExportProgRes, err error) {
	err = api.NewGetExportProgressApi(m.httpClient, m.ss, req).Request().Invoke(ctx, &res)
	if err != nil {
		return res, fmt.Errorf("get export progress request failed: %w", err)
	}

	return res, nil
}

// ExportTableSheets exports table sheets to Excel.
// ExportTableSheets 导出应用表格为 Excel。
func (m *Manager) ExportTableSheets(ctx context.Context, req api.ExportTableSheetsReq) (res api.ExportTableSheetsRes, err error) {
	err = api.NewExportTableSheetsApi(m.httpClient, m.ss, req).Request().Invoke(ctx, &res)
	if err != nil {
		return res, fmt.Errorf("export table sheets request failed: %w", err)
	}

	return res, nil
}

// CreatePreview creates a preview.
// CreatePreview 创建预览。
func (m *Manager) CreatePreview(ctx context.Context, req api.CreatePreviewReq) (res api.CreatePreviewRes, err error) {
	err = api.NewCreatePreviewApi(m.httpClient, m.ss, req).Request().Invoke(ctx, &res)
	if err != nil {
		return res, fmt.Errorf("create preview request failed: %w", err)
	}

	return res, nil
}

// AccessPreview accesses a preview.
// AccessPreview 访问预览。
func (m *Manager) AccessPreview(ctx context.Context, req api.AccessPreviewReq) (res api.AccessPreviewRes, err error) {
	err = api.NewAccessPreviewApi(m.httpClient, m.ss, req).Request().Invoke(ctx, &res)
	if err != nil {
		return res, fmt.Errorf("access preview request failed: %w", err)
	}

	return res, nil
}

// ImportFileToAiKnowledgeBase imports a file to AI knowledge base.
// ImportFileToAiKnowledgeBase 导入文件到知识库。
func (m *Manager) ImportFileToAiKnowledgeBase(ctx context.Context, req api.ImportFileToAiKnowledgeBaseReq) (res api.ImportFileToAiKnowledgeBaseRes, err error) {
	err = api.NewImportFileToAiKnowledgeBaseApi(m.httpClient, m.ss, req).Request().Invoke(ctx, &res)
	if err != nil {
		return res, fmt.Errorf("import file to ai knowledge base request failed: %w", err)
	}

	return res, nil
}

// DeleteFileFromAiKnowledgeBase deletes a file from AI knowledge base.
// DeleteFileFromAiKnowledgeBase 从知识库中删除文档。
func (m *Manager) DeleteFileFromAiKnowledgeBase(ctx context.Context, req api.DeleteFileFromAiKnowledgeBaseReq) (res api.DeleteFileFromAiKnowledgeBaseRes, err error) {
	err = api.NewDeleteFileFromAiKnowledgeBaseApi(m.httpClient, m.ss, req).Request().Invoke(ctx, &res)
	if err != nil {
		return res, fmt.Errorf("delete file from ai knowledge base request failed: %w", err)
	}

	return res, nil
}

// ImportFileToAiKnowledgeBaseV2 imports a file to AI knowledge base using v2 API.
// ImportFileToAiKnowledgeBaseV2 使用 v2 接口导入文件到知识库。
func (m *Manager) ImportFileToAiKnowledgeBaseV2(ctx context.Context, req api.ImportFileToAiKnowledgeBaseV2Req) (res api.ImportFileToAiV2Res, err error) {
	err = api.NewImportFileToAiKnowledgeBaseV2Api(m.httpClient, m.ss, req).Request().Invoke(ctx, &res)
	if err != nil {
		return res, fmt.Errorf("import file to ai knowledge base v2 request failed: %w", err)
	}

	return res, nil
}

// GetImportFileToAiProgressV2 retrieves the import progress to AI knowledge base using v2 API.
// GetImportFileToAiProgressV2 使用 v2 接口获取导入到知识库的进度。
func (m *Manager) GetImportFileToAiProgressV2(ctx context.Context, req api.GetImportFileToAiProgressV2Req) (res api.GetImportFileToAiProgressV2Res, err error) {
	err = api.NewGetImportFileToAiProgressV2Api(m.httpClient, m.ss, req).Request().Invoke(ctx, &res)
	if err != nil {
		return res, fmt.Errorf("get import file to ai progress v2 request failed: %w", err)
	}

	return res, nil
}
