package common

type CollabFileType string

const (
	CollabFileType_Document     CollabFileType = "document"
	CollabFileType_Spreadsheet  CollabFileType = "spreadsheet"
	CollabFileType_DocumentPro  CollabFileType = "documentPro"
	CollabFileType_Presentation CollabFileType = "presentation"
	CollabFileType_Table        CollabFileType = "table"
)
