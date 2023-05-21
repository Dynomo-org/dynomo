package github

type UploadFileParam struct {
	FilePathLocal         string
	FileName              string
	DestinationFolderPath string
	ReplaceIfNameExists   bool
}
