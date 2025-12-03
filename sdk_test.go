package sdk

import (
	"context"
	"os"
	"path/filepath"
	"testing"

	"github.com/gotomicro/ego/client/ehttp"
	"github.com/stretchr/testify/assert"

	"github.com/shimo-open/sdk-kit-go/api"
)

var (
	sdkMgr  *Manager
	auth    api.Metadata
	ctxTest = context.TODO()
)

func TestMain(m *testing.M) {
	// 初始化 sdkMgr 和 auth
	httpClient := ehttp.Load("").Build(ehttp.WithRawDebug(true), ehttp.WithAddr("http://co-dev-16.shimorelease.com"))
	sdkMgr = NewManager(WithHTTPClient(httpClient))
	auth = api.Metadata{
		ShimoToken: "eyJhbGciOiJIUzI1NiIsImtpZCI6InNoaW1vZGV2IiwidHlwIjoiSldUIn0.eyJleHAiOjE3MzU5ODIwMTQsInVzZXJJZCI6OTMsIm1vZGUiOiJzaGltbyJ9.yZmaTtbKWb9cw9yXjbzqWEKgP9RJA1eO-tLlpMdrRCA",
	}

	// 运行测试
	exitCode := m.Run()

	// 退出程序
	os.Exit(exitCode)
}

func TestCreateFile(t *testing.T) {
	createFileParams := api.CreateFileReq{
		Metadata: auth,
		FileType: api.CollabFileTypeDocument,
		FileID:   "test_file_99",
		Lang:     api.LangZhCN,
	}
	_, err := sdkMgr.CreateFile(ctxTest, createFileParams)
	assert.NoError(t, err)
}

func TestCreateFileCopy(t *testing.T) {
	createFileCopyParams := api.CreateFileCopyReq{
		Metadata:     auth,
		OriginFileID: "1adb10eb32fe4d6c",
		TargetFileID: "test_file_copy_5",
	}
	_, err := sdkMgr.CreateFileCopy(ctxTest, createFileCopyParams)
	assert.Equal(t, nil, err)
}

func TestDeleteFile(t *testing.T) {
	deleteFileParams := api.DeleteFileReq{
		Metadata: auth,
		FileID:   "01bf99751c2d472f",
	}
	_, err := sdkMgr.DeleteFile(ctxTest, deleteFileParams)
	assert.NoError(t, err)
}

func TestGetHistoryList(t *testing.T) {
	getHistoryListParams := api.GetHistoryListReq{
		Metadata:    auth,
		FileID:      "7b8558bd4df84dd2",
		PageSize:    10,
		Count:       0,
		HistoryType: 1,
	}
	_, err := sdkMgr.GetHistoryList(ctxTest, getHistoryListParams)
	assert.Equal(t, nil, err)
}

func TestGetRevisionList(t *testing.T) {
	_, err := sdkMgr.GetRevisionList(ctxTest, api.GetRevisionListReq{
		Metadata: auth,
		FileID:   "1adb10eb32fe4d6c",
	})
	assert.Equal(t, nil, err)
}

func TestGetPlainText(t *testing.T) {
	_, err := sdkMgr.GetPlainText(ctxTest, api.GetPlainTextReq{
		Metadata: auth,
		FileID:   "1adb10eb32fe4d6c",
	})
	assert.Equal(t, nil, err)
}

func TestGetPlainTextWC(t *testing.T) {
	_, err := sdkMgr.GetPlainTextWC(ctxTest, api.GetPlainTextWCReq{
		Metadata: auth,
		FileID:   "910759b9232a4c73",
	})
	assert.Equal(t, nil, err)
}

func TestGetMentionAt(t *testing.T) {
	_, err := sdkMgr.GetMentionAt(ctxTest, api.GetMentionAtReq{
		Metadata: auth,
		FileID:   "910759b9232a4c73",
	})
	assert.Equal(t, nil, err)
}

func TestGetCommentCount(t *testing.T) {
	_, err := sdkMgr.GetCommentCount(ctxTest, api.GetCommentCountReq{
		Metadata: auth,
		FileID:   "81e7a08e651f4052",
	})
	assert.Equal(t, nil, err)
}

func TestGetTableContent(t *testing.T) {
	_, err := sdkMgr.GetTableContent(ctxTest, api.GetTableContentReq{
		Metadata: auth,
		FileID:   "81e7a08e651f4052",
		Rg:       "工作表1!A1:C3",
	})
	assert.Equal(t, nil, err)
}

func TestUpdateTableContent(t *testing.T) {
	reqBody := api.UpdateTableContentRequestBody{
		Rg: "工作表1!A1:C3",
		Resource: struct {
			Values [][]interface{} `json:"values"`
		}{
			Values: [][]interface{}{
				{"第一行第一列的值", "第一行第二列的值"},
				{"第二行第一列的值", "第二行第二列的值"},
			},
		},
	}
	_, err := sdkMgr.UpdateTableContent(ctxTest, api.UpdateTableContentReq{
		Metadata:                      auth,
		FileID:                        "81e7a08e651f4052",
		UpdateTableContentRequestBody: reqBody,
	})
	assert.Equal(t, nil, err)
}

func TestAppendTableContent(t *testing.T) {
	resource := api.Resource{
		Values: [][]interface{}{
			{"第一行第一列追加文本", "第一行第二列追加文本"},
			{"第二行第一列追加文本", "第二行第二列追加文本"},
		},
	}
	_, err := sdkMgr.AppendTableContent(ctxTest, api.AppendTableContentReq{
		Metadata: auth,
		FileID:   "81e7a08e651f4052",
		AppendTableContentReqBody: api.AppendTableContentReqBody{
			Rg:       "工作表1!A1:C3",
			Resource: resource,
		},
	})
	assert.Equal(t, nil, err)
}

func TestDeleteTableRow(t *testing.T) {
	_, err := sdkMgr.DeleteTableRow(ctxTest, api.DeleteTableRowReq{
		Metadata:  auth,
		FileID:    "81e7a08e651f4052",
		SheetName: "工作表1",
		Index:     0,
		Count:     1,
	})
	assert.Equal(t, nil, err)
}

func TestAddTableSheet(t *testing.T) {
	_, err := sdkMgr.AddTableSheet(ctxTest, api.AddTableSheetReq{
		Metadata: auth,
		FileID:   "81e7a08e651f4052",
		AddTableSheetReqBody: api.AddTableSheetReqBody{
			Name: "工作表 2",
		},
	})
	assert.Equal(t, nil, err)
}

func TestUploadAttachment(t *testing.T) {

	relPath := "./resource/import/test.doc"
	rootPath, err := filepath.Abs(relPath)
	file, err := os.Open(rootPath)
	defer file.Close()
	fileInfo, err := file.Stat()
	fileSize := fileInfo.Size()

	reqBody := []api.GetUploadTokenReqBody{
		{
			Bucket:   "attachments", // 别名，必须为以下几种之一：attachments,assets,avatar,images
			Filename: "text.doc",    // 指定上传文件的文件名, 为空时使用 untitled{ext} 作为文件名
			FileGuid: "haha",
			FileSize: fileSize, // 文件大小，单位字节
			Encrypt:  "",       // 是否开启加密，若开启填 default
		},
	}

	_, err = sdkMgr.UploadAttachment(ctxTest, api.GetTokenReq{
		Metadata:              auth,
		GetUploadTokenReqBody: reqBody,
	},
		file,
	)
	assert.Equal(t, nil, err)
}

func TestReadBookmarkContent(t *testing.T) {
	_, err := sdkMgr.ReadBookmarkContent(ctxTest, api.ReadBookmarkContentReq{
		Metadata:  auth,
		FileID:    "2ef5679a064a4769",
		Bookmarks: []string{"guid"},
	})
	assert.Equal(t, nil, err)
}

func TestReplaceBookmarkContent(t *testing.T) {
	replacement := api.Replacement{
		Bookmark: "bookmark",
		Type:     "text",
		Value:    "bookmark value",
	}
	reqBody := []api.Replacement{replacement}
	_, err := sdkMgr.ReplaceBookmarkContent(ctxTest, api.RepBookmarkContentReq{
		Metadata:                  auth,
		FileID:                    "2ef5679a064a4769",
		RepBookmarkContentReqBody: api.RepBookmarkContentReqBody{Replacements: reqBody},
	})
	assert.Equal(t, nil, err)
}

// 系统接口

func TestGetAppDetail(t *testing.T) {
	authLicense := api.Metadata{
		ShimoToken: "eyJhbGciOiJIUzI1NiIsImtpZCI6InNoaW1vZGV2IiwidHlwIjoiSldUIn0.eyJleHAiOjE3MzEwNTM3NDAsInVzZXJJZCI6ODYsIm1vZGUiOiJzaGltbyJ9.Z3P240q3Chd6r2PrsWtysBN50ixeGNrxZ-5PpAiBRUg",
	}
	_, err := sdkMgr.GetAppDetail(ctxTest, api.GetAppDetailReq{
		AppID:    "shimodev",
		Metadata: authLicense,
	})
	assert.Equal(t, nil, err)
}

func TestUpdateCallbackURL(t *testing.T) {
	authLicense := api.Metadata{
		ShimoToken: "eyJhbGciOiJIUzI1NiIsImtpZCI6InNoaW1vZGV2IiwidHlwIjoiSldUIn0.eyJleHAiOjE3MzEwNTM3NDAsInVzZXJJZCI6ODYsIm1vZGUiOiJzaGltbyJ9.Z3P240q3Chd6r2PrsWtysBN50ixeGNrxZ-5PpAiBRUg",
	}
	_, err := sdkMgr.UpdateCallbackURL(ctxTest, api.UpdateCallbackURLReq{
		AppID:                    "f83b43b4ba644f26837238e983f86f86",
		Metadata:                 authLicense,
		UpdateCallbackURLReqBody: api.UpdateCallbackURLReqBody{URL: "https://co-sdk-dev.shimorelease.com"},
	})
	assert.Equal(t, nil, err)
}

func TestGetUserAndStatus(t *testing.T) {
	authLicense := api.Metadata{
		ShimoToken: "eyJhbGciOiJIUzI1NiIsImtpZCI6InNoaW1vZGV2IiwidHlwIjoiSldUIn0.eyJleHAiOjE3MzEwNTM3NDAsInVzZXJJZCI6ODYsIm1vZGUiOiJzaGltbyJ9.Z3P240q3Chd6r2PrsWtysBN50ixeGNrxZ-5PpAiBRUg",
	}
	_, err := sdkMgr.GetUserAndStatus(ctxTest, api.GetUserAndStatusReq{
		Metadata: authLicense,
		Page:     1,
		Size:     30,
	})
	assert.Equal(t, nil, err)
}

func TestActivateUserSeat(t *testing.T) {
	authLicense := api.Metadata{
		ShimoToken: "eyJhbGciOiJIUzI1NiIsImtpZCI6InNoaW1vZGV2IiwidHlwIjoiSldUIn0.eyJleHAiOjE3MzEwNTM3NDAsInVzZXJJZCI6ODYsIm1vZGUiOiJzaGltbyJ9.Z3P240q3Chd6r2PrsWtysBN50ixeGNrxZ-5PpAiBRUg",
	}
	_, err := sdkMgr.ActivateUserSeat(ctxTest, api.ActivateUserSeatReq{
		Metadata:                authLicense,
		ActivateUserSeatReqBody: api.ActivateUserSeatReqBody{UserIds: []string{"1", "2"}},
	})
	assert.Equal(t, nil, err)
}

func TestCancelUserSeat(t *testing.T) {
	authLicense := api.Metadata{
		ShimoToken: "eyJhbGciOiJIUzI1NiIsImtpZCI6InNoaW1vZGV2IiwidHlwIjoiSldUIn0.eyJleHAiOjE3MzEwNTM3NDAsInVzZXJJZCI6ODYsIm1vZGUiOiJzaGltbyJ9.Z3P240q3Chd6r2PrsWtysBN50ixeGNrxZ-5PpAiBRUg",
	}
	_, err := sdkMgr.CancelUserSeat(ctxTest, api.CancelUserSeatReq{
		Metadata:              authLicense,
		CancelUserSeatReqBody: api.CancelUserSeatReqBody{UserIds: []string{"1", "2"}},
	})
	assert.Equal(t, nil, err)
}

func TestBatchSetUserSeat(t *testing.T) {
	authLicense := api.Metadata{
		ShimoToken: "eyJhbGciOiJIUzI1NiIsImtpZCI6InNoaW1vZGV2IiwidHlwIjoiSldUIn0.eyJleHAiOjE3MzEwNTM3NDAsInVzZXJJZCI6ODYsIm1vZGUiOiJzaGltbyJ9.Z3P240q3Chd6r2PrsWtysBN50ixeGNrxZ-5PpAiBRUg",
	}
	_, err := sdkMgr.BatchSetUserSeat(ctxTest, api.BatchSetUserSeatReq{
		Metadata:                authLicense,
		BatchSetUserSeatReqBody: api.BatchSetUserSeatReqBody{UserIds: []string{"1", "2"}},
		Status:                  -1,
	})
	assert.Equal(t, nil, err)
}

// 文件操作

func TestImportFile(t *testing.T) {
	filePath := "./resource/import/test.doc"
	rootPath, err := filepath.Abs(filePath)
	file, err := os.Open(rootPath)
	defer file.Close()
	reqBody := api.ImportFileReqBody{
		FileID: "import-file-03",
		Type:   string(api.CollabFileTypeDocument),
		// File:   file,
		FileUrl:  "http://obs.cn-north-4.myhuaweicloud.com/shimo-devops/shimo-inspect-test/files/2024_12_09/task_3853_functional_dispatch_report.xlsx?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=KSXLSHEJJXCUHHX0TCKL%2F20241223%2Fcn-north-4%2Fs3%2Faws4_request&X-Amz-Date=20241223T071259Z&X-Amz-Expires=86400&X-Amz-SignedHeaders=host&X-Amz-Signature=8e39e54e5d344a58771a2661fca30c0b06e7d16135d37dd95d21f1e908292610",
		FileName: "task_3853_functional_dispatch_report.xlsx",
	}

	_, err = sdkMgr.ImportFile(ctxTest, api.ImportFileReq{
		Metadata:          auth,
		ImportFileReqBody: reqBody,
	})
	assert.Equal(t, nil, err)
}

func TestGetImportProgress(t *testing.T) {
	_, err := sdkMgr.GetImportProgress(ctxTest, api.GetImportProgReq{
		Metadata: auth,
		TaskId:   "5mwoAlZFv7vTyYkH",
	})
	assert.Equal(t, nil, err)
}

func TestExportFile(t *testing.T) {
	_, err := sdkMgr.ExportFile(ctxTest, api.ExportFileReq{
		Metadata: auth,
		FileID:   "910759b9232a4c73",
		Type:     string(api.CollabFileTypeTable),
	})
	assert.Equal(t, nil, err)
}

func TestGetExportProgress(t *testing.T) {
	_, err := sdkMgr.GetExportProgress(ctxTest, api.GetExportProgReq{
		Metadata: auth,
		TaskId:   "ItY7hT2bzC6LJXkN:1:2044354355:10000000087:table",
	})
	assert.Equal(t, nil, err)
}

func TestExportTableSheets(t *testing.T) {
	_, err := sdkMgr.ExportTableSheets(ctxTest, api.ExportTableSheetsReq{
		Metadata: auth,
		FileID:   "c912c5c9b1ed41bd",
	})
	assert.Equal(t, nil, err)
}

func TestCreatePreview(t *testing.T) {
	_, err := sdkMgr.CreatePreview(ctxTest, api.CreatePreviewReq{
		Metadata: auth,
		FileID:   "910759b9232a4c73",
	})
	assert.Equal(t, nil, err)
}

func TestAccessPreview(t *testing.T) {
	_, err := sdkMgr.AccessPreview(ctxTest, api.AccessPreviewReq{
		Metadata: auth,
		FileID:   "910759b9232a4c73",
	})
	assert.Equal(t, nil, err)
}
